package flagx

import (
	"flag"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestFlagSet(t *testing.T) {
	var ids []int
	wantIDs := []int{1, 2, 3}

	fset := NewFlagSet("testing", os.Stderr)
	_ = fset.Duration("timeout", "t", 10*time.Second, "just a timeout")
	fset.IntSliceVar(&ids, "ids", "", wantIDs, "just a timeout")

	names := map[string]struct{}{}
	fs := fset.AsStdlib()
	fs.VisitAll(func(f *flag.Flag) {
		names[f.Name] = struct{}{}
	})

	want := map[string]struct{}{
		"timeout": {}, "t": {}, "ids": {},
	}
	if !reflect.DeepEqual(names, want) {
		t.Fatalf("got %v want %v", names, want)
	}

	if !reflect.DeepEqual(ids, wantIDs) {
		t.Fatalf("got %v want %v", want, wantIDs)
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
