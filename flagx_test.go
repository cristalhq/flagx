package flagx

import (
	"flag"
	"reflect"
	"testing"
	"time"
)

func TestFlagSet(t *testing.T) {
	fset := NewFlagSet("testing")
	_ = fset.Duration("timeout", "t", 10*time.Second, "just a timeout")

	names := map[string]struct{}{}
	fs := fset.AsStdlibFlagSet()
	fs.VisitAll(func(f *flag.Flag) {
		names[f.Name] = struct{}{}
	})

	want := map[string]struct{}{
		"timeout": {}, "t": {},
	}
	if !reflect.DeepEqual(names, want) {
		t.Fatalf("got %v want %v", names, want)
	}
}

func failIfErr(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func mustEqual(t testing.TB, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
