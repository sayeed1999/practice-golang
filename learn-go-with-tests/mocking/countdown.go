package countdown

import (
	"fmt"
	"io"
)

func CountDown(out io.Writer) { // bytes.Buffer implements io.Writer
	fmt.Fprintf(out, "3")
}
