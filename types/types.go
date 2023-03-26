package types

// PointerVal gets value of the pointer or zero value in case of nil pointer.
func PointerVal[T any](ptr *T) T {
	if ptr == nil {
		return *new(T)
	}

	return *ptr
}
