package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func tryTypeAssert() {
	var w io.Writer
	w = os.Stdout
	f, ok := w.(*os.File)
	if ok {
		fmt.Println("ok File")
	}
	c, ok := w.(*bytes.Buffer)
	if ok {
		fmt.Println("ok bytesBuffer")
	}
	d, ok := w.(io.ReadWriter)
	if ok {
		fmt.Println("ok ReadWriter")
	}
	fmt.Printf("%T and %T and %T\n", f, c, d)
	fmt.Printf("the end\n")
}
