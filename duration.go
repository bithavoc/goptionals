package goptionals

import "time"

// Duration returns a pointer to the time.Duration in.
func Duration(v time.Duration) *time.Duration {
	return &v
}

// DurationValue returns the value of the time.Duration pointer passed in or time.Duration(0) if the pointer is nil.
func DurationValue(v *time.Duration) time.Duration {
	if v != nil {
		return *v
	}
	return time.Duration(0)
}

// DurationSlice converts a slice of time.Duration values into a slice of time.Duration pointers
func DurationSlice(src []time.Duration) []*time.Duration {
	dst := make([]*time.Duration, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// DurationValueSlice converts a slice of time.Duration pointers into a slice of time.Duration values
func DurationValueSlice(src []*time.Duration) []time.Duration {
	dst := make([]time.Duration, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// DurationMap converts a string map of time.Duration values into a string
// map of time.Duration pointers
func DurationMap(src map[string]time.Duration) map[string]*time.Duration {
	dst := make(map[string]*time.Duration, len(src))
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// DurationValueMap converts a string map of time.Duration pointers into a string
// map of time.Duration values
func DurationValueMap(src map[string]*time.Duration) map[string]time.Duration {
	dst := make(map[string]time.Duration, len(src))
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}
