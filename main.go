package main

import "fmt"

func main() {

	fmt.Println(FindMultiples(3, 20))

}

//function takes two values which is the integer, you find multiples of integer upto the limit value
func FindMultiples(integer, limit int) []int {
	var sliceint []int
	n := limit
	for i := integer; i <= n; i++ {
		if i%integer == 0 {
			sliceint = append(sliceint, i)

		}

	}
	return sliceint
}
