package main

import (
	"errors"
)

func Division(num, division_by int) (int, error) {
	if division_by == 0 {
		return 0, errors.New("Zero division error")
	}
	return num / division_by, nil
}
