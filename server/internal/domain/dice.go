package domain

import "math/rand"

func D10() int {
	return rand.Intn(10)
}

func TwoD10() int {
	return D10() + D10()
}

func ThreeD10() int {
	return D10() + D10() + D10()
}

func D100() int {
	return rand.Intn(100)
}
