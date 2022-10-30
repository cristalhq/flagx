package flagx

import (
	"testing"
	"time"
)

const sep = ","

func TestSliceBool(t *testing.T) {
	s := boolSlice{sep: sep, value: new([]bool)}
	err := s.Set("true,false,t,f,1,0")
	failIfErr(t, err)

	want := []bool{true, false, true, false, true, false}
	mustEqual(t, *s.value, want)

	str := s.String()
	wantStr := "true,false,true,false,true,false"
	mustEqual(t, str, wantStr)
}

func TestSliceBool_Bad(t *testing.T) {
	s := boolSlice{sep: sep, value: new([]bool)}
	err := s.Set("true,false,nono,yes")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSliceInt(t *testing.T) {
	s := intSlice{sep: sep, value: new([]int)}
	err := s.Set("1,2,-3,4")
	failIfErr(t, err)

	want := []int{1, 2, -3, 4}
	mustEqual(t, *s.value, want)

	str := s.String()
	wantStr := "1,2,-3,4"
	mustEqual(t, str, wantStr)
}

func TestSliceInt_Bad(t *testing.T) {
	s := intSlice{sep: sep, value: new([]int)}
	err := s.Set("1,2,3.3,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSliceInt64(t *testing.T) {
	s := int64Slice{sep: sep, value: new([]int64)}
	err := s.Set("1,2,-3,4")
	failIfErr(t, err)

	want := []int64{1, 2, -3, 4}
	mustEqual(t, *s.value, want)

	str := s.String()
	wantStr := "1,2,-3,4"
	mustEqual(t, str, wantStr)
}

func TestSliceInt64_Bad(t *testing.T) {
	s := int64Slice{sep: sep, value: new([]int64)}
	err := s.Set("1,2,3.3,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSliceUint(t *testing.T) {
	s := uintSlice{sep: sep, value: new([]uint)}
	err := s.Set("1,2,3,4")
	failIfErr(t, err)

	want := []uint{1, 2, 3, 4}
	mustEqual(t, *s.value, want)

	str := s.String()
	wantStr := "1,2,3,4"
	mustEqual(t, str, wantStr)
}

func TestSliceUint_Bad(t *testing.T) {
	s := uintSlice{sep: sep, value: new([]uint)}
	err := s.Set("1,2,-3,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSliceUint64(t *testing.T) {
	s := uint64Slice{sep: sep, value: new([]uint64)}
	err := s.Set("1,2,3,4")
	failIfErr(t, err)

	want := []uint64{1, 2, 3, 4}
	mustEqual(t, *s.value, want)

	str := s.String()
	wantStr := "1,2,3,4"
	mustEqual(t, str, wantStr)
}

func TestSliceUint64_Bad(t *testing.T) {
	s := uint64Slice{sep: sep, value: new([]uint64)}
	err := s.Set("1,2,-3,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSliceString(t *testing.T) {
	s := stringSlice{sep: sep, value: new([]string)}
	err := s.Set("1,2e20,3.78,rabbit,-4.20")
	failIfErr(t, err)

	want := []string{"1", "2e20", "3.78", "rabbit", "-4.20"}
	mustEqual(t, *s.value, want)

	str := s.String()
	wantStr := "1,2e20,3.78,rabbit,-4.20"
	mustEqual(t, str, wantStr)
}

func TestSliceFloat64(t *testing.T) {
	s := float64Slice{sep: sep, value: new([]float64)}
	err := s.Set("1,2e20,3.78,-4.20")
	failIfErr(t, err)

	want := []float64{1, 2e20, 3.78, -4.20}
	mustEqual(t, *s.value, want)

	str := s.String()
	wantStr := "1,2e+20,3.78,-4.2"
	mustEqual(t, str, wantStr)
}

func TestSliceFloat64_Bad(t *testing.T) {
	s := float64Slice{sep: sep, value: new([]float64)}
	err := s.Set("1,2,3/2,4")
	if err == nil {
		t.Fatal(err)
	}
}

func TestSliceDuration(t *testing.T) {
	s := durationSlice{sep: sep, value: new([]time.Duration)}
	err := s.Set("1ns,1s,-1h")
	failIfErr(t, err)

	want := []time.Duration{1, time.Second, -time.Hour}
	mustEqual(t, *s.value, want)

	str := s.String()
	wantStr := "1ns,1s,-1h0m0s"
	mustEqual(t, str, wantStr)
}

func TestSliceDuration_Bad(t *testing.T) {
	s := durationSlice{sep: sep, value: new([]time.Duration)}
	err := s.Set("1s,2day")
	if err == nil {
		t.Fatal(err)
	}
}
