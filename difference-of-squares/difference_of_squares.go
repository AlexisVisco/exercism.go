package diffsquares

import "math"

type apply func(int, int) int

func Applier(number int, fn apply) int {
	result := 0
	for i := 1; i <= number; i++ {
		result = fn(result, i)
	}
	return result
}

func SumOfSquares(number int) int {
	return Applier(number, func(last int, next int) int {
		return (last) + (next * next)
	})
}

func SquareOfSums(number int) int {
	return int(math.Pow(float64(Applier(number, func(last int, next int) int {
		return last + next
	})), 2))
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
