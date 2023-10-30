package utils

import "math/rand"

var (
	max int = 6
	min int = 1
)

func Roll() int {
	return rand.Intn(max) + min
}
