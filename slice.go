package flagx

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type boolSlice struct {
	sep   string
	value *[]bool
}

// Set implements the flag.Value interface.
func (s boolSlice) Set(str string) error {
	var bools []bool
	for _, v := range strings.Split(str, s.sep) {
		b, err := strconv.ParseBool(v)
		if err != nil {
			return fmt.Errorf("parsing bool: %w", err)
		}
		bools = append(bools, b)
	}
	*s.value = bools
	return nil
}

// String implements the flag.Value interface.
func (s boolSlice) String() string {
	res := make([]string, len(*s.value))
	for i, v := range *s.value {
		res[i] = strconv.FormatBool(v)
	}
	return strings.Join(res, s.sep)
}

type intSlice struct {
	sep   string
	value *[]int
}

// Set implements the flag.Value interface.
func (s intSlice) Set(str string) error {
	var ints []int
	for _, v := range strings.Split(str, s.sep) {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return fmt.Errorf("parsing int: %w", err)
		}
		ints = append(ints, int(i))
	}
	*s.value = ints
	return nil
}

// String implements the flag.Value interface.
func (s intSlice) String() string {
	res := make([]string, len(*s.value))
	for i, v := range *s.value {
		res[i] = strconv.FormatInt(int64(v), 10)
	}
	return strings.Join(res, s.sep)
}

type int64Slice struct {
	sep   string
	value *[]int64
}

// Set implements the flag.Value interface.
func (s int64Slice) Set(str string) error {
	var ints []int64
	for _, v := range strings.Split(str, s.sep) {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return fmt.Errorf("parsing int: %w", err)
		}
		ints = append(ints, i)
	}
	*s.value = ints
	return nil
}

// String implements the flag.Value interface.
func (s int64Slice) String() string {
	res := make([]string, len(*s.value))
	for i, v := range *s.value {
		res[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(res, s.sep)
}

type uintSlice struct {
	sep   string
	value *[]uint
}

// Set implements the flag.Value interface.
func (s uintSlice) Set(str string) error {
	var uints []uint
	for _, v := range strings.Split(str, s.sep) {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return fmt.Errorf("parsing uint: %w", err)
		}
		uints = append(uints, uint(i))
	}
	*s.value = uints
	return nil
}

// String implements the flag.Value interface.
func (s uintSlice) String() string {
	res := make([]string, len(*s.value))
	for i, v := range *s.value {
		res[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(res, s.sep)
}

type uint64Slice struct {
	sep   string
	value *[]uint64
}

// Set implements the flag.Value interface.
func (s uint64Slice) Set(str string) error {
	var uints []uint64
	for _, v := range strings.Split(str, s.sep) {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return fmt.Errorf("parsing uint: %w", err)
		}
		uints = append(uints, i)
	}
	*s.value = uints
	return nil
}

// String implements the flag.Value interface.
func (s uint64Slice) String() string {
	res := make([]string, len(*s.value))
	for i, v := range *s.value {
		res[i] = strconv.FormatUint(v, 10)
	}
	return strings.Join(res, s.sep)
}

type stringSlice struct {
	sep   string
	value *[]string
}

// Set implements the flag.Value interface.
func (s stringSlice) Set(str string) error {
	*s.value = strings.Split(str, s.sep)
	return nil
}

// String implements the flag.Value interface.
func (s stringSlice) String() string {
	return strings.Join(*s.value, s.sep)
}

type float64Slice struct {
	sep   string
	value *[]float64
}

// Set implements the flag.Value interface.
func (s float64Slice) Set(str string) error {
	var floats []float64
	for _, v := range strings.Split(str, s.sep) {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return fmt.Errorf("parsing float: %w", err)
		}
		floats = append(floats, f)
	}
	*s.value = floats
	return nil
}

// String implements the flag.Value interface.
func (s float64Slice) String() string {
	res := make([]string, len(*s.value))
	for i, v := range *s.value {
		res[i] = strconv.FormatFloat(v, 'g', 10, 64)
	}
	return strings.Join(res, s.sep)
}

type durationSlice struct {
	sep   string
	value *[]time.Duration
}

// Set implements the flag.Value interface.
func (s durationSlice) Set(str string) error {
	var durs []time.Duration
	for _, v := range strings.Split(str, s.sep) {
		d, err := time.ParseDuration(v)
		if err != nil {
			return fmt.Errorf("parsing duration: %w", err)
		}
		durs = append(durs, d)
	}
	*s.value = durs
	return nil
}

// String implements the flag.Value interface.
func (s durationSlice) String() string {
	res := make([]string, len(*s.value))
	for i, v := range *s.value {
		res[i] = v.String()
	}
	return strings.Join(res, s.sep)
}
