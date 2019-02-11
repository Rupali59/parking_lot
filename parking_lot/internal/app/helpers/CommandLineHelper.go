//Package helper provides helpers to read commands
package helpers

import (
	commandProcessor "../processor"
	"bufio"
	"fmt"
	"os"
)

//CommandLineHelper helps reading the commands from the console
type CommandLineHelper struct {
}

//Process the command sent on the console. As soon as exit gets called, exits the program
func (helper *CommandLineHelper) Process() (err error) {
	reader := bufio.NewReader(os.Stdin)
	processor := commandProcessor.CommandProcessor{}
	for {
		command, _, _ := reader.ReadLine()
		if string(command) == "exit" {
			return nil
		}
		output, cmd_err := processor.Process(string(command))
		if cmd_err != nil {
			fmt.Println(cmd_err)
		} else {
			fmt.Println(output)
		}
	}
}
