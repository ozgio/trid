package trid

import (
	"fmt"
	"strconv"
	"strings"
)

// Valid returns if the id is valid. It is a convenience function to easily get the bool result of
// the validation.
func Valid(id string) bool {
	return Validate(id) != nil
}

// Validate validates the id and returns an error if it is not valid.
func Validate(id string) error {
	if len(id) != 11 {
		return fmt.Errorf("invalid length")
	} else if id[0] == '0' {
		return fmt.Errorf("id cannot start with 0")
	}

	var (
		sumEven int
		sumOdd  int
		sumAll  int
		d       int
		err     error
	)

	for i := 0; i < 9; i++ {
		d, err = strconv.Atoi(string(id[i]))
		if err != nil {
			return fmt.Errorf("'%c' at digit %d is not a number", id[i], i+1)
		}
		sumAll += d
		if (i+1)%2 == 0 {
			sumEven += d
		} else {
			sumOdd += d
		}
	}

	cs1 := (sumOdd*7 - sumEven) % 10
	cs2 := (sumAll + cs1) % 10

	checksum := fmt.Sprintf("%d%d", cs1, cs2)
	if !strings.HasSuffix(id, checksum) {
		return fmt.Errorf("invalid checksum, id should end with %s", checksum)
	}
	return nil
}
