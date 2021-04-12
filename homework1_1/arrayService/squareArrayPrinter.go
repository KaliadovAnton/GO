package arrayService

import (
	"fmt"
	"strconv"
)

func PrintArray(arraysToPrint [][]int) {
	for _, arrayToPrint := range arraysToPrint {
		s:= strconv.Itoa(arrayToPrint[0])
		for i:=1; i<len(arrayToPrint); i++ {
			s+=" "
			s += strconv.Itoa(arrayToPrint[i])
		}
		fmt.Println(s)
	}
}
