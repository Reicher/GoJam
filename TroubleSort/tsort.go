package main

import (
	//"log"
	"fmt"
	"strconv"
)

func tsort(d int, v []int) []int {
	done := false
	for !done {
		done = true
		for i := 0; i < d-2; i++ {
			if v[i] > v[i+2]{
				done = false
				tmp := v[i]
				v[i] = v[i+2]
				v[i+2] = tmp
			}
		}
	}
	return v
}

func isSorted(d int, v []int) bool {
	return false
}

func solver(d int, v []int ) string{
	sorted := tsort(d, v)
	//fmt.Println(sorted)
	if isSorted(d, sorted){
		return "OK"
	} else {
		return "NOT OK"
	}
}

func main() {
	T := 0
	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		n := 0
		fmt.Scanf("%d", &n)

		var v []int
		for i := 0; i < n; i++ {
			a := 0
			fmt.Scanf("%d", &a)
			v = append(v, a)
		}

		solution := solver(n, v)

		line := "Case #" + strconv.Itoa(i+1) + ": " + solution
		fmt.Println(line)
	}
}
