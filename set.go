package flagx

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

type SetOfInt map[int]struct{}

// Set is flag.Value.Set
func (si *SetOfInt) Set(v string) error {
	ints := make(map[int]struct{})
	for _, v := range strings.Split(v, ",") {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		ints[int(i)] = struct{}{}
	}
	*si = SetOfInt(ints)
	return nil
}

func (si *SetOfInt) String() string {
	if si == nil {
		return ""
	}
	res := make([]string, 0, len(*si))
	for v := range *si {
		res = append(res, strconv.FormatInt(int64(v), 10))
	}
	sort.Strings(res)
	return strings.Join(res, ",")
}

type SetOfInt64 map[int64]struct{}

// Set is flag.Value.Set
func (si *SetOfInt64) Set(v string) error {
	ints := make(map[int64]struct{})
	for _, v := range strings.Split(v, ",") {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		ints[i] = struct{}{}
	}
	*si = SetOfInt64(ints)
	return nil
}

func (si *SetOfInt64) String() string {
	if si == nil {
		return ""
	}
	res := make([]string, 0, len(*si))
	for v := range *si {
		res = append(res, strconv.FormatInt(v, 10))
	}
	sort.Strings(res)
	return strings.Join(res, ",")
}

type SetOfUint map[uint]struct{}

// Set is flag.Value.Set
func (su *SetOfUint) Set(v string) error {
	ints := make(map[uint]struct{})
	for _, v := range strings.Split(v, ",") {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		ints[uint(i)] = struct{}{}
	}
	*su = SetOfUint(ints)
	return nil
}

func (su *SetOfUint) String() string {
	if su == nil {
		return ""
	}
	res := make([]string, 0, len(*su))
	for v := range *su {
		res = append(res, strconv.FormatUint(uint64(v), 10))
	}
	sort.Strings(res)
	return strings.Join(res, ",")
}

type SetOfUint64 map[uint64]struct{}

// Set is flag.Value.Set
func (su *SetOfUint64) Set(v string) error {
	ints := make(map[uint64]struct{})
	for _, v := range strings.Split(v, ",") {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		ints[i] = struct{}{}
	}
	*su = SetOfUint64(ints)
	return nil
}

func (su *SetOfUint64) String() string {
	if su == nil {
		return ""
	}
	res := make([]string, 0, len(*su))
	for v := range *su {
		res = append(res, strconv.FormatUint(v, 10))
	}
	sort.Strings(res)
	return strings.Join(res, ",")
}

type SetOfString map[string]struct{}

// Set is flag.Value.Set
func (ss *SetOfString) Set(v string) error {
	// TODO: how to configure separator? , \t |
	return nil
}

func (ss *SetOfString) String() string {
	if ss == nil {
		return ""
	}
	res := make([]string, 0, len(*ss))
	for v := range *ss {
		res = append(res, v)
	}
	sort.Strings(res)
	return strings.Join(res, ",")
}

type SetOfFloat64 map[float64]struct{}

// Set is flag.Value.Set
func (sf *SetOfFloat64) Set(v string) error {
	floats := make(map[float64]struct{})
	for _, v := range strings.Split(v, ",") {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		floats[f] = struct{}{}
	}
	*sf = SetOfFloat64(floats)
	return nil
}

func (sf *SetOfFloat64) String() string {
	if sf == nil {
		return ""
	}
	res := make([]string, 0, len(*sf))
	for v := range *sf {
		res = append(res, strconv.FormatFloat(v, 'g', 10, 64))
	}
	sort.Strings(res)
	return strings.Join(res, ",")
}

type SetOfDuration map[time.Duration]struct{}

// Set is flag.Value.Set
func (sd *SetOfDuration) Set(v string) error {
	durs := make(map[time.Duration]struct{})
	for _, v := range strings.Split(v, ",") {
		d, err := time.ParseDuration(v)
		if err != nil {
			return err
		}
		durs[d] = struct{}{}
	}
	*sd = SetOfDuration(durs)
	return nil
}

func (sd *SetOfDuration) String() string {
	if sd == nil {
		return ""
	}
	res := make([]string, 0, len(*sd))
	for v := range *sd {
		res = append(res, v.String())
	}
	sort.Strings(res)
	return strings.Join(res, ",")
}
