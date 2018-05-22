package main

import "fmt"

func main() {

	slc := []int{10, 7, 8, 9, 1, 5}

	fmt.Println("raw array")
	printArray(slc)

	n := len(slc)
	sort(slc, 0, n-1)

	fmt.Println("sorted array")
	printArray(slc);

}

/* The main function that implements QuickSort()
	 arr[] --> Array to be sorted,
	 low  --> Starting index,
	 high  --> Ending index */
func sort(slc []int, low int, high int) {

	var pi int

	if (low < high) {

		/* pi is partitioning index, arr[pi] is
  		now at right place */
		pi = partition(slc, low, high)

		// Recursively sort elements before
		// partition and after partition
		go sort(slc, low, pi-1)
		go sort(slc, pi+1, high)
	}
}

/* This function takes last element as pivot,
       places the pivot element at its correct
       position in sorted array, and places all
       smaller (smaller than pivot) to left of
       pivot and all greater elements to right
       of pivot */
func partition(slc []int, low int, high int) int {


	pivot := slc[high]
	i := low - 1 // index of smaller element

	for j := low; j < high; j++ {
		// If current element is smaller than or
		// equal to pivot
		if slc[j] <= pivot {
			i++

			// swap arr[i] and arr[j]
			temp := slc[i]
			slc[i] = slc[j]
			slc[j] = temp
		}
	}

	// swap arr[i+1] and arr[high] (or pivot)
	temp := slc[i + 1]
	slc[i + 1] = slc[high]
	slc[high] = temp

	return i + 1
}

func printArray(arr []int) {

	for key, val := range  arr {
		fmt.Println(key, " - ", val)
	}
	fmt.Println()
}