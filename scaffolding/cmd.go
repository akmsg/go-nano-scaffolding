package scaffolding

import (
	"os"
	"path/filepath"
	"text/template"
)

func createCmd(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/cmd/main.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "cmd/main.go"))
	if nil != errCreate {
		return errCreate
	}
	defer wr.Close()

	errExecuteMain := t.Execute(wr, templateData)
	if nil != errExecuteMain {
		return errExecuteMain
	}
	return nil
}
