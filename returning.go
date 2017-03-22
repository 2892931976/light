package main

import (
	"strings"

	"github.com/arstd/log"
)

func getReturnings(m *Method) (rs []*VarType) {
	switch m.Kind {
	case Count:
		return nil
	case Batch, Update, Delete:
		return nil
	case Insert:
		return getInsertReturnings(m)
	case Get, List:
		return getFieldsReturings(m, 0)
	case Page:
		return getFieldsReturings(m, 1)
	}
	return rs
}

func getInsertReturnings(m *Method) (rs []*VarType) {
	stmt := m.Fragments[len(m.Fragments)-1].Stmt

	fs := strings.Split(stmt[(strings.Index(stmt, "returning ")+len("returning ")):], ",")
	rs = make([]*VarType, len(fs))
	for i, f := range fs {
		f = strings.TrimSpace(f)
		// TODO model index ?= 1
		for _, vt := range m.Params[1].Fields {
			if strings.HasPrefix(vt.Tag, f) {
				rs[i] = vt
				break
			}
		}
		if rs[i] == nil {
			log.Panicf("returning `%s` no matched field for method `%s`", f, m.Name)
		}
	}

	return rs
}

func getFieldsReturings(m *Method, idx int) (rs []*VarType) {
	stmt := m.Fragments[0].Stmt

	stmt = stmt[len("select "):strings.Index(stmt, " from ")]
	fs := strings.Split(stmt, ",")
	rs = make([]*VarType, len(fs))
	for i, f := range fs {
		fs := strings.Split(f, " ")
		f = fs[len(fs)-1]
		f = strings.TrimSpace(f)
		// TODO model index
		for _, vt := range m.Results[idx].Fields {
			if strings.HasPrefix(vt.Tag, f) {
				rs[i] = vt
				break
			}
		}
		if rs[i] == nil {
			log.Panicf("returning `%s` no matched field for method `%s`", f, m.Name)
		}
	}

	return rs
}
