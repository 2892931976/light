package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"golang.org/x/tools/imports"

	"github.com/arstd/log"
)

func usage() {
	fmt.Fprintln(os.Stderr, `usage: light [flags] [file.go]
	//go:generate light [flags] [file.go]
`)
	flag.PrintDefaults()

	fmt.Fprintln(os.Stderr, `
examples:
	light -force -dbvar=db.DB -dbpath=github.com/arstd/light/example/mapper
	light -force -dbvar=db2.DB -dbpath=github.com/arstd/light/example/mapper
`)
	os.Exit(2)
}

func main() {
	log.SetLevel(log.Linfo)
	log.SetFormat("2006-01-02 15:04:05.999 info examples/main.go:88 message")

	dbVar := flag.String("dbvar", "db", "variable of db to open transaction and execute SQL statements")
	dbPath := flag.String("dbpath", "", "path of db to open transaction and execute SQL statements")
	skip := flag.Bool("skip", true, "skip generate if sourceimpl.go file newer than source.go file")
	quick := flag.Bool("quick", true, "if true, use go/types to parse dependences, much fast when built pkg cached; \n        if false, use go/loader parse source and dependences, much slow")
	version := flag.Bool("v", false, "variable of db to open transaction and execute SQL statements")
	flag.Usage = usage

	flag.Parse()
	if *version {
		fmt.Println("0.5.5")
		return
	}

	goFile := os.Getenv("GOFILE")
	if goFile == "" {
		if flag.NArg() > 0 {
			goFile = flag.Arg(0)
			if !strings.HasSuffix(goFile, ".go") {
				fmt.Println("file suffix must match *.go")
				return
			}
		} else {
			flag.Usage()
		}
	}
	fmt.Printf("Found source file %s.\n", goFile)

	outFile := goFile[:len(goFile)-3] + "impl.go"
	if *skip {
		outStat, err := os.Stat(outFile)
		if err != nil {
			// log.Info(err)
		} else {
			goStat, _ := os.Stat(goFile)
			if !outStat.ModTime().Before(goStat.ModTime()) {
				fmt.Print("Skip!\n")
				return
			}
		}
	}
	os.Remove(outFile)

	*dbVar = strings.Trim(*dbVar, `'"`)
	pkg := &Package{
		Source:  goFile,
		DBVar:   *dbVar,
		Imports: map[string]string{},
	}
	*dbPath = strings.Trim(*dbPath, `'"`)
	if *dbPath != "" {
		ss := strings.Split(*dbVar, ".")
		if len(ss) != 2 {
			fmt.Println("arg 'dbvar' must be <package-name>:<variable-name")
			flag.Usage()
			return
		}
		pkg.Imports[ss[0]] = strings.Trim(*dbPath, `'"`)
	}

	if *quick {
		parseGoFile(pkg)
	} else {
		parseGoFileByLoader(pkg)
	}

	prepare(pkg)
	// log.JSONIndent(pkg)

	paths := strings.Split(os.Getenv("GOPATH"), string(filepath.ListSeparator))
	tmplFile := filepath.Join(paths[0], "src", "github.com/arstd/light", "postgresql.pq.gotemplate")

	funcMap := template.FuncMap{
		"timestamp": func() string { return time.Now().Format("2006-01-02 15:04:05") },
	}

	tmpl, err := template.New("postgresql.pq.gotemplate").Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		log.Panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, pkg)
	if err != nil {
		log.Panic(err)
	}

	ioutil.WriteFile(outFile, buf.Bytes(), 0644)
	fmt.Printf("Generate implementation file %s.\n", outFile)

	pretty, err := imports.Process(outFile, buf.Bytes(), nil)
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(outFile, pretty, 0644)
	if err != nil {
		log.Panic(err)
	}
}
