package main

import "fmt"

func main() {

	var arr = []int{10, 7, 8, 9, 1, 5}

	fmt.Println("raw array")
	printArray(arr)

	n := len(arr)
	sort(arr, 0, n-1)

	fmt.Println("sorted array")
	printArray(arr);

}

/* The main function that implements QuickSort()
	 arr[] --> Array to be sorted,
	 low  --> Starting index,
	 high  --> Ending index */
func sort(arr []int, low int, high int) {

	if (low < high) {
		var pi int
		/* pi is partitioning index, arr[pi] is
  		now at right place */
		pi = partition(arr, low, high)

		// Recursively sort elements before
		// partition and after partition
		sort(arr, low, pi-1)
		sort(arr, pi+1, high)
	}
}

/* This function takes last element as pivot,
       places the pivot element at its correct
       position in sorted array, and places all
       smaller (smaller than pivot) to left of
       pivot and all greater elements to right
       of pivot */
func partition(arr []int, low int, high int) int {


	pivot := arr[high]
	i := low - 1 // index of smaller element

	for j := low; j < high; j++ {
		// If current element is smaller than or
		// equal to pivot
		if arr[j] <= pivot {
			i++

			// swap arr[i] and arr[j]
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	// swap arr[i+1] and arr[high] (or pivot)
	temp := arr[i+1]
	arr[i + 1] = arr[high]
	arr[high] = temp

	return i + 1
}

func printArray(arr []int) {

	n := len(arr)

	for i := 0; i < n; i++ {
		fmt.Println(i, " - ", arr[i])
	}
	fmt.Println()
}