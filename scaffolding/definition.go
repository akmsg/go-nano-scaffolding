package scaffolding

import (
	"os"
	"path/filepath"
	"text/template"
)

func createDefinitionDoc(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/definition/doc.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "definition/doc.go"))
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

func createDefinitionServiceProto(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/definition/service.proto.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "definition/service.proto"))
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

func createDefinitionMakefile(servicePath string, templateData map[string]string) error {
	t, errParse := template.ParseFiles("templates/definition/makefile.tmpl")
	if nil != errParse {
		return errParse
	}

	wr, errCreate := os.Create(filepath.Join(servicePath, "definition/Makefile"))
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

func createDefinition(servicePath string, templateData map[string]string) error {
	errCreateDefinitionDoc := createDefinitionDoc(servicePath, templateData)
	if nil != errCreateDefinitionDoc {
		return errCreateDefinitionDoc
	}

	errCreateDefinitionServiceProto := createDefinitionServiceProto(servicePath, templateData)
	if nil != errCreateDefinitionServiceProto {
		return errCreateDefinitionServiceProto
	}

	errCreateDefinitionMakefile := createDefinitionMakefile(servicePath, templateData)
	if nil != errCreateDefinitionMakefile {
		return errCreateDefinitionMakefile
	}

	return nil
}
