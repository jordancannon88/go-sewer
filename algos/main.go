package main

import (
	"fmt"
)

func main() {
	// Set the value you wish to search for in an array.
	var valueToFind int = 2

	// Create an array.
	// var customArray = [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	// Create a slice.
	var customSlice = []int{2, 4, 6, 8, 10}

	// Functions return true or false; print on those results.
	if binarySearch(customSlice, valueToFind) {
		// Success.
		fmt.Print("True")
	} else {
		// Fail.
		fmt.Print("False")
	}

	// Call single functions.
	// fizzBuzz(1, 100)
	// fizzBuzzRedux(1, 100)
}

// A function that does a linear search on an ordered or unordered array to find a specfic value.
// The goal is to iteratate over each value in the array and determine if the
// the value matches the needle provided.
//
// Big O Notation: O(n)
func linearSearch(haystack []int, needle int) bool {
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

// A function that does a binary search on a ordered array to find a specfic value.
// This function can only work with a ordered array as the code is written for that
// assumption.
//
// Big O Notation: O(log n)
func binarySearch(haystack []int, needle int) bool {
	// Set the low variable.
	var lo = 0
	// Set the upper(high) variable.
	var hi = len(haystack)

	// Create a loop.
	for i := 0; i < len(haystack); i++ {
		// Get the middle index of the array by splitting it.
		var middleIndex int = lo + (hi-lo)/2
		// Initiliaze var.
		var middleValue int
		// If the middleIndex is greater than the length of the array then the needle doesn't exist.
		if len(haystack)-1 < middleIndex {
			return false
		}
		// Set the value of the middle index.
		middleValue = haystack[middleIndex]
		// First, check if middle index matches needle.
		if middleValue == needle {
			// Success.
			return true
		}
		// Nothing found yet:
		// Second, check if the middle index of the array is greater than the needle.
		if middleValue > needle {
			// Success. Set the upper(high) variable to the middle index.
			hi = middleIndex
		} else {
			// Otherwise, it is lower. Set the low variable to the midle index and increment by 1.
			lo = middleIndex + 1
		}
	}
	// Fail.
	return false
}

// Return a new array after doubling every item in the original array.
//
// Big O Notation: O(n)
func double(orginal_array []int) []int {
	for i := 0; i < len(orginal_array); i++ {
		orginal_array[i] = orginal_array[i] * 2
	}
	return orginal_array
}

// Add every item in the original array and return the result.
//
// Big O Notation: O(n)
func add(original_array []int) int {
	result := 0
	for i := 0; i < len(original_array); i++ {
		result += original_array[i]
	}
	return result
}

func binarySearchRedux(haystack []int, needle int) bool {
	// Set the low variable.
	var lo = 0
	// Set the upper(high) variable.
	var hi = len(haystack)

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
		// Second, check if the middle index of the array is greater than the needle.
		if haystack[middle] > needle {
			// Success. Set the upper(high) variable to the middle index.
			hi = middle
		} else {
			// Otherwise, it is lower. Set the low variable to the midle index and increment by 1.
			lo = middle + 1
		}
	}
	// Fail.
	return false
}

// For multiples of 3 say Fizz. For multiples of 5 say Buzz. And for multiples of both 3 and 5 say FizzBuzz.
// Print everything else.
func fizzBuzz(start int, end int) {
	for i := start; i < (end + 1); i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// For multiples of 3 say Fizz. For multiples of 5 say Buzz. And for multiples of both 3 and 5 say FizzBuzz.
// Print everything else.
func fizzBuzzRedux(start int, end int) {
	for i := start; i < (end + 1); i++ {
		mixedOutput := []interface{}{}

		if i%3 == 0 {
			mixedOutput = append(mixedOutput, "Fizz")
		}
		if i%5 == 0 {
			mixedOutput = append(mixedOutput, "Buzz")
		}

		if mixedOutput != nil {
			mixedOutput = append(mixedOutput, i)
		}

		fmt.Println(mixedOutput[0])
	}
}
