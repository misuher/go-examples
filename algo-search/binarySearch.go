package binarySearch

func Bsearch(arr []int, elemToFind int) int {
	//var low, high int
	i := 0
	j := len(arr)
	for i < j {
		middle := i + (j-i)/2 //avoid overflow
		if arr[middle] < elemToFind {
			i = middle + 1
		} else if arr[middle] > elemToFind {
			j = middle
		} else {
			return middle
		}
	}
	return -1
}
