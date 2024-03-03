package main

import (
	"fmt"
	"os"
	"strconv"
)

func main (){
	sum, product := 0, 0

	for i :=1; i<len(os.Args); i++ {

		value, err:=strconv.Atoi(os.Args[i])

		sum = sum + value
		product *= value
		_=err

	} 
	fmt.Println("Sum: ", sum, "product: ", product)
}

/*
Create a Go program that reads some numbers from the command line and 
then calculates the sum and the product of all the numbers given at the command line.

The user should give between 2 and 10 numbers.

Notes:

- the program should be run in a terminal (go run main.go) not in Go Playground

- example:

go run main.go 3 2 5

Expected output: Sum: 10, Product: 30 */