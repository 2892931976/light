package prepare

import (
	"strings"

	"github.com/arstd/light/domain"
	"github.com/arstd/log"
)

func getReturnings(m *domain.Method) (rs []*domain.VarType) {
	switch m.Kind {
	case domain.Batch, domain.Update, domain.Delete:
		return nil
	case domain.Insert:
		return getInsertReturnings(m)
	case domain.Get, domain.List, domain.Page:
		return getFieldsReturings(m)
	}
	return rs
}

func getInsertReturnings(m *domain.Method) (rs []*domain.VarType) {
	stmt := m.Fragments[len(m.Fragments)-1].Stmt

	fs := strings.Split(stmt[(strings.Index(stmt, "returning ")+len("returning ")):], ",")
	rs = make([]*domain.VarType, len(fs))
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

func getFieldsReturings(m *domain.Method) (rs []*domain.VarType) {
	stmt := m.Fragments[0].Stmt

	stmt = stmt[len("select "):strings.Index(stmt, " from ")]
	fs := strings.Split(stmt, ",")
	rs = make([]*domain.VarType, len(fs))
	for i, f := range fs {
		f = strings.TrimSpace(f)
		// TODO model index
		for _, vt := range m.Results[0].Fields {
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
