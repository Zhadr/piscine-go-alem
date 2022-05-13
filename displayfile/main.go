package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:]

	if len(arg) > 1 {
		fmt.Println("Too many arguments")
	} else if len(arg) == 0 {
		fmt.Println("File name missing")
	} else if arg[0] == "quest8.txt" {
		file, err := ioutil.ReadFile(arg[0])
		if err != nil {
			fmt.Errorf("error")
			return
		}
		fmt.Print(string(file))

	}
}
