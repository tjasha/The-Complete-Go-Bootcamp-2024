/*
Exercise 6
Create a Go Program that:
1. Using single function creates a file called info.txt in the current directory.
If the file already exists it will truncate it to zero size.
2. Write the string "The Go gopher is an iconic mascot!" to the file.
*/

package main

import (
	"io/ioutil"
	"log"
)

func main(){
	// file, err := os.OpenFile(
	// 	"info.txt",
	// 	os.O_CREATE|os.O_TRUNC,
	// 	0644,
	// )
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	byteSlice := []byte("The Go gopher is an iconic mascot!")
	// byteWritter, err := file.Write(byteSlice)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _ = byteWritter
	

	err := ioutil.WriteFile("info.txt", byteSlice, 0644)
		if err != nil {
		log.Fatal(err)
	}



}