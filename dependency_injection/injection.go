package dependency_injection

import (
	"io"
	"fmt"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
