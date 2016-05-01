package persist

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func prepare(itf *Interface) (impl *Implement) {
	impl = &Implement{}
	impl.Source = os.Getenv("GOFILE")
	impl.Package = os.Getenv("GOPACKAGE")

	impl.Imports = itf.Imports

	impl.Name = itf.Name
	if strings.HasSuffix(impl.Name, "Persister") {
		impl.Name = impl.Name[:len(impl.Name)-2]
	} else {
		impl.Name += impl.Name + "Impl"
	}

	for _, f := range itf.Methods {
		var m Method
		impl.Methods = append(impl.Methods, &m)

		prepareMethod(&m, f)
	}

	return impl
}

func prepareMethod(m *Method, f *Func) {
	m.Name = f.Name
	m.Params = f.Params
	m.Returns = f.Returns

	sql := strings.Trim(f.Doc, " \t")
	m.Type = getMethodType(sql, f)

	calcInResult(m, f)

	calcArgs(m, sql)

	calcMarshals(m)

	calcScans(m, sql)

	calcUnmarshals(m)
}

func calcInResult(m *Method, f *Func) {
	if m.Type == MethodTypeGet {
		m.In = f.Params[0].Name
		m.Result, m.ResultType = f.Returns[0].Name, f.Returns[0].Type[1:]
	} else if m.Type == MethodTypeList {
		m.In = f.Params[0].Name
		m.Result, m.ResultType = f.Returns[0].Name, f.Returns[0].Type[3:]
	} else if m.Type == MethodTypePage {
		m.In = f.Params[0].Name
		m.Result, m.ResultType = f.Returns[1].Name, f.Returns[1].Type[3:]
	} else {
		m.In = f.Params[0].Name
		m.Result, m.ResultType = f.Params[0].Name, f.Params[0].Type[1:]
	}
}

func getMethodType(sql string, f *Func) MethodType {
	switch sql[:strings.Index(sql, " ")] {
	case "insert":
		return MethodTypeAdd
	case "update":
		return MethodTypeModify
	case "delete":
		return MethodTypeRemove
	case "select":
		if len(f.Returns) == 3 {
			return MethodTypePage
		} else {
			if len(f.Returns) > 0 && strings.HasPrefix(f.Returns[0].Type, "[]") {
				return MethodTypeList
			} else {
				return MethodTypeGet
			}
		}
	default:
		panic("sql error: " + sql)
	}
}

func calcArgs(m *Method, sql string) {
	re := regexp.MustCompile(`\$\{(.+?)\}`)
	matched := re.FindAllStringSubmatchIndex(sql, -1)

	if len(matched) == 0 {
		m.Prefix = sql
		return
	}

	// i.e.
	// select id, demo_name, demo_status
	// from demos
	// where id < ${id} and demo_name=${d.demeName} and demo_status=1
	var from int
	for i, group := range matched {
		//= select ... from demos where id <
		m.Prefix += sql[from:group[0]]

		//= select ... from demos where id < $1
		m.Prefix += "$" + strconv.Itoa(i+1)

		m.Args = append(m.Args, sql[group[2]:group[3]])

		from = group[1]
	}

	//= select ... from demos where id < $1 ... and demo_status=1
	m.Prefix += sql[matched[len(matched)-1][1]:]
}

func calcScans(m *Method, sql string) {
	var start, end int
	switch m.Type {
	case MethodTypeGet, MethodTypeList, MethodTypePage:
		start, end = 6, strings.Index(sql, " from ")

	case MethodTypeAdd, MethodTypeModify, MethodTypeRemove:
		start = strings.Index(sql, " returning ")
		if start == -1 {
			return
		}
		start += 11
		end = strings.Index(sql, " on conflict ")
		if end == -1 {
			end = len(sql)
		}

	default:
		panic("unreachable code")
	}

	fields := strings.Split(sql[start:end], ",")
	for _, f := range fields {
		f = strings.Trim(f, " \t")
		f = strings.Replace(f, "_", " ", -1)
		f = strings.Title(f)
		f = strings.Replace(f, " ", "", -1)
		m.Scans = append(m.Scans, "x."+f)
	}
}

func calcMarshals(m *Method) {
	for _, p := range m.Params {
		for _, prop := range p.Props {
			switch prop.Type {
			case "int", "int64", "int32", "int16", "int8":
			case "uint", "uint64", "uint32", "uin16", "uint8", "byte":
			case "string":
			default:
				for i, arg := range m.Args {
					// TODO must not use m.Result
					if arg == m.In+"."+prop.Name {
						m.Args[i] = m.In + "_" + prop.Name
						m.Marshals = append(m.Marshals, prop.Name)
					}
				}
			}
		}
	}
}

func calcUnmarshals(m *Method) {
	for _, p := range m.Returns {
		for _, prop := range p.Props {
			switch prop.Type {
			case "int", "int64", "int32", "int16", "int8":
			case "uint", "uint64", "uint32", "uin16", "uint8", "byte":
			case "string":
			default:
				for i, scan := range m.Scans {
					if scan == "x."+prop.Name {
						m.Scans[i] = "x_" + prop.Name
						m.Unmarshals = append(m.Unmarshals, prop.Name)
					}
				}
			}
		}
	}
}
