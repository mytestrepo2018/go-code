package slices

//Sum takes an array of ints and returns the sum of integers in the array
func Sum(in []int) (total int) {
	total = 0
	for _, v := range in {
		total += v
	}
	return
}

//SumAll takes n slices and sums them and returns them as a new slice
func SumAll(in ...[]int) (sums []int) {
	for _, v := range in {
		sums = append(sums, Sum(v))
	}
	return
}

//SumAllTails takes n slices and sums them and returns them as a new slice
func SumAllTails(in ...[]int) (sums []int) {
	for _, v := range in {
		//		sums = append(sums, sumTail(v))
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(v[1:]))
		}
	}
	return
}

/* func sumTail(in []int) (total int) {
	total = 0
	for k, v := range in {
		if k != 0 {
			total += v
		}
	}
	return
} */
