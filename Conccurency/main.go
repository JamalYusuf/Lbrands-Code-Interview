/*

Write a buffered channel,
then write and receive to that channel

*/

package main

import (
	"fmt"
)

func main() {
	//Buffer length
	buffLen := 20

	//Buffered channel - bi directional
	bchan := make(chan int, buffLen)

	//Send to the channel a fixed amount of numbers up to buffLen
	//Will send a fixed amount of ints to chan
	go func() {
		defer close(bchan)
		//Iterate over loop N times
		for i := 0; i < buffLen; i++ {

			//send to the channel
			bchan <- i

		}

	}() //envoke

	//Receive data
	//Will print the ints received
	for v := range bchan {
		fmt.Println(v)
	}

}
