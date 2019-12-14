package options

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/akmsg/go-nano-scaffolding/scaffolding"
)

func Run() {
	name := flag.String("name", "", "provide <service-name> to scaffold a service")
	out := flag.String("out", "", "provide output directory, default is cwd")
	flag.Parse()

	if name == nil || 1 > len(*name) {
		flag.Usage()
		return
	}

	if out == nil || 1 > len(*out) {
		outputDir, errGetwd := os.Getwd()
		if errGetwd != nil {
			panic(fmt.Errorf("failed to get current working dir: %v", errGetwd))
		}
		out = &outputDir
	}
	servicePath := filepath.Join(*out, *name)
	fmt.Printf("Will scaffold service: %s at: %s\n", *name, servicePath)

	errorCrateService := scaffolding.CreateService(*name, *out)
	if nil != errorCrateService {
		fmt.Printf("Failed to create service: %v", errorCrateService)
		os.Exit(1)
	}

	fmt.Printf("Scaffolding done ...\n")
	errChdir := os.Chdir(servicePath)
	if nil != errChdir {
		fmt.Printf("Failed to create service: %v", errChdir)
		os.Exit(2)
	}

	c := exec.Command("go", "generate")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	errGoGen := c.Run()

	if nil != errGoGen {
		fmt.Printf("Failed to create service: %v", errGoGen)
		os.Exit(3)
	}

	fmt.Printf("Service created!\n")
}
