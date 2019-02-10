package cmd

import (
	"os"
	helpers "../internal/app/helpers"
)

func main()  {
	if(len(os.Args)>1 && os.Args[1]!=""){
		helpers.FileProcessor{os.Args[1]}.Process();
	}
	helpers.CommandLineHelper.Process();
}