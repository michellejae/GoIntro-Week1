package main

import "fmt"

func main() {
	a, b := 1, "1"
	fmt.Println(a, b)
	fmt.Printf("a=%d, b=%q\n", a, b)
	fmt.Printf("a=%#v, b=%#v\n", a, b)

	// %v mean print value with type
	// %#v value with type
	// %d for numbers
	// %q for quoted strings

	n := 21
	fmt.Printf("dec: %d, hexa: %x, octa: %o, bin:%b\n", n, n, n, n)
	fmt.Printf("%04d\n", n) // want it with four characters, pair it with 0 if not there
	// 0021

	name := "Joe"
	fmt.Printf("%20s\n", name) // want it with at least 20 characters long, so indent to the right .. can do -20 to go indenting to the left

	item, price := "carrot", 10
	fmt.Printf("%[1]s cost $%[2]d, $%[2]d\n", item, price) //can refer to each item but position as well

	type Player struct {
		Name  string
		Level int
	}

	p1 := Player{"Boop", 100}
	fmt.Println(p1)         // prints out with curly braces {Boop 100}
	fmt.Printf("%+v\n", p1) // prints out with field next to value {Name:Boop Level:100}
	fmt.Printf("%#v\n", p1) // with quotes main.Player{Name:"Boop", Level:100}

}
