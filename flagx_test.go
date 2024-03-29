package flagx

import (
	"bytes"
	"flag"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestFlagSet(t *testing.T) {
	var d time.Duration
	var ids []int
	wantIDs := []int{1, 2, 3}
	var offsets map[float64]struct{}
	wantOffsets := map[float64]struct{}{1: {}, 2: {}, 3: {}}

	fset := NewFlagSet("testing", os.Stderr)
	fset.Duration(&d, "timeout", "t", 10*time.Second, "just a timeout")
	fset.IntSlice(&ids, "ids", "", wantIDs, ",", "just a timeout")
	fset.Float64Set(&offsets, "offsets", "", wantOffsets, "just a timeout")

	names := map[string]struct{}{}
	fs := fset.AsStdlib()
	fs.VisitAll(func(f *flag.Flag) {
		names[f.Name] = struct{}{}
	})

	want := map[string]struct{}{
		"timeout": {}, "t": {}, "ids": {}, "offsets": {},
	}
	mustEqual(t, names, want)
	mustEqual(t, ids, wantIDs)
	mustEqual(t, offsets, wantOffsets)
}

func TestFlagSet_PrintDefaults(t *testing.T) {
	const usage = `  -timeout (-t) duration
    	just a timeout (default 10s)
`
	var buf bytes.Buffer
	fset := NewFlagSet("testing", &buf)
	fset.Duration(new(time.Duration), "timeout", "t", 10*time.Second, "just a timeout")
	fset.PrintDefaults()
	mustEqual(t, buf.String(), usage)
}

func failIfErr(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatal(err)
	}
}

func mustEqual(tb testing.TB, have, want interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(have, want) {
		tb.Fatalf("\nhave: %v\nwant: %v", have, want)
	}
}
