package flagx

import (
	"testing"
	"time"
)

func TestSliceBool(t *testing.T) {
	mustEqual(t, (*SliceOfBool)(nil).String(), "")

	var sb SliceOfBool
	err := sb.Set("true,false,t,f,y,n,yes,no")
	failIfErr(t, err)

	want := []bool{true, false, true, false, true, false, true, false}
	mustEqual(t, []bool(sb), want)

	str := sb.String()
	wantStr := "true,false,true,false,true,false,true,false"
	mustEqual(t, str, wantStr)
}

func TestSliceBool_Bad(t *testing.T) {
	var sb SliceOfBool
	err := sb.Set("true,false,nono,yes")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSliceInt(t *testing.T) {
	mustEqual(t, (*SliceOfInt)(nil).String(), "")

	var sb SliceOfInt
	err := sb.Set("1,2,-3,4")
	failIfErr(t, err)

	want := []int{1, 2, -3, 4}
	mustEqual(t, []int(sb), want)

	str := sb.String()
	wantStr := "1,2,-3,4"
	mustEqual(t, str, wantStr)
}

func TestSliceInt_Bad(t *testing.T) {
	var sb SliceOfInt
	err := sb.Set("1,2,3.3,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSliceInt64(t *testing.T) {
	mustEqual(t, (*SliceOfInt64)(nil).String(), "")

	var sb SliceOfInt64
	err := sb.Set("1,2,-3,4")
	failIfErr(t, err)

	want := []int64{1, 2, -3, 4}
	mustEqual(t, []int64(sb), want)

	str := sb.String()
	wantStr := "1,2,-3,4"
	mustEqual(t, str, wantStr)
}

func TestSliceInt64_Bad(t *testing.T) {
	var sb SliceOfInt64
	err := sb.Set("1,2,3.3,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSliceUint(t *testing.T) {
	mustEqual(t, (*SliceOfUint)(nil).String(), "")

	var sb SliceOfUint
	err := sb.Set("1,2,3,4")
	failIfErr(t, err)

	want := []uint{1, 2, 3, 4}
	mustEqual(t, []uint(sb), want)

	str := sb.String()
	wantStr := "1,2,3,4"
	mustEqual(t, str, wantStr)
}

func TestSliceUint_Bad(t *testing.T) {
	var sb SliceOfUint
	err := sb.Set("1,2,-3,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSliceUint64(t *testing.T) {
	mustEqual(t, (*SliceOfUint64)(nil).String(), "")

	var sb SliceOfUint64
	err := sb.Set("1,2,3,4")
	failIfErr(t, err)

	want := []uint64{1, 2, 3, 4}
	mustEqual(t, []uint64(sb), want)

	str := sb.String()
	wantStr := "1,2,3,4"
	mustEqual(t, str, wantStr)
}

func TestSliceUint64_Bad(t *testing.T) {
	var sb SliceOfUint64
	err := sb.Set("1,2,-3,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSliceString(t *testing.T) {
}

func TestSliceString_Bad(t *testing.T) {
}

func TestSliceFloat64(t *testing.T) {
	mustEqual(t, (*SliceOfFloat64)(nil).String(), "")

	var sb SliceOfFloat64
	err := sb.Set("1,2e20,3.78,-4.20")
	failIfErr(t, err)

	want := []float64{1, 2e20, 3.78, -4.20}
	mustEqual(t, []float64(sb), want)

	str := sb.String()
	wantStr := "1,2e+20,3.78,-4.2"
	mustEqual(t, str, wantStr)
}

func TestSliceFloat64_Bad(t *testing.T) {
	var sb SliceOfFloat64
	err := sb.Set("1,2,3/2,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSliceDuration(t *testing.T) {
	mustEqual(t, (*SliceOfDuration)(nil).String(), "")

	var sb SliceOfDuration
	err := sb.Set("1ns,1s,-1h")
	failIfErr(t, err)

	want := []time.Duration{1, time.Second, -time.Hour}
	mustEqual(t, []time.Duration(sb), want)

	str := sb.String()
	wantStr := "1ns,1s,-1h0m0s"
	mustEqual(t, str, wantStr)
}

func TestSliceDuration_Bad(t *testing.T) {
	var sb SliceOfDuration
	err := sb.Set("1s,2day")
	if err == nil {
		t.Fatal(err)
	}
}
