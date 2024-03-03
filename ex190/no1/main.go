/*
1. Using the var keyword, declare a bidirectional unbuffered channel called c1 that works with values of type float64
2. Using the make() built-in function declare and initialize a receive-only channel called c2 and a send-only channel called c3.
Both work with data of type rune.
3. Declare a bidirectional buffered channel  called c4 with a capacity of 10 ints.
4. Print out the type of all the channels declared.
*/

package main

import "fmt"


func main (){

	var c1 chan float64 
	c1 = make(chan float64)

	c2 := make(<-chan rune) //recieving only
	c3 := make(chan<- rune) //sending only

	c4 := make(chan int, 10)

	fmt.Printf ("c1 is type: %T \n", c1)
	fmt.Printf ("c2 is type: %T \n", c2)
	fmt.Printf ("c3 is type: %T \n", c3)
	fmt.Printf ("c4 is type: %T \n", c4)

}

