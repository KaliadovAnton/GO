package main

import (
	"homework1_1/arrayService"
)

func main() {
	squareArray := arrayService.CreateSquareArray(5)
	arrayWithDiagonals := arrayService.FillDiagonals(1, squareArray)
	arrayService.PrintArray(arrayWithDiagonals)
}
