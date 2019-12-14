package scaffolding

import (
	"os"
	"path/filepath"
	"text/template"
)

func createAppConf(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/app_conf.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "app_conf.yaml"))
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
