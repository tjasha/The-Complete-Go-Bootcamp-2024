/*
Change the program from Exercise #4 and calculate the square of all values between 1 and 50 included using an anonymous function.
*/

package main

import "fmt"

func main(){

	ch := make(chan int)

	for i:=1; i<=50; i++{
		go func(num int, c chan int) {
			square := num * num 
			c <- square
		}(i, ch)
	}

	for i:=1; i<=50; i++{
		fmt.Println(<-ch )
	}


}
