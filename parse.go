package main

import (
	"go/ast"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strconv"
	"strings"

	"github.com/wothing/log"
)

func parseGofile(file string) {

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// ast.Print(fset, f)
	// format.Node(os.Stdout, fset, f)

	mapper.Package = f.Name.Name

	for _, d := range f.Decls {
		genDecl, ok := d.(*ast.GenDecl)
		if !ok {
			format.Node(os.Stdout, fset, d)
			continue
		}

		switch genDecl.Tok {

		case token.IMPORT:
			mapper.Imports = getImports(genDecl)

		case token.TYPE:
			mapper.Name, mapper.Methods = getNameAndMethods(fset, genDecl)

		default:
			format.Node(os.Stdout, fset, d)
		}
	}

	// log.JSONIndent(mapper)

	info := types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
	}
	conf := types.Config{Importer: importer.Default()}
	_, err = conf.Check(mapper.Name, fset, []*ast.File{f}, &info)
	if err != nil {
		panic(err)
	}

	// log.Warn(pkg.Name())

	for _, obj := range info.Defs {
		if obj == nil {
			continue
		}

		itf, ok := obj.Type().Underlying().(*types.Interface)
		if !ok {
			// log.Warnf("%#v", obj)
			continue
		}

		for i := 0; i < itf.NumMethods(); i++ {
			for _, o := range mapper.Methods {
				if o.Name == itf.Method(i).Name() {
					getOperationInfo(o, itf.Method(i).Type().(*types.Signature))
					break
				}
			}
		}
	}
}

func getImports(gd *ast.GenDecl) map[string]string {
	imports := map[string]string{}
	for _, spec := range gd.Specs {
		is, _ := spec.(*ast.ImportSpec) // TODO must check error

		k, err := strconv.Unquote(is.Path.Value)
		if err != nil {
			panic(err)
		}
		if is.Name != nil {
			imports[k] = is.Name.Name + " " + is.Path.Value
		} else {
			imports[k] = is.Path.Value
		}
	}
	return imports
}

func getNameAndMethods(fset *token.FileSet, gd *ast.GenDecl) (name string, ms []*Operation) {
	ts, _ := gd.Specs[0].(*ast.TypeSpec)
	name = ts.Name.Name

	it, _ := ts.Type.(*ast.InterfaceType)

	ms = make([]*Operation, it.Methods.NumFields())
	for i, m := range it.Methods.List {
		var o Operation
		ms[i] = &o

		o.Doc = getComment(m.Doc)
		o.Name = m.Names[0].Name // TODO multiple

		// ft, _ := mapper.Type.(*ast.FuncType)

		// o.Params = getFeilds(ft.Params)
		// o.Results = getFeilds(ft.Results)
	}

	return
}

func getComment(cg *ast.CommentGroup) (comment string) {
	for _, c := range cg.List {
		comment += strings.TrimSpace(c.Text[2:]) + " " // remove `//`
	}
	return
}

//
// func getFeilds(fl *ast.FieldList) (vts []*Type) {
// 	vts = []*Type{}
// 	for _, field := range fl.List {
// 		for _, name := range field.Names {
// 			vts = append(vts, &Type{
// 				Var: name.Name,
// 			})
// 		}
// 	}
// 	return vts
// }

func getOperationInfo(o *Operation, sig *types.Signature) {
	o.ParamsOrder, o.Params = getFields(sig.Params())
	o.ResultsOrder, o.Results = getFields(sig.Results())
}

func getFields(ts *types.Tuple) (fos []*VarType, fs map[string]*Type) {
	fos = make([]*VarType, ts.Len())
	fs = make(map[string]*Type, ts.Len())
	for i := 0; i < ts.Len(); i++ {
		u := ts.At(i)

		v := u.Name()
		if v == "" {
			v = makeVarName(u.Type().String())
		}

		if t, ok := uses[u.Type().String()]; ok {
			fos[i] = &VarType{v, t, false}
			fs[v] = t
			continue
		}

		t := &Type{
			Type: u.Type().String(),
		}
		fos[i] = &VarType{v, t, false}
		fs[v] = t

		getOtherInfo(t, u.Type())

		uses[u.Type().String()] = t
	}
	return
}

func getOtherInfo(vt *Type, t types.Type) {
	// log.Debugf("%#v", t)
	// time.Sleep(100 * time.Millisecond)

	switch t.(type) {

	case *types.Slice:
		s := t.(*types.Slice)
		vt.Slice = true
		getOtherInfo(vt, s.Elem())

	case *types.Pointer:
		p := t.(*types.Pointer)
		vt.Pointer = true
		getOtherInfo(vt, p.Elem())

	case *types.Named:
		n := t.(*types.Named)
		if n.Obj().Pkg() != nil {
			vt.Name = n.Obj().Name()
			vt.Path = n.Obj().Pkg().Path()
			if imp, ok := mapper.Imports[vt.Path]; ok {
				if i := strings.Index(imp, " "); i != -1 {
					vt.Package = imp[:i]
				} else {
					vt.Package = n.Obj().Pkg().Name()
				}
			} else {
				mapper.Imports[vt.Path] = `"` + vt.Path + `"`
				vt.Package = n.Obj().Pkg().Name()
			}

			if vt.Name == "Tx" || vt.Name == "Time" {
				return
			}
			getOtherInfo(vt, n.Underlying())
		}

	case *types.Struct:
		s := t.(*types.Struct)
		vt.Fields = make(map[string]*Type, s.NumFields())
		for i := 0; i < s.NumFields(); i++ {
			f := s.Field(i)

			if v, ok := uses[f.Type().String()]; ok {
				// 避免 json 打印时 递归
				x := *v
				x.Fields = nil
				vt.Fields[f.Name()] = &x

				continue
			}

			x := &Type{
				Type: f.Type().String(),
			}
			vt.Fields[f.Name()] = x

			if x.Type == "time.Time" {
				x.Primitive = x.Type
				return
			}

			uses[f.Type().String()] = x

			getOtherInfo(x, f.Type())
		}

	case *types.Basic:
		b := t.(*types.Basic)
		vt.Primitive = b.Name()

	case *types.Map:
		// map[x]y

	default:
		log.Infof("%s %#v", t.String(), t)
	}
}

func makeVarName(t string) string {
	if t == "error" {
		return "err"
	}

	i := strings.LastIndex(t, ".")

	s := ""
	if t[0:2] == "[]" {
		s = "s"
	}

	return strings.ToLower(t[i+1:i+2]) + s
}
