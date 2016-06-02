package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

var db, path string

func main() {
	flag.StringVar(&db, "db", "db", "db variable to Query/QueryRow/Exec")
	flag.StringVar(&path, "path", "", "db variable path")

	flag.Parse()
	if path != "" && path[0] != '"' {
		path = `"` + path + `"`
	}

	gofile := os.Getenv("GOFILE")

	filename := gofile[:len(gofile)-3] + "impl.go"
	os.Remove(filename)

	meta, err := parseFile(gofile)
	if err != nil {
		panic(err)
	}
	// log.JSONIndent(meta)

	data := prepare(meta)
	// log.JSONIndent(data)

	tpl, err := Asset("template.txt")
	if err != nil {
		panic(err)
	}
	t, err := template.New("template.txt").Parse(string(tpl))
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		panic(err)
	}

	// ioutil.WriteFile(filename, buf.Bytes(), 0644)

	pretty, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(filename, pretty, 0644)

	pwd, _ := os.Getwd()
	fmt.Printf("Generate file %s/%s\n", pwd, filename)

	cmd := exec.Command("goimports", "-w", pwd+"/"+filename)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
