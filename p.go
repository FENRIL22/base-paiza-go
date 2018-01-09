package main

import (
	"fmt"
	//"bufio"
	"os"
	"io"
)

func main() {

	err := Question(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println("[ERR]", err)
	}
}

func Question(in io.Reader, out io.Writer) error {
	var hoge string
	var fuga string

	fmt.Fscan(in, &hoge)
	fuga += hoge;
	fmt.Fscan(in, &hoge)
	fuga += "\n" + hoge;

	fmt.Fprintln(out, fuga)
	return nil
}
