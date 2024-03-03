package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}
 
type car struct {
	licenceNo string
	brand     string
}

func (c car) Licence() string{
 return c.licenceNo
}

func (c car) Name() string{
	return c.brand
}


func main(){

	myCar := car{licenceNo: "uh2379", brand: "Audi"} 
	fmt.Println(myCar.Licence())
	fmt.Println(myCar.Name())

}




/* 1. Create a Go program where car type implements the vehicle interface.

2. Create a variable of type vehicle that holds a car struct value.

3. Call the methods (Licence and Name) on the interface value declared at step 2

4. Run the program without errors. */