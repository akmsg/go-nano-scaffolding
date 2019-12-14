package scaffolding

import (
	"fmt"
	"os"
	"path/filepath"
)

// createDirs creates directory structure for serviceName at outputDir
// serviceName
//	├── bin
//	├── cmd
//	├── definition
//	├── logger
//	└── server
//		└── conf
// 		└── handler
//
func createDirs(serviceName, outputDir string) []error {
	errListCreateDirs := []error{}

	errCreateDirs := os.Mkdir(filepath.Join(outputDir, serviceName), 0754)
	if errCreateDirs != nil {
		errListCreateDirs = append(errListCreateDirs, fmt.Errorf("failed to createDirs.serviceName: %v", errCreateDirs))
	}

	errCreateDirs = os.Mkdir(filepath.Join(outputDir, serviceName, "bin"), 0754)
	if errCreateDirs != nil {
		errListCreateDirs = append(errListCreateDirs, fmt.Errorf("failed to createDirs.serviceName.bin: %v", errCreateDirs))
	}

	errCreateDirs = os.Mkdir(filepath.Join(outputDir, serviceName, "cmd"), 0754)
	if errCreateDirs != nil {
		errListCreateDirs = append(errListCreateDirs, fmt.Errorf("failed to createDirs.serviceName.cmd: %v", errCreateDirs))
	}

	errCreateDirs = os.Mkdir(filepath.Join(outputDir, serviceName, "definition"), 0754)
	if errCreateDirs != nil {
		errListCreateDirs = append(errListCreateDirs, fmt.Errorf("failed to createDirs.serviceName.definition: %v", errCreateDirs))
	}

	errCreateDirs = os.Mkdir(filepath.Join(outputDir, serviceName, "logger"), 0754)
	if errCreateDirs != nil {
		errListCreateDirs = append(errListCreateDirs, fmt.Errorf("failed to createDirs.serviceName.logger: %v", errCreateDirs))
	}

	errCreateDirs = os.Mkdir(filepath.Join(outputDir, serviceName, "server"), 0754)
	if errCreateDirs != nil {
		errListCreateDirs = append(errListCreateDirs, fmt.Errorf("failed to createDirs.serviceName.server: %v", errCreateDirs))
	}

	errCreateDirs = os.Mkdir(filepath.Join(outputDir, serviceName, "server/handler"), 0754)
	if errCreateDirs != nil {
		errListCreateDirs = append(errListCreateDirs, fmt.Errorf("failed to createDirs.serviceName.server.handler: %v", errCreateDirs))
	}

	errCreateDirs = os.Mkdir(filepath.Join(outputDir, serviceName, "server/conf"), 0754)
	if errCreateDirs != nil {
		errListCreateDirs = append(errListCreateDirs, fmt.Errorf("failed to createDirs.serviceName.server.conf: %v", errCreateDirs))
	}

	return nil
}
