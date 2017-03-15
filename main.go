package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/arstd/light/parse"
	"github.com/arstd/light/prepare"
	"github.com/arstd/light/util"
	"github.com/arstd/log"
)

func init() {
	// log.SetLevel(log.Lwarn)
	log.SetFormat("2006-01-02 15:04:05.999 info examples/main.go:88 message")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: light [flags] [file.go]\n\t//go:generate light [flags] [file.go]")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
}

func main() {
	goFile := os.Getenv("GOFILE")
	if goFile == "" {
		if flag.NArg() > 1 {
			goFile = flag.Arg(0)
			if !strings.HasSuffix(goFile, ".go") {
				fmt.Println("file suffix must match *.go")
				return
			}
		} else {
			flag.Usage()
		}
	}
	fmt.Printf("Found  go file: %s\n", goFile)

	p := parse.ParseGoFile(goFile)

	prepare.Prepare(p)

	// log.JSONIndent(p)

	paths := strings.Split(os.Getenv("GOPATH"), string(filepath.ListSeparator))
	tmplFile := filepath.Join(paths[0], "src", "github.com/arstd/light", "templates/pq.tmpl")

	funcMap := template.FuncMap{
		"timestamp": func() string { return time.Now().Format("2006-01-02 15:04:05") },
	}

	tmpl, err := template.New("pq.tmpl").Funcs(funcMap).ParseFiles(tmplFile)
	util.CheckError(err)

	// var buf bytes.Buffer
	err = tmpl.Execute(os.Stdout, p)
	util.CheckError(err)

	// outFile := goFile[:len(goFile)-3] + "impl.go"
	// pretty, err := imports.Process(outFile, buf.Bytes(), nil)
	// util.CheckError(err)
	// err = ioutil.WriteFile(outFile, pretty, 0644)
	// util.CheckError(err)
	// fmt.Printf("Generated file: %s\n", outFile)
}
