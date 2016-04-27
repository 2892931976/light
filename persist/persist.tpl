package persist

import (
	"github.com/arstd/persist/examples/domain"

	// import sql driver
	_ "github.com/lib/pq"
	"github.com/wothing/log"
)

type {{ .Name }} struct{}

{{- range .Funcs }}

func (*{{ $.Name }}) {{ .Name }}({{range $i,$vt := .Params}}{{if $i | ne 0}}, {{end}}{{$vt.Var}} {{$vt.Type}}{{ end }}) ({{range $i,$vt := .Returns}}{{if $i | ne 0}}, {{end}}{{$vt.Var}} {{$vt.Type}}{{ end }}) {
	q := "{{ .Stmt | print }}"
{{if .Type | eq "add" }}
	err = db.QueryRow(q{{range .Args}}, {{.}}{{ end }}).Scan({{range $i, $r := .Returning}}{{ if ne $i 0 }}, {{end}}&{{ $r}}{{ end }})
	if err != nil {
		log.Errorf("insert(%s{{range .Args}}, %s{{ end }}) error: %s", q{{range .Args}}, {{.}}{{ end }}, err)
		return err
	}
	return nil

{{- else if .Type | eq "modify"}}
	res, err := db.Exec(q{{range .Args}}, {{.}}{{ end }})
	if err != nil {
		log.Errorf("update(%s{{range .Args}}, %s{{ end }}) error: %s", q{{range .Args}}, {{.}}{{ end }}, err)
		return err
	}
	a, err := res.RowsAffected()
	if err != nil {
		log.Errorf("update(%s{{range .Args}}, %s{{ end }}) error: %s", q{{range .Args}}, {{.}}{{ end }}, err)
		return err
	} else if a != 1 {
		log.Errorf("update(%s{{range .Args}}, %s{{ end }}) expected affected 1 row, but actual affected %d rows",
			q{{range .Args}}, {{.}}{{ end }}, a)
		return err
	}
	return nil

{{- else if .Type | eq "remove"}}
	res, err := db.Exec(q{{range .Args}}, {{.}}{{ end }})
	if err != nil {
		log.Errorf("delete(%s{{range .Args}}, %s{{ end }}) error: %s", q{{range .Args}}, {{.}}{{ end }}, err)
		return err
	}
	a, err := res.RowsAffected()
	if err != nil {
		log.Errorf("delete(%s{{range .Args}}, %s{{ end }}) error: %s", q{{range .Args}}, {{.}}{{ end }}, err)
		return err
	} else if a != 1 {
		log.Errorf("delete(%s{{range .Args}}, %s{{ end }}) expected affected 1 row, but actual affected %d rows",
			q{{range .Args}}, {{.}}{{ end }}, a)
		return err
	}
	return nil

{{- else if .Type | eq "get"}}
	x := {{.ResultType}}{}
	err = db.QueryRow(q{{range .Args}}, {{.}}{{ end }}).
		Scan({{range $i, $r := .Scans}}{{ if ne $i 0 }}, {{end}}&x.{{$r}}{{ end }})
	if err != nil {
		log.Errorf("query(%s{{range .Args}}, %s{{ end }}) error: %s", q{{range .Args}}, {{.}}{{ end }}, err)
		return nil, err
	}
	return &x, nil

{{- else if .Type | eq "list"}}
	rows, err := db.Query(q{{range .Args}}, {{.}}{{ end }})
	if err != nil {
		log.Errorf("query(%s{{range .Args}}, %s{{ end }}) error: %s", q{{range .Args}}, {{.}}{{ end }}, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var x {{.ResultType}}
		err = rows.Scan({{range $i, $r := .Scans}}{{ if ne $i 0 }}, {{end}}&x.{{$r}}{{ end }})
		if err != nil {
			log.Errorf("scan rows for query(%s{{range .Args}}, %s{{ end }}) error: %s", q{{range .Args}}, {{.}}{{ end }}, err)
			return nil, err
		}
		{{.Result}} = append({{.Result}}, &x)
	}
	if err = rows.Err(); err != nil {
		log.Errorf("scan rows for query(%s{{range .Args}}, %s{{ end }}) last error: %s", q{{range .Args}}, {{.}}{{ end }}, err)
		return nil, err
	}
	return {{.Result}}, nil

{{- end}}
}
{{- end}}
