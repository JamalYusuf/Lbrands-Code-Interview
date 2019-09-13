/*

given a input, find the priority string in the source string

example:

given the priority string "hello"

find all other strings that are similar,

You cannot use standard library sort


*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	//The source stream
	input := "Hello lleho HllEo Cat cat tac acT"

	//Priority Key
	priKey := "Hello"

	//stores the relations between our source and string
	relationMap := make(map[string][]string)

	//Split string into []string
	//split by space
	inputSpl := strings.Split(input, " ")

	//iterate through slice of string
	for _, v := range inputSpl {

		//Create Key
		key := SortStringByRune(v)

		//Add value to map to build relationship
		relationMap[key] = append(relationMap[key], v)
	}

	//Print Relation of Priority Key to string

	key := SortStringByRune(priKey) //Get the Key

	fmt.Printf("\n\nInputString: %q\nPriority Key: [ %s ]\nRelationMap:%+v\n\n", input, priKey, relationMap[key])

}

//Sort the string by code point
func SortStringByRune(s string) string {

	//normalize data
	s = strings.ToLower(s) //To prevent mismatching of cases

	var input ByCodePoints = ConvertStringToRuneSlice(s) //get our rune slice, store in alias ByCodePoints

	//We can use the amazing sort package

	//sort the string by rune value
	////

	// n is the number of items in our list
	n := len(input)

	// set swapped to true
	swapped := true

	// loop
	for swapped {

		// set swapped to false
		swapped = false

		// iterate through all of the elements in our list
		for i := 1; i < n; i++ {
			// if the current element is greater than the next
			// element, swap them
			if input[i-1] > input[i] {

				// swap values using Go's tuple assignment
				input[i], input[i-1] = input[i-1], input[i]

				// set swapped to true - this is important
				// if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is
				// fully sorted.
				swapped = true
			}
		}
	}
	// finally, print out the sorted list
	fmt.Println(input)

	//return it back as a string

	return string(input)
}

//Func ConvertStringToRuneSlice, will given a string, extract
//individual UTF-8 code points and store them in a slice of rune
func ConvertStringToRuneSlice(s string) []rune {
	var xiRune []rune

	for _, codePoint := range s {
		xiRune = append(xiRune, codePoint) //Add our codePoint to slice of rune
	}

	//we have finished converting it, return it

	return xiRune
}

/*

Important notes, in Golang source files are encoded in
utf-8 format, meaning that each codePoint(a single symbol) can
be larger than a single byte which is int8, in unicode a single
codePoint can be stored in a int32, which in golang can
be thought of as a rune.

Because your strings are already in UTF-8, you can easily pipe
them into a rune.
*/
type ByCodePoints []rune //This is a string by UTF-8, codepoints
