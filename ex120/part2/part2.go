package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	/*
	   Exercise 4
	   Create a Go Program that reads the entire contents of a file called info.txt into a string.
	   You can use ioutil.ReadAll() or any other function you want.

	   The file exists in the current working directory.
	*/

	var fileName = "info.txt"
	file, err := os.Open(fileName)
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	fmt.Println("####")

	// opton 2
	data2, err := os.ReadFile(fileName)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(string(data2))
	fmt.Println("####")


	/*
	Exercise 5 
	Create a Go Program that reads the entire contents of a file called info.txt  
	using a scanner (bufio package) line by line.

	The file exists in the current working directory.
	*/

	file2, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	scanner:= bufio.NewScanner(file2)
	for scanner.Scan(){
		println(scanner.Text())
	}
	fmt.Println("####")
	if err := scanner.Err(); err !=nil{
		log.Fatal()
	}

}
