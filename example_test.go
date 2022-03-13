package flagx_test

import (
	"os"
	"testing"
	"time"

	"github.com/cristalhq/flagx"
)

func ExampleFlagSet(t *testing.T) {
	var d time.Duration
	fset := flagx.NewFlagSet("testing", os.Stderr)
	fset.Duration(&d, "timeout", "t", 10*time.Second, "just a timeout")

	err := fset.Parse([]string{"-t", "20s"})
	if err != nil {
		t.Fatal(err)
	}

	if d != 20*time.Second {
		t.Fatalf("got %v want %v", d, 20*time.Second)
	}
}
