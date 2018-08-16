package goptionals

// Bool returns a pointer to the bool value passed in.
func Bool(v bool) *bool {
	return &v
}

// BoolValue returns the value of the bool pointer passed in or false if the pointer is nil.
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// BoolSlice converts a slice of bool values into a slice of bool pointers
func BoolSlice(src []bool) []*bool {
	dst := make([]*bool, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// BoolValueSlice converts a slice of bool pointers into a slice of bool values
func BoolValueSlice(src []*bool) []bool {
	dst := make([]bool, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// BoolMap converts a string map of bool values into a string map of bool pointers
func BoolMap(src map[string]bool) map[string]*bool {
	dst := make(map[string]*bool, len(src))
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// BoolValueMap converts a string map of bool pointers into a string map of bool values
func BoolValueMap(src map[string]*bool) map[string]bool {
	dst := make(map[string]bool, len(src))
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}
