package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countDownStart = 3

func main() {
	CountDown(os.Stdout)
}

func CountDown(out io.Writer) { // bytes.Buffer implements io.Writer
	for i := countDownStart; i >= 1; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord) // Fprintf vs Fprint ??
}
