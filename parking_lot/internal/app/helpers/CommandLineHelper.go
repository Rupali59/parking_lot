package helpers

import (
	commandProcessor "../processor"
	"bufio"
	"fmt"
	"os"
)

type CommandLineHelper struct {
}

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
