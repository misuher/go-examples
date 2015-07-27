package binarySearch

func Bsearch(arr []int, elemToFind int) int {
	low := 0
	high := len(arr)
	for low < high {
		middle := low + (high-low)/2 //avoid overflow
		if arr[middle] < elemToFind {
			low = middle + 1
		} else if arr[middle] > elemToFind {
			high = middle
		} else {
			return middle
		}
	}
	return -1
}
