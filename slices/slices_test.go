package slices

import (
	"reflect"
	"testing"
)

func assertResult(want int, got int, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSum(t *testing.T) {

	t.Run("sum the integers passed in an array of ints", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		want := 15
		got := Sum(in)
		assertResult(want, got, t)
	})
}

func assertResultSlice(want, got []int, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("sum the 2 slices and return a new slice of the sums", func(t *testing.T) {
		in1 := []int{1, 2}
		in2 := []int{0, 9}
		want := []int{3, 9}
		got := SumAll(in1, in2)
		assertResultSlice(want, got, t)
	})
	t.Run("sum the 3 slices of differing lengths and return a new slice of the sums", func(t *testing.T) {
		in1 := []int{1, 2, 4, 5}
		in2 := []int{0, 9}
		in3 := []int{1}
		want := []int{12, 9, 1}
		got := SumAll(in1, in2, in3)
		assertResultSlice(want, got, t)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum tails (trailing items not heads) of 2 slices of differing lengths and return a new slice of the sums", func(t *testing.T) {
		in1 := []int{1, 2, 4, 5}
		in2 := []int{2, 9}
		in3 := []int{100}
		want := []int{11, 9, 0}
		got := SumAllTails(in1, in2, in3)
		assertResultSlice(want, got, t)
	})
	t.Run("sum tails (trailing items not heads) of 2 slices and one is empty return a new slice where empty one is 0", func(t *testing.T) {
		in1 := []int{2, 9}
		in2 := []int{}
		want := []int{9, 0}
		got := SumAllTails(in1, in2)
		assertResultSlice(want, got, t)
	})
}
