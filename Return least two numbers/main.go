package main

import "fmt"

/*
given a slice of ints, return the two smallest numbers

*/

func main() {

	input := []int{10, 23, 4008, 84832, 333, 1}

	var (
		smallest, smallestSecond int
	)
	//iterate through the numbers
	for i, v := range input {
		//First run
		if i == 0 {
			smallest = v
			smallestSecond = v
		}

		//Find Smallest
		if smallest > v {
			//Assign second smallest
			smallestSecond = smallest

			//Assign smallest
			smallest = v

		}
	}

	//Print the output

	fmt.Printf("\n\nSmallest Number is:%d\nSecond Smallest is:%d\n\n", smallest, smallestSecond)

}
