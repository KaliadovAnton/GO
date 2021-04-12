package arrayService

func FillDiagonals(symbol int, arrays [][]int) [][]int {
	for i:=0;i<len(arrays);i++ {
		arrays[i][i] = symbol
		arrays[i][len(arrays)-i-1] = symbol
	}
	return arrays
}