package flagx

import (
	"testing"
	"time"
)

func TestSetInt(t *testing.T) {
	mustEqual(t, (*SetOfInt)(nil).String(), "")

	var sb SetOfInt
	err := sb.Set("1,2,-3,4")
	failIfErr(t, err)

	want := map[int]struct{}{1: {}, 2: {}, -3: {}, 4: {}}
	mustEqual(t, map[int]struct{}(sb), want)

	str := sb.String()
	wantStr := "-3,1,2,4"
	mustEqual(t, str, wantStr)
}

func TestSetInt_Bad(t *testing.T) {
	var sb SetOfInt
	err := sb.Set("1,2,3.3,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSetInt64(t *testing.T) {
	mustEqual(t, (*SetOfInt64)(nil).String(), "")

	var sb SetOfInt64
	err := sb.Set("1,2,-3,4")
	failIfErr(t, err)

	want := map[int64]struct{}{1: {}, 2: {}, -3: {}, 4: {}}
	mustEqual(t, map[int64]struct{}(sb), want)

	str := sb.String()
	wantStr := "-3,1,2,4"
	mustEqual(t, str, wantStr)
}

func TestSetInt64_Bad(t *testing.T) {
	var sb SetOfInt64
	err := sb.Set("1,2,3.3,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSetUint(t *testing.T) {
	mustEqual(t, (*SetOfUint)(nil).String(), "")

	var sb SetOfUint
	err := sb.Set("1,2,3,4")
	failIfErr(t, err)

	want := map[uint]struct{}{1: {}, 2: {}, 3: {}, 4: {}}
	mustEqual(t, map[uint]struct{}(sb), want)

	str := sb.String()
	wantStr := "1,2,3,4"
	mustEqual(t, str, wantStr)
}

func TestSetUint_Bad(t *testing.T) {
	var sb SetOfUint
	err := sb.Set("1,2,-3,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSetUint64(t *testing.T) {
	mustEqual(t, (*SetOfUint64)(nil).String(), "")

	var sb SetOfUint64
	err := sb.Set("1,2,3,4")
	failIfErr(t, err)

	want := map[uint64]struct{}{1: {}, 2: {}, 3: {}, 4: {}}
	mustEqual(t, map[uint64]struct{}(sb), want)

	str := sb.String()
	wantStr := "1,2,3,4"
	mustEqual(t, str, wantStr)
}

func TestSetUint64_Bad(t *testing.T) {
	var sb SetOfUint64
	err := sb.Set("1,2,-3,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSetString(t *testing.T) {
}

func TestSetString_Bad(t *testing.T) {
}

func TestSetFloat64(t *testing.T) {
	mustEqual(t, (*SetOfFloat64)(nil).String(), "")

	var sb SetOfFloat64
	err := sb.Set("1,2e20,3.78,-4.20")
	failIfErr(t, err)

	want := map[float64]struct{}{1: {}, 2e20: {}, 3.78: {}, -4.20: {}}
	mustEqual(t, map[float64]struct{}(sb), want)

	str := sb.String()
	wantStr := "-4.2,1,2e+20,3.78"
	mustEqual(t, str, wantStr)
}

func TestSetFloat64_Bad(t *testing.T) {
	var sb SetOfFloat64
	err := sb.Set("1,2,3/2,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSetDuration(t *testing.T) {
	mustEqual(t, (*SetOfDuration)(nil).String(), "")

	var sb SetOfDuration
	err := sb.Set("1ns,1s,-1h")
	failIfErr(t, err)

	want := map[time.Duration]struct{}{1: {}, time.Second: {}, -time.Hour: {}}
	mustEqual(t, map[time.Duration]struct{}(sb), want)

	str := sb.String()
	wantStr := "-1h0m0s,1ns,1s"
	mustEqual(t, str, wantStr)
}

func TestSetDuration_Bad(t *testing.T) {
	var sb SetOfDuration
	err := sb.Set("1s,2day")
	if err == nil {
		t.Fatal(err)
	}
}
