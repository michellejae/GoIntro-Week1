package main

import (
	"compress/bzip2"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func revenue(csvFile string) (float64, error) {

	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	defer file.Close()

	r := csv.NewReader(bzip2.NewReader(file))
	// which reader gets called first?

	lNum, total := 0, 0.0

	//INFINITE WHILE LOOP USED TO JUST BUST OUT WHEN IT HITS A RETURN lol
	for {
		record, err := r.Read()
		if err == io.EOF {
			// stream is completed
			break
		}
		if err != nil {
			return 0, err
		}
		// each line number is a stream or record
		lNum++
		if lNum == 1 {
			continue // skip header
		}
		amount, err := strconv.ParseFloat(record[len(record)-2], 64)
		if err != nil {
			return 0, err
		}
		total += amount

	}

	return total, nil

}

func main() {

	rev, err := revenue("./taxi-1k.csv.bz2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rev)

}
