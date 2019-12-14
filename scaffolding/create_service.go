package scaffolding

import (
	"fmt"
	"go/build"
	"path/filepath"
	"strings"
)

func generateBoilerPlate(serviceName, outputDir string) error {
	servicePath := filepath.Join(outputDir, serviceName)
	goSrc := filepath.Join(build.Default.GOPATH, "src")
	baseImportPath, errBaseImportPath := filepath.Rel(goSrc, servicePath)
	exportableServiceName := strings.ToUpper(serviceName[0:1]) + serviceName[1:]

	if nil != errBaseImportPath {
		return fmt.Errorf("failed to create baseImportPath: %v", errBaseImportPath)
	}

	td := map[string]string{
		"ServiceName":           serviceName,
		"ExportableServiceName": exportableServiceName,
		"ServicePath":           servicePath,
		"BaseImportPath":        baseImportPath,
	}

	// 1. cmd/
	errCreateCmd := createCmd(serviceName, td)
	if nil != errCreateCmd {
		return fmt.Errorf("failed to createCmd: %v", errCreateCmd)
	}

	// 2. definition/
	errCreateDefinition := createDefinition(serviceName, td)
	if nil != errCreateDefinition {
		return fmt.Errorf("failed to createDefinition: %v", errCreateDefinition)
	}

	// 3. logger/
	errCreateLogger := createLogger(serviceName, td)
	if nil != errCreateLogger {
		return fmt.Errorf("failed to createLogger: %v", errCreateLogger)
	}

	// 4. server/
	errCreateServer := createServer(serviceName, td)
	if nil != errCreateServer {
		return fmt.Errorf("failed to createServer: %v", errCreateServer)
	}

	// 5. app_config.yaml
	errCreateAppConf := createAppConf(serviceName, td)
	if nil != errCreateAppConf {
		return fmt.Errorf("failed to createAppConf: %v", errCreateAppConf)
	}

	// 6. impl.go
	errCreateGenerator := createImpl(serviceName, td)
	if nil != errCreateGenerator {
		return fmt.Errorf("failed to createGenerator: %v", errCreateGenerator)
	}

	// 7. service.go
	errCreateServiceGenerator := createServiceGenerator(serviceName, td)
	if nil != errCreateServiceGenerator {
		return fmt.Errorf("failed to createServiceGenerator: %v", errCreateServiceGenerator)
	}

	// 8. dep Gopkg.toml
	errCreateDepGopkg := createDepGopkg(serviceName, td)
	if nil != errCreateDepGopkg {
		return fmt.Errorf("failed to createDepGopkg: %v", errCreateDepGopkg)
	}

	return nil
}

// CreateService will scaffold serviceName at outputDir
func CreateService(serviceName, outputDir string) error {
	//todo: handle dir creation individually
	errCreateDirs := createDirs(serviceName, outputDir)
	if errCreateDirs != nil {
		fmt.Printf("errCreateDirs: %v\n", errCreateDirs)
	}

	errGenerateBoilerPlate := generateBoilerPlate(serviceName, outputDir)
	if errGenerateBoilerPlate != nil {
		return errGenerateBoilerPlate
	}
	return nil
}
