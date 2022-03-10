package flagx

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type SliceOfBool []bool

// Set is flag.Value.Set
func (sb *SliceOfBool) Set(v string) error {
	var bools []bool
	for i, v := range strings.Split(v, ",") {
		if v == "true" || v == "t" || v == "yes" || v == "y" {
			bools = append(bools, true)
		} else if v == "false" || v == "f" || v == "no" || v == "n" {
			bools = append(bools, false)
		} else {
			return fmt.Errorf("unknown bool at %d: %s", i, v)
		}
	}
	*sb = SliceOfBool(bools)
	return nil
}

func (sb *SliceOfBool) String() string {
	if sb == nil {
		return ""
	}
	res := make([]string, len(*sb))
	for i, v := range []bool(*sb) {
		if v {
			res[i] = "true"
		} else {
			res[i] = "false"
		}
	}
	return strings.Join(res, ",")
}

type SliceOfInt []int

// Set is flag.Value.Set
func (si *SliceOfInt) Set(v string) error {
	var ints []int
	for _, v := range strings.Split(v, ",") {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		ints = append(ints, int(i))
	}
	*si = SliceOfInt(ints)
	return nil
}

func (si *SliceOfInt) String() string {
	if si == nil {
		return ""
	}
	res := make([]string, len(*si))
	for i, v := range *si {
		res[i] = strconv.FormatInt(int64(v), 10)
	}
	return strings.Join(res, ",")
}

type SliceOfInt64 []int64

// Set is flag.Value.Set
func (si *SliceOfInt64) Set(v string) error {
	var ints []int64
	for _, v := range strings.Split(v, ",") {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		ints = append(ints, i)
	}
	*si = SliceOfInt64(ints)
	return nil
}

func (si *SliceOfInt64) String() string {
	if si == nil {
		return ""
	}
	res := make([]string, len(*si))
	for i, v := range *si {
		res[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(res, ",")
}

type SliceOfUint []uint

// Set is flag.Value.Set
func (su *SliceOfUint) Set(v string) error {
	var ints []uint
	for _, v := range strings.Split(v, ",") {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		ints = append(ints, uint(i))
	}
	*su = SliceOfUint(ints)
	return nil
}

func (su *SliceOfUint) String() string {
	if su == nil {
		return ""
	}
	res := make([]string, len(*su))
	for i, v := range *su {
		res[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(res, ",")
}

type SliceOfUint64 []uint64

// Set is flag.Value.Set
func (su *SliceOfUint64) Set(v string) error {
	var ints []uint64
	for _, v := range strings.Split(v, ",") {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		ints = append(ints, i)
	}
	*su = SliceOfUint64(ints)
	return nil
}

func (su *SliceOfUint64) String() string {
	if su == nil {
		return ""
	}
	res := make([]string, len(*su))
	for i, v := range *su {
		res[i] = strconv.FormatUint(v, 10)
	}
	return strings.Join(res, ",")
}

type SliceOfString []string

// Set is flag.Value.Set
func (ss *SliceOfString) Set(v string) error {
	// TODO: how to configure separator? , \t |
	return nil
}

func (ss *SliceOfString) String() string {
	if ss == nil {
		return ""
	}
	return strings.Join([]string(*ss), ",")
}

type SliceOfFloat64 []float64

// Set is flag.Value.Set
func (sf *SliceOfFloat64) Set(v string) error {
	var floats []float64
	for _, v := range strings.Split(v, ",") {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		floats = append(floats, f)
	}
	*sf = SliceOfFloat64(floats)
	return nil
}

func (sf *SliceOfFloat64) String() string {
	if sf == nil {
		return ""
	}
	res := make([]string, len(*sf))
	for i, v := range *sf {
		res[i] = strconv.FormatFloat(v, 'g', 10, 64)
	}
	return strings.Join(res, ",")
}

type SliceOfDuration []time.Duration

// Set is flag.Value.Set
func (sd *SliceOfDuration) Set(v string) error {
	var durs []time.Duration
	for _, v := range strings.Split(v, ",") {
		d, err := time.ParseDuration(v)
		if err != nil {
			return err
		}
		durs = append(durs, d)
	}
	*sd = SliceOfDuration(durs)
	return nil
}

func (sd *SliceOfDuration) String() string {
	if sd == nil {
		return ""
	}
	res := make([]string, len(*sd))
	for i, v := range *sd {
		res[i] = v.String()
	}
	return strings.Join(res, ",")
}
