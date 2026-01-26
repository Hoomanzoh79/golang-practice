package main

import (
	"errors"
	"fmt"
)

type Employee struct{
	ID int
	FirstName string
	LastName string
	Address string
}

const InvalidID string = "Invalid ID"
const EmptyField string = "Empty Field"

// sentinel error
var ErrInvalidID error = errors.New(InvalidID)

// custom error 
type EmptyFieldError struct{
	Field string
	Message string
}
func (err EmptyFieldError) Error() string {
	return err.Message
}

func validateEmployee(emp Employee) error {
	var errorsSlice[] error
	if emp.ID < 0 || emp.ID > 1000 {
		errorsSlice = append(errorsSlice, ErrInvalidID)
	} 
	if emp.FirstName == "" {
		errorsSlice = append(errorsSlice, EmptyFieldError{Field: "ّFirstName",Message: EmptyField})
	}
	if emp.LastName == "" {
		errorsSlice = append(errorsSlice, EmptyFieldError{Field: "LastName",Message: EmptyField})
	}
	if emp.Address == "" {
		errorsSlice = append(errorsSlice, EmptyFieldError{Field: "Address",Message: EmptyField})
	}
	return errors.Join(errorsSlice...)
}

func main() {
	testCases := []Employee{
		{ID: 0, FirstName: "John", LastName: "Doe", Address: "123 Main St"},
		{ID: 1, FirstName: "", LastName: "Smith", Address: "456 Oak Ave"},
		{ID: 2, FirstName: "Jane", LastName: "", Address: ""},
		{ID: -1, FirstName: "", LastName: "", Address: ""},
		{ID: 3, FirstName: "Bob", LastName: "Johnson", Address: "789 Pine Rd"},
	}
	for _,emp := range testCases{
		err := validateEmployee(emp)
		if err != nil {
			if errors.Is(err,ErrInvalidID){
				fmt.Printf("Invalid ID for %+v \n",emp)
			}
			var emptyFieldErr EmptyFieldError
			if errors.As(err,&emptyFieldErr) {
				fmt.Printf("Found empty field error: %s\n", emptyFieldErr.Field)
			}
			fmt.Printf("Validation errors: %s\n", err.Error())
		} else {
			fmt.Println("Validation passed")
		}
	}
}