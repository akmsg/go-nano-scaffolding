package options

import (
	"flag"
	"fmt"
	"os"

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
	fmt.Printf("Will scaffold service: %s at: %s\n", *name, *out)

	errorCrateService := scaffolding.CreateService(*name, *out)
	if nil != errorCrateService {
		fmt.Printf("Failed to create service: %v", errorCrateService)
		os.Exit(1)
	}

	fmt.Printf("Scaffolding completed!\n")
}
