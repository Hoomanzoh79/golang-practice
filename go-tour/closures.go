package main

// Adder is a factory function that returns a 
// closure which captures the sum variable
func Adder()func(int)int{
	sum := 0
	return func (x int)int{
		sum += x
		return sum
	}
}
