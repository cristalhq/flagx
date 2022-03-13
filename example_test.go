package flagx_test

import (
	"os"
	"testing"
	"time"

	"github.com/cristalhq/flagx"
)

func ExampleFlagSet(t *testing.T) {
	fset := flagx.NewFlagSet("testing", os.Stderr)
	d := fset.Duration("timeout", "t", 10*time.Second, "just a timeout")

	err := fset.Parse([]string{"-t", "20s"})
	if err != nil {
		t.Fatal(err)
	}

	if d == nil {
		t.Fatal()
	}
	if val := *d; val != 20*time.Second {
		t.Fatalf("got %v want %v", val, 20*time.Second)
	}
}
