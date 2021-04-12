package main

import (
	"fmt"
	"homework1_2/diceService"
)

func main() {
	var sumOfDices int
	var everyThrow []int
	for i := 0; i<1000; i++{
		sumOfDices = diceService.ThrowADice()
		everyThrow = append(everyThrow, sumOfDices)
	}
	result := make(map[int]int)
	for _, throw := range everyThrow{
		result[throw]++
	}
	var sortedResults = make([]float64, 11)
	for numberOnDices, amount := range result {
		percentage := float64(amount)
		percentage = percentage/10
		sortedResults[numberOnDices-2] = percentage
	}
	for i:=0;i<len(sortedResults);i++ {
		fmt.Printf("%d - %.1f%%\n", i+2, sortedResults[i])
	}
}
