package main

import (
	"fmt"
	"runtime"
)


type empty interface{

}


func main(){

	//var a interface{} 2 options
	var a empty
	fmt.Printf("%T\n", a)

	a = 5
	fmt.Printf("%T\n", a)

	a = 45.46
	fmt.Printf("%T\n", a)

	a = []int{1, 2, 3, 4}
	fmt.Printf("%T\n", a)

	a = append(a.([]int), 5)
	fmt.Printf("%v\n", a)

	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())


}


/*
1. Declare a variable called empty of type empty interface. Print out its type.
2. Assign an int value to the variable called empty. Print out its type.
3. Assign a float64 value to empty. Print out its type.
4. Assign an int slice value to empty. Print out its type.
5. Add a new int value to the slice (empty variable).
6. Print out the slice (empty variable).
*/