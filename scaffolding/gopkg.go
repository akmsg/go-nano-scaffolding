package scaffolding

import (
	"os"
	"path/filepath"
	"text/template"
)

func createDepGopkg(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/gopkg.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "Gopkg.toml"))
	if nil != errCreate {
		return errCreate
	}
	defer wr.Close()

	errExecute := t.Execute(wr, templateData)
	if nil != errExecute {
		return errExecute
	}
	return nil
}
