package main

import "fmt"

type Car struct {
	name     string
	model_no int
}

// fmt package and other packages in go look for Stringer which is an interface that declares itself for various operations
// here there is a stringer that accepts Car and print and returns some value based on the car
// println in the main would look if there is any stringer for that specific value if so use that.
func (c Car) String() string {
	return fmt.Sprintf("Dang bro you got the \"%v\" model %v\n", c.name, c.model_no)
}

func main() {
	a := Car{"Bugatti", 5}
	b := Car{"Ferrari", 1}
	fmt.Println(a,b)
}
