package main

// Example for variadic parameters
func Sum(numbers ... int) int{
	mySum := 0 
	for _,num := range(numbers) {
		mySum += num
	}
	return mySum
}