package main

import (
	"errors"
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