
package main

import (
	"../internal/app/helpers"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] != "" {
		fileProcessor := helpers.FileProcessor{os.Args[1]}
		fileProcessor.Process()
	}
	commandLineHelper := helpers.CommandLineHelper{}
	commandLineHelper.Process()
}
