package main

import "fmt"

func main() {
	// Set the value you wish to search for in an array.
	var value_to_find int = 8888

	// Create an array to be searched.
	var custom_array = [12]int{1, 2, 3, 69, 420, 666, 6969, 8888, 10000, 15000, 25000, 50000}

	// Functions return true or false; print on those results.
	if binary_search(custom_array, value_to_find) {
		// Success.
		fmt.Print("True")
	} else {
		// Fail.
		fmt.Print("False")
	}
}

// A function that does a linear search on an sorted or unsorted array to find a specfic value.
// The goal is to iteratate over each value in the array and determine if the
// the value matches the needle provided.
//
// Big O Notation: O(N)
func linear_search(haystack [12]int, needle int) bool {
	// Create a loop. Use the length of the haystack array as the constraint for
	// the loop. Increment by 1.
	for i := 0; i < len(haystack); i++ {
		// Check the index for a match.
		if haystack[i] == needle {
			// Success.
			return true
		}
	}
	// Fail.
	return false
}

// A function that does a binary search on a sorted array to find a specfic value.
// This function can only work with a sorted array as the code is written for that
// assumption.
//
// Big O Notation: O(LOG)
func binary_search(haystack [12]int, needle int) bool {
	// Set the low variable.
	lo := 0
	// Set the upper(high) variable.
	hi := len(haystack)

	// Create a loop. Use the low and upper(high) variables for setting the
	// loops constraints. Increment by 1.
	for i := lo; i < hi; i++ {
		// Get the middle of the array by splitting it.
		var middle int = lo + (hi-lo)/2
		// First, check if middle index matches needle.
		if haystack[middle] == needle {
			// Success.
			return true
		}
		// Nothing found yet:
		// Second, check if the needle is greater than the middle index.
		if haystack[middle] > needle {
			// Success. Set the upper(high) variable to the middle index.
			hi = middle
		} else {
			// Otherwise, set the low variable to the midle index and increment by 1.
			lo = middle + 1
		}
	}
	// Fail.
	return false
}
