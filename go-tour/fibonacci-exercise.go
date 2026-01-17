package main

func Fibo() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}
