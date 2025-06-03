package otp

import (
	"math"
	"math/rand"
)

func GetOtpWithNumbers(digitCount int) int {
    return generateRandomNumber(digitCount)
}
func generateRandomNumber(digits int) int {
	min := int(math.Pow(10, float64(digits-1)))
	max := int(math.Pow(10, float64(digits))) - 1
	return rand.Intn(max-min+1) + min
}