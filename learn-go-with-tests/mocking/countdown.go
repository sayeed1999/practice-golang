package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CountDown(os.Stdout)
}

func CountDown(out io.Writer) { // bytes.Buffer implements io.Writer
	fmt.Fprintf(out, "3")
}
