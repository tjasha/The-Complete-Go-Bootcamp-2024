/*
Create a function literal (a.k.a. anonymous function) that sends the string value if receives as argument to main func using a channel.
*/

package main

import "fmt"

func main (){

	c1 := make(chan string)
	a := "lalala"

	go func(str string){
		c1 <- str
	}(a)

	geslo:=<-c1
	fmt.Println(geslo)

}