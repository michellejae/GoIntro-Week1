package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func grep(term, inFile string) error {
	// return errors to try and mitigate and don't panic lol

	/*
		var file *os.File
		var err error
		file, err = os.Open(inFile)
	*/

	//os.Open returns the file object and a possible error value
	file, err := os.Open(inFile)
	// alwaqys have to manually check for error in GO
	if err != nil {
		return err
	}
	// if there is an error in the file, the file.Close would not happen
	// file is from os package and it has multiple different types, i guess it automaticaly knows it's a file?
	defer file.Close() // defer is executed when the function exits

	// line number
	lNum := 0
	s := bufio.NewScanner((file))
	// this for loop is like a while loop, it will loop through until there are no more lines
	for s.Scan() {
		lNum++
		// s.Text is the line, and term is what word we are looking for
		if strings.Contains(s.Text(), term) {
			// going to print file name, line number and the line
			fmt.Printf("%s:%d: %s\n", inFile, lNum, s.Text())
		}

	}

	// with a file it will likely not be an error, but if we pass in a socket it may have a network error and this will give us that
	return s.Err()
}

func main() {
	if err := grep("Print", "grep.go"); err != nil {
		// this prints to a file
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		// os.Exit means to end the program, which generally happens naturally if there are not errors but needs to be forced to happen
		// if there are erros. status 1 means general error
		os.Exit(1)
	}
}
