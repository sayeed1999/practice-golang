package main

import (
	"fmt"
	"io"
	"os"

	"countdown.go/sleeper"
)

const finalWord = "Go!"
const countDownStart = 3

func main() {
	CountDown(os.Stdout, &sleeper.RealSleeper{})
}

func CountDown(out io.Writer, sleeper sleeper.Sleeper) { // bytes.Buffer implements io.Writer,
	for i := countDownStart; i >= 1; i-- {
		fmt.Fprintln(out, i)
		// time.Sleep(1 * time.Second)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord) // Fprintf vs Fprint ??
}
