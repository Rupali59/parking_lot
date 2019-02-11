//Package helper provides helpers to read commands
package helpers

import (
	commandProcessor "../processor"
	"bufio"
	"fmt"
	"os"
)

//FileProcessor reads a file to get the commands to be executed
type FileProcessor struct {
	FileName string //Note FileName is the fully qualified name of the path where the commands are stored
}

//Process reads the file given its fully qualified path and executes the commands present in the file and exits
//after reading the file
func (fileProcessor *FileProcessor) Process() (err error) {
	file, err := os.Open(fileProcessor.FileName)

	if err != nil {
		return err
	}

	defer file.Close()
	reader := bufio.NewScanner(file)
	processor := commandProcessor.CommandProcessor{}
	for reader.Scan() {
		command := reader.Text()
		output, cmd_err := processor.Process(command)
		if cmd_err != nil {
			fmt.Println(cmd_err)
		} else {
			fmt.Println(output)
		}
	}
	if err := reader.Err(); err != nil {
		os.Exit(1)
		return err
	}

	os.Exit(0)
	return nil
}
