package generator

import (
	"bytes"

	"github.com/arstd/light/goparser"
	"github.com/arstd/light/sqlparser"
)

func writeGet(buf *bytes.Buffer, m *goparser.Method, stmt *sqlparser.Statement) {
	w := buf.WriteString
	wln := func(s string) { buf.WriteString(s + "\n") }

	wln("query := buf.String()")
	if m.Store.Log {
		wln("log.Debug(query)")
		wln("log.Debug(args...)")
	}

	wln("row := exec.QueryRow(query, args...)")

	v := m.Results.Result()

	w("xu := new(")
	w(v.TypeName())
	wln(")")
	w("xdst := []interface{}{")
	for _, f := range stmt.Fields {
		s := m.Results.Result()
		v := s.VarByTag(f)
		name := "xu." + v.VName
		w(v.Scan(name))
		w(",")
	}
	buf.Truncate(buf.Len() - 1)
	wln("}")

	wln("err := row.Scan(xdst...)")
	if m.Store.Log {
		wln("if err != nil {")
		wln(`if err == sql.ErrNoRows {
			return nil, nil
		}`)
		wln("log.Error(query)")
		wln("log.Error(args...)")
		wln("log.Error(err)")
		wln("return nil, err")
		wln("}")
		wln(`log.JSON(xdst)`)
	}

	wln("return xu, err")
}
