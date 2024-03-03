package main

import "fmt"

type book struct {
	price float64
	title  string
}

func (b book) vat() float64 {
	return b.price*0.09
}

func (b *book) discount() {
	b.price *=  0.90
}

func main(){


	myBook := book{price:40.5, title: "newTitle"}
	fmt.Printf("Vat: %v\n", myBook.vat())

	myBook.discount()
	fmt.Printf("New price: %v\n", myBook)


}





// 1. Create a new struct type called book with 2 fields: price and title of type float64 and string.
// 2. Create a method for the newly defined type called vat that returns the vat value if vat is 9%.
// 3. Create a method called discount that takes a book value as receiver and decreases the price of the book by 10%.