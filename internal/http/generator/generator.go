package generator

import (
	"embed"
	"os"
	"text/template"
)

//go:embed templates
var efs embed.FS

type generator struct {
	err  error
	tmpl *template.Template
}

func (a *generator) Error() error {
	return a.err
}

func (a *generator) Setup() {
	if a.err != nil {
		return
	}
	entries := a.readEntries()
	wd, _ := os.Getwd()
	data := struct{ Functions []string }{Functions: entries}
	a.writeTemplate("templates/router.tmpl", data, wd+"/router.go")
}

func (a *generator) Teardown() {
	if a.err != nil {
		return
	}
}

func (a *generator) readEntries() (entries []string) {
	wd, _ := os.Getwd()
	infos, err := os.ReadDir(wd + "/functions")
	if err != nil {
		a.err = err
		return
	}
	for _, info := range infos {
		if !info.IsDir() || info.Name() == "" {
			continue
		}
		entries = append(entries, info.Name())
	}
	return
}

func (a *generator) writeTemplate(src string, data any, dst string) {
	if a.err != nil {
		return
	}
	content, err := efs.ReadFile(src)
	if err != nil {
		a.err = err
		return
	}
	tmpl, err := template.New("tmpl").Parse(string(content))
	if err != nil {
		a.err = err
		return
	}
	file, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_SYNC|os.O_TRUNC, 0644)
	if err != nil {
		a.err = err
		return
	}
	defer file.Close()
	if err := tmpl.Execute(file, data); err != nil {
		a.err = err
	}
}

func New() *generator {
	return &generator{}
}
