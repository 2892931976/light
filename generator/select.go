package generator

import (
	"bytes"

	"github.com/arstd/light/goparser"
	"github.com/arstd/light/sqlparser"
)

func writeSelect(buf *bytes.Buffer, m *goparser.Method, stmt *sqlparser.Statement) {
	wln := func(s string) { buf.WriteString(s + "\n") }

	wln("var buf bytes.Buffer")
	wln("var args []interface{}")

	for _, f := range stmt.Fragments {
		writeFragment(buf, m, f)
	}

	if m.Results.Len() == 3 {
		writePage(buf, m, stmt)
	} else if m.Results.Len() == 2 {
		if m.Results.At(0).IsSlice() {
			writeList(buf, m, stmt)
		} else {
			writeGet(buf, m, stmt)
		}
	}
}
