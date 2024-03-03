/*
1. Create a function called sum() that calculates and then prints out  the sum of 2 float numbers it receives as arguments.

Format the result with 2 decimal points.

2. From main launch 3 goroutines that execute the function you have just created (sum)

3. Synchronize the goroutines and the main function using WaitGroups
*/

package main

import (
	"fmt"
	"sync"
)


func sum(a, b float64, wg *sync.WaitGroup) {
	fmt.Printf("%.2f\n", (a+b))
	wg.Done()
}

func main (){

	var wg sync.WaitGroup
	wg.Add(3)
	for i:= 0; i<=2; i++{
		go sum(23.4345, 23.343433, &wg)
	} 
	wg.Wait()

}
