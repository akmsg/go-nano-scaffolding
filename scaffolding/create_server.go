package scaffolding

import (
	"os"
	"path/filepath"
	"text/template"
)

func createServerHandler(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/server/handlers/echo.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "server/handler/echo.go"))
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

func createServerConf(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/server/conf/conf.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "server/conf/conf.go"))
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

func createServer(servicePath string, templateData map[string]string) error {
	errCreateServerHandler := createServerHandler(servicePath, templateData)
	if nil != errCreateServerHandler {
		return errCreateServerHandler
	}

	errCreateServerConf := createServerConf(servicePath, templateData)
	if nil != errCreateServerConf {
		return errCreateServerConf
	}

	return nil
}
