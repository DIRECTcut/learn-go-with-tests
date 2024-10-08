package countdown

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Countdown(os.Stdout)
}

func Countdown(out io.Writer) {
	fmt.Printf("Third: %p\n", out)
	fmt.Printf("fourth: %p\n", &out)

	fmt.Fprintf(out, "3")
}
