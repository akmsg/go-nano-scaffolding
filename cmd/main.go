package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/akmsg/go-nano-scaffolding/options"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
		os.Exit(127)
	}
	options.Run()
}
