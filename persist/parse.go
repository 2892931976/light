package persist

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gotips/log"
)

func parseFile(gofile string) (i *Interface, err error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, gofile, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// ast.Print(fset, f)
	// format.Node(os.Stdout, fset, f)

	i = &Interface{}

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		switch genDecl.Tok {
		case token.IMPORT:
			parseImports(genDecl, i)

		case token.TYPE:
			parseType(genDecl, i)
		}
	}

	deepParseStruct(i)

	return i, nil
}

func deepParseStruct(i *Interface) {
	for _, fun := range i.Methods {
		for _, p := range fun.Params {
			if strings.Contains(p.Type, ".") {
				parseStruct(p, i.Imports)
			}
		}
		for _, p := range fun.Returns {
			if strings.Contains(p.Type, ".") {
				parseStruct(p, i.Imports)
			}
		}
	}
}

// TODO
func parseStruct(p *Param, imps []string) {
	// get path and type name
	path, name := getPathAndTypeName(p, imps)

	// parse file to ast, then get type properties
	typeSpec := getTypeSpec(path, name)

	stype, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		log.Fatal(typeSpec.Name.Name + " not struct")
	}

	for _, f := range stype.Fields.List {
		if len(f.Names) == 0 {
			// Embedded struct: recurse
			// TODO
			continue
		}
		typ := parseExpr(f.Type)
		for _, nam := range f.Names {
			p.Props = append(p.Props, &Param{
				Name: nam.Name,
				Type: typ,
			})
		}
	}
}

func getPathAndTypeName(p *Param, imps []string) (path, name string) {
	typ := p.Type

	if strings.HasPrefix(typ, "[]") {
		typ = typ[2:]
	}
	if typ[0] == '*' {
		typ = typ[1:]
	}
	dotIdx := strings.Index(typ, ".")
	path, name = typ[:dotIdx], typ[dotIdx+1:]

	var err error
	for _, imp := range imps {
		if strings.HasSuffix(imp, path+`"`) {
			path, err = strconv.Unquote(imp)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
		if strings.HasPrefix(imp, path+" ") {
			path = imp[len(path)+1:]
			path, err = strconv.Unquote(path)
			if err != nil {
				log.Fatalf("unquote %s error: %s", imp[len(path)+1:], err)
			}
			break
		}
	}
	return path, name
}

func getTypeSpec(path, name string) (typeSpec *ast.TypeSpec) {
	pkg, _ := build.Import(path, "", 0)
	fset := token.NewFileSet() // share one fset across the whole package
	for _, file := range pkg.GoFiles {
		f, err := parser.ParseFile(fset, filepath.Join(pkg.Dir, file), nil, 0)
		if err != nil {
			continue
		}

		for _, decl := range f.Decls {
			decl, ok := decl.(*ast.GenDecl)
			if !ok || decl.Tok != token.TYPE {
				log.Debugf("%#v", decl)
				continue
			}
			for _, spec := range decl.Specs {
				spec := spec.(*ast.TypeSpec)
				if spec.Name.Name != name {
					continue
				}
				typeSpec = spec
			}
		}
	}

	if typeSpec == nil {
		log.Fatalf("%s.%s not exist", path, name)
	}

	return typeSpec
}

func parseImports(genDecl *ast.GenDecl, i *Interface) {
	for _, spec := range genDecl.Specs {
		importSpec, ok := spec.(*ast.ImportSpec)
		if !ok {
			continue
		}

		path := ""
		if importSpec.Name != nil {
			path += importSpec.Name.Name + " "
		}
		path += importSpec.Path.Value

		i.Imports = append(i.Imports, path)
	}
}

func parseType(genDecl *ast.GenDecl, i *Interface) {
	for _, spec := range genDecl.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		i.Name = typeSpec.Name.Name

		interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)
		if !ok {
			continue
		}

		parseMethods(interfaceType, i)
	}
}

func parseMethods(interfaceType *ast.InterfaceType, i *Interface) {
	for _, m := range interfaceType.Methods.List {
		var f Func
		i.Methods = append(i.Methods, &f)

		f.Doc = getDoc(m.Doc)

		f.Name = m.Names[0].Name

		funcType, ok := m.Type.(*ast.FuncType)
		if !ok {
			continue
		}

		parseFuncType(funcType, &f)
	}
}

func parseFuncType(funcType *ast.FuncType, i *Func) {
	for _, field := range funcType.Params.List {
		i.Params = append(i.Params, parseField(field)...)
	}

	for _, field := range funcType.Results.List {
		i.Returns = append(i.Returns, parseField(field)...)
	}
}

func parseField(field *ast.Field) (rets []*Param) {
	typ := parseExpr(field.Type)

	for _, name := range field.Names {
		rets = append(rets, &Param{
			Name: name.Name,
			Type: typ,
		})
	}
	if len(field.Names) == 0 {
		rets = append(rets, &Param{
			Name: "",
			Type: typ,
		})
	}

	return rets
}

func parseExpr(expr ast.Expr) (x string) {
	switch expr.(type) {
	case *ast.Ident:
		ident := expr.(*ast.Ident)
		return ident.Name

	case *ast.StarExpr:
		starExpr := expr.(*ast.StarExpr)
		return "*" + parseExpr(starExpr.X)

	case *ast.SelectorExpr:
		selectorExpr := expr.(*ast.SelectorExpr)
		return parseExpr(selectorExpr.X) + "." + selectorExpr.Sel.Name

	case *ast.ArrayType:
		arrayType := expr.(*ast.ArrayType)
		return "[]" + parseExpr(arrayType.Elt)

	default:
		panic("not implemented")
	}
}

func getDoc(g *ast.CommentGroup) (doc string) {
	for _, comment := range g.List {
		doc += " " + strings.TrimLeft(comment.Text, " /")
	}

	return doc
}