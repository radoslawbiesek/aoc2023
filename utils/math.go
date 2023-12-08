package utils

import "math"

func GCD(a, b int) int {
	for a%b > 0 {
		r := a % b
		a = b
		b = r
	}

	return b
}

func LCM(a, b int) int {
	return int(math.Abs(float64(a)*float64(b))) / GCD(a, b)
}

func LCMSlice(s []int) int {
	result := 1
	for _, num := range s {
		result = LCM(result, num)
	}

	return result
}
