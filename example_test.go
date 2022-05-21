package flagx_test

import (
	"fmt"
	"os"
	"time"

	"github.com/cristalhq/flagx"
)

func ExampleFlagSet() {
	var d time.Duration
	fset := flagx.NewFlagSet("testing", os.Stderr)
	fset.Duration(&d, "timeout", "t", 10*time.Second, "just a timeout")

	err := fset.Parse([]string{"-t", "20s"})
	if err != nil {
		panic(err)
	}

	if d != 20*time.Second {
		panic(fmt.Sprintf("got %v want %v", d, 20*time.Second))
	}
}
