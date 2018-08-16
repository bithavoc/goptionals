package goptionals

// String returns a pointer to the string passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or "" when the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// StringSlice converts a slice of string values into a slice of string pointers
func StringSlice(src []string) []*string {
	dst := make([]*string, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// StringValueSlice converts a slice of string pointers into a slice of string values
func StringValueSlice(src []*string) []string {
	dst := make([]string, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// StringMap converts a string map of string values into a string map of string pointers
func StringMap(src map[string]string) map[string]*string {
	dst := make(map[string]*string, len(src))
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// StringValueMap converts a string map of string pointers into a string map of string values
func StringValueMap(src map[string]*string) map[string]string {
	dst := make(map[string]string, len(src))
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}
