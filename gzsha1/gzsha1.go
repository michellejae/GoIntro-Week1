// want to find digital signature of uncompressed file

package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("httpd.log.gz")
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	defer file.Close()

	// need to DECOMPRESS FILE
	// newReader again will use ioReader and will return another reader and when you read from this new reader it will decompress the file
	r, err := gzip.NewReader(file)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	w := sha1.New()
	// io.Copy very versatile, copies from a writer to a reader here our reader (r)
	//is the one returned from newReader and writer (w) is from sha1
	// io copy checks for number of bytes written and possible error
	if _, err := io.Copy(w, r); err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	// print digital signature of file
	fmt.Printf("%x\n", w.Sum(nil))

}
