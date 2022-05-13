package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		data, _ := io.ReadAll(os.Stdin)

		PrintString(string(data))
	}
	for _, v := range args {
		file, err := ioutil.ReadFile(v)
		if err == nil {
			PrintString(string(file))
		} else {
			PrintString("ERROR: open " + v + ": no such file or directory")
			PrintString("\n")
			os.Exit(1)
		}
	}
}

func PrintString(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}
