// all files in same directory must be under same package
// does this mean file? or actual directory?
package main

//most prewritten things in go you'll need a package that you have to import
import (
	"fmt"
)

func main() {
	// package name before the function call -- Println is function call
	// can only access Println cause it's a capital letter. lower case letters are not exported
	fmt.Println(("Hello BOOOOP"))

}
