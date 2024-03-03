package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("what? %.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("yey %.2f", m)
}

func main() {

	myMoney := money(32.33245544)

	myMoney.print()

	var string = myMoney.printStr()
	fmt.Println(string)

}

//1. Create a new type called money. Its underlying type is float64.
// . Create a method called print that prints out the money value with exact 2 decimal points.
//Create a new method for the money type called printStr that returns the money value as a string with 2 decimal points.
