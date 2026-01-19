package main

// uses comparable keyword to make sure T is a comparable type (constraint)
func findIndex[T comparable](input_slice[]T,value T) int{
	for i,v := range input_slice {
		if value == v{
			return i
		}
	}
	return 0
}