package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	a := 1
	{
		// a := 2	// shadows outer a
		a = 2                   // change outer a
		fmt.Println("inner", a) // affects only inner a
	}
	fmt.Println("outer")
	sig, err := sha1Sum("http.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)

	sig, err = sha1Sum("sha1.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
}

// if fale name ends with .gz
// $ cat http.log.gz | gunzip | sha1sum
// else $ cat filename | sha1sum
func sha1Sum(fileName string) (string, error) {
	// $ ulimit -a to check how many file descriptors available per process
	// important to close these calls, hard to debug
	// idiom: open a resource, check for error, defer release
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	// happens when the surrounding function for defer statement returns
	// multiple defer statements get pushed onto a stack and run in reverse order
	defer file.Close()
	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer gz.Close()
		r = gz
	}

	// io.CopyN(os.Stdout, r, 100)
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", nil
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil

}
