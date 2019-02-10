package helpers

import (
	"bufio"
	"fmt"
	"os"
	commandProcessor "../processor"
)

type FileProcessor struct {
 fileName string
}

func (fileProcessor *FileProcessor)Process() (err error){
	file,err:=os.Open(fileProcessor.fileName)
	if err!=nil{
		return err
	}

	defer file.Close()
	reader := bufio.NewScanner(file)
	processor:= commandProcessor.CommandProcessor{}
	for reader.Scan(){
		command := reader.Text()
		output,cmd_err:= processor.Process(command)
		if cmd_err !=nil{
			fmt.Println(cmd_err)
		}else{
			fmt.Println(output)
		}
	}

	if err:=reader.Err(); err!=nil{
		return err
	}

	return nil
}
