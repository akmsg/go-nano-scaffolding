package scaffolding

import (
	"os"
	"path/filepath"
	"text/template"
)

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

func createServerHandlerServiceImpl(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/server/handlers/service_impl.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "server/handler/service_impl.go"))
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

func createServerDoc(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/server/doc.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "server/doc.go"))
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

func createServerCleanup(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/server/cleanup.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "server/cleanup.go"))
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

func createServerServe(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/server/serve.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "server/serve.go"))
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

func createServerServeUtils(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/server/serve_utils.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "server/serve_utils.go"))
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
	errCreateServerConf := createServerConf(servicePath, templateData)
	if nil != errCreateServerConf {
		return errCreateServerConf
	}

	errCreateServerHandler := createServerHandler(servicePath, templateData)
	if nil != errCreateServerHandler {
		return errCreateServerHandler
	}

	errCreateServerHandlerServiceImpl := createServerHandlerServiceImpl(servicePath, templateData)
	if nil != errCreateServerHandlerServiceImpl {
		return errCreateServerHandlerServiceImpl
	}

	errCreateServerDoc := createServerDoc(servicePath, templateData)
	if nil != errCreateServerDoc {
		return errCreateServerDoc
	}
	errCreateServerCleanup := createServerCleanup(servicePath, templateData)
	if nil != errCreateServerCleanup {
		return errCreateServerCleanup
	}
	errCreateServerServe := createServerServe(servicePath, templateData)
	if nil != errCreateServerServe {
		return errCreateServerServe
	}

	errCreateServerServeUtils := createServerServeUtils(servicePath, templateData)
	if nil != errCreateServerServeUtils {
		return errCreateServerServeUtils
	}

	return nil
}
