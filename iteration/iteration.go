package iteration

import "strings"

//IterateChar takes a character and repeats it the specified times and returns it
func IterateChar(char string, times int) string {
	//r := ""
	//for i := 0; i < times; i++ {
	//	r += char
	//}
	//return r
	return strings.Repeat(char, times)
}
