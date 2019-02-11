
package main

import (
	"../internal/app/helpers"
	"os"
)

//Main checks for the filename. If found, then the fileProcessor handles the file else commandlinehelper
//reads from the console
func main() {
	if len(os.Args) > 1 && os.Args[1] != "" {
		fileProcessor := helpers.FileProcessor{os.Args[1]}
		fileProcessor.Process()
	}
	commandLineHelper := helpers.CommandLineHelper{}
	commandLineHelper.Process()
}
