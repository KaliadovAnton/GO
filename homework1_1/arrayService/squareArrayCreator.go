package arrayService

func CreateSquareArray(size int) [][]int {
	squareArray := make([][]int, size)
	for i:= 0; i<size;i++{
		squareArray[i] = make([]int,size)
	}
	return squareArray
}