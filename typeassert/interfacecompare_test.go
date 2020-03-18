package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTypes(t *testing.T) {
	types()
}

func TestCompare(t *testing.T) {
	var x interface{} = time.Now()
	x2 := x
	var y interface{} = time.Now()
	var z interface{} = []int{1, 2, 3}
	r := compare(x, y)
	s := compare(x, z)
	u := compare(x, x2)
	fmt.Printf("r %v\n", r)
	fmt.Printf("s %v\n", s)
	fmt.Printf("u %v\n", u)
}
