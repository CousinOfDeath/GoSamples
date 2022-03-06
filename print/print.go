package print

import (
	"fmt"
	"io"
)

func PrintLine(message string, writer io.Writer) {
	fmt.Fprintln(writer, message)
}
