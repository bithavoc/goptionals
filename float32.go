package goptionals

// Float32 returns a pointer to the float32 value passed in.
func Float32(v float32) *float32 {
	return &v
}

// Float32Value returns the value of the float32 pointer passed in or
// 0 if the pointer is nil.
func Float32Value(v *float32) float32 {
	if v != nil {
		return *v
	}
	return 0
}

// Float32Slice converts a slice of float32 values into a slice of float32 pointers
func Float32Slice(src []float32) []*float32 {
	dst := make([]*float32, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Float32ValueSlice converts a slice of float32 pointers into a slice of float32 values
func Float32ValueSlice(src []*float32) []float32 {
	dst := make([]float32, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Float32Map converts a string map of float32 values into a string map of float32 pointers
func Float32Map(src map[string]float32) map[string]*float32 {
	dst := make(map[string]*float32, len(src))
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Float32ValueMap converts a string map of float32 pointers into a string map of float32 values
func Float32ValueMap(src map[string]*float32) map[string]float32 {
	dst := make(map[string]float32, len(src))
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}
