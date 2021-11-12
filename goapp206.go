package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, dept = "Trat", "IT"
	s := fmt.Sprintln(name,dept)
	io.WriteString(os.Stdout,s);
}
