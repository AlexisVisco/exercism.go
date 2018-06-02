package grains

import (
	"errors"
	"math"
)

// Sum of a geometric sequence of U(n+1)=2*U(n) == 2^(64+1) - 1.
// 2^(64+1) - 1 == 3.68934881e19 == 36899348810000000000
// 36899348810000000000 > math.MaxUint64
// value of this exercise cannot be performed in despite of
// the architecture of our computer and how go has implemented
// integers.
// Since we must return something, we must return what we are able to return.
func Total() uint64 {
	return math.MaxUint64
}

func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("negative greather than 64")
	}
	return 1 << uint64(n-1), nil
}
