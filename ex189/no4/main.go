/*
Change the code from Exercise #3 and launch 50 goroutines that calculate concurrently
the square root of all the numbers between 100 and 149 (both included).
*/

package main

import (
	"fmt"
	"math"
	"sync"
)

func main (){

var wg sync.WaitGroup
wg.Add(50)

for i := 100; i<=149; i++{
	go func (a float64, wg *sync.WaitGroup){
		x := math.Sqrt(a)
		fmt.Println("Square root from " , a ," is " , x)
		wg.Done()
	}(float64(i), &wg)
}
wg.Wait()
}