package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func types() {
	var w io.Writer
	//w.Write([]byte("hello")) // panic

	w = os.Stdout
	w.Write([]byte("hello\n")) // will print "hello"
	// same as os.Stdout.Write()
	os.Stdout.Write([]byte("hello again\n"))

	w = new(bytes.Buffer)
	w.Write([]byte("hello - new bytes buffer\n")) // will print "hello"
	fmt.Printf("buffer = %v\n", w)

	w = nil
	//w.Write([]byte("hello\n")) // panic
}

func compare(x, y interface{}) bool {
	if x == y {
		return true
	}
	// will cause panic
	return false
}

const debug = false

func main() {
	//var buf *bytes.Buffer // panic debug false
	var buf io.Writer // solution this is valid for receiver
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {
		fmt.Printf("type %T\n", buf)
	}
	fmt.Println("ending")
}

func f(out io.Writer) {
	fmt.Println("do something")
	if out != nil {
		out.Write([]byte("done!\n"))
		fmt.Printf("write to output\n")
	}
}
