package diceService

import "math/rand"

func ThrowADice() int {
	return 2 + rand.Intn(6) + rand.Intn(6)
}
