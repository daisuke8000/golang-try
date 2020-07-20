/*
mylib is myspecial lib.
*/
package mylib

import "fmt"

// Person2
type Person2 struct{
	//Name
	Name string
	//Age
	Age int
}

func (p *Person2) Say(){
	fmt.Println("Person2")
}

//Average return the average of a series of number
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
