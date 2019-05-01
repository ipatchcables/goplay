/*
Write a program which prompts the user to enter integers and stores the
integers in a sorted slice. The program should be written as a loop. Before
entering the loop, the program should create an empty integer slice of size
(length) 3. During each pass through the loop, the program prompts the user
to enter an integer to be added to the slice. The program adds the integer
to the slice, sorts the slice, and prints the contents of the slice in
sorted order. The slice must grow in size to accommodate any number of
integers which the user decides to enter. The program should only quit
(exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	myArr := [3]int{} // Per requirement, create array slice of len 3
	mySlice := myArr[0:3]
	var myInt string    // Declared as string to allow "X" entry for quitting
	for i := 0; ; i++ { // Loop forever, incrementing i
		fmt.Printf("Enter an integer (or \"X\" to quit): ") // Used printf to avoid trailing \n
		fmt.Scan(&myInt)                                    // Read user input into a temporary variable
		if myInt != "X" {                                   // Check if entered value is X, if not, run loop
			userInt, _ := strconv.Atoi(myInt) // Convert value entered to int
			if i > len(mySlice)-1 {           // If the current instance of the loop goes beyond the size of the slice
				mySlice = append(mySlice, 1) // Increase the length of the slice by 1
			}
			mySlice[i] = userInt // Store the entered value into the current slice position
			sort.Ints(mySlice)   // Sort the slice
			fmt.Println(mySlice) // Print the sorted slice
		} else {
			break // Exits the loop (only if entered value is X)
		}
	}
}
