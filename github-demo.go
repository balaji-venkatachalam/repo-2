//Fibonacci is afunction which is a function; function that is being returned is a function that returns integer

package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		returnValue := a
		a, b = b, a+b
		return returnValue
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
