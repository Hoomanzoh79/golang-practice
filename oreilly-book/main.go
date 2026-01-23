package main

import (
	"errors"
	"fmt"
)

type Person struct{
	FirstName string
	LastName string
}

// Example for joining errors 
func (p Person) Validator() (Person,error) {
	var myErrors []error
	if len(p.FirstName) == 0 {
		myErrors = append(myErrors, errors.New("first name can't be empty"))
	}
	if len(p.LastName) == 0 {
		myErrors = append(myErrors, errors.New("last name can't be empty"))
	}
	if len(myErrors) > 0{
		// variadic parameters (...) is used because errors.Join() expects it
		return Person{},errors.Join(myErrors...)
	} 
	return p,nil
}

// Example for variadic parameters
func Sum(numbers ... int) int{
	mySum := 0 
	for _,num := range(numbers) {
		mySum += num
	}
	return mySum
}

func main(){
	p := Person{"",""}
	_,err := p.Validator()
	if err != nil{
		fmt.Println(err)
	}
	nums := make([]int,3)
	nums = append(nums, 1,4,10)
	// you can either pass a slice like this
	fmt.Println(Sum(nums...))
	// or you can pass multiple variables like this 
	// (variadic params will handle this and make it a slice)
	fmt.Println(Sum(1,4,10))
}