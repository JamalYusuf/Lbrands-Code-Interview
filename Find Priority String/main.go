/*

given a input, find the priority string in the source string

example:

given the priority string "hello"

find all other strings that are similar,

you can use standard library sort interface


*/

package main

import (
	"fmt"
	"sort"
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

	var Cp ByCodePoints = ConvertStringToRuneSlice(s) //get our rune slice, store in alias ByCodePoints

	//We can use the amazing sort package

	//sort the string by rune value
	sort.Sort(Cp)

	//return it back as a string

	return string(Cp)
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

//Implement the sort interface

//Len
func (r ByCodePoints) Len() int { return len(r) }

//Swap
func (r ByCodePoints) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

//Less
func (r ByCodePoints) Less(i, j int) bool { return r[i] < r[j] }
