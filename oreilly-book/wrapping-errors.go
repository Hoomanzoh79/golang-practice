package main

import (
	"errors"
	"fmt"
)

func doThing(val int) (int, error) {
	if val < 10 {
		return 0, errors.New("value should be bigger than 10")
	}
	val += 5
	return val,nil
}

// wrapping multiple errors with the same message 
// This can be handled with defer (beter way)
func DoSomeThings(val1, val2 int) (int, error) {
	val3,err := doThing(val1)
	if err != nil{
		return 0,fmt.Errorf("in do somethings: %w",err)
	}
	val4,err := doThing(val2)
	if err != nil{
		return 0,fmt.Errorf("in do somethings: %w",err)
	}
	result := val3 + val4
	return result,nil
}

// wrapping multiple errors with the same message using defer
func DoSomeThingsWithDefer(val1,val2 int)(_ int,err error) {
	defer func() {
		if err != nil{
			err = fmt.Errorf("in do somethings: %w",err)
		}
	}()
	val3,err := doThing(val1)
	if err != nil{
		return 0,err
	}
	val4,err := doThing(val2)
	if err != nil{
		return 0,err
	}
	result := val3 + val4
	return result,nil
}