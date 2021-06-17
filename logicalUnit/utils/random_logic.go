package utils

func MakeArray(len int) []int {
	newArray := make([]int, len)
	for i := 0; i < len; i++ {
		newArray[i] = i
	}
	return newArray
}
