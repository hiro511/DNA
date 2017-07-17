package main

import (
	"fmt"

	"github.com/hiro511/DNA/args"
)

func main() {

	if err := args.Parse(); err != nil {
		fmt.Println(err.Error())
		args.Usage()
		return
	}

	args.Debug()
	// _, err := ioutil.ReadFile(args.FileName)
	// if err != nil {
	// 	fmt.Println("failed to read the specified file:", args.FileName)
	// 	return
	// }
}
