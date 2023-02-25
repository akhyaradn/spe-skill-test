package main

import (
	"fmt"
)

func main() {   
	fmt.Println(narcissisticNumber(153))
	fmt.Println(needleInHaystack([]string{"a","b", "c"}, "b"))
	fmt.Println(blueOceanReverse([]int{1,4,5,6,3}, []int{5}))
}

func blueOceanReverse(first []int, second []int) []int {
	result := []int{}
	for i := 0; i < len(first); i++ {
		exist := false
		for j := 0; j < len(second); j++ {
			if first[i] == second[j] {
				exist = true
			}
		}
		if !exist {
			result = append(result, first[i])
		}
	}

	return result
}

func needleInHaystack(arr []string, needle string) int {
	for i,v := range arr {
		if v == needle {
			return i
		}
	}
	return -1
}

func narcissisticNumber(input int) bool {
	cp := input
	if cp <= 0 {
		return false
	}

	result := make([]int, 0)

	for cp > 0 {
		result = append(result, cp%10)
		cp /= 10
	}

	digits := result
	sum := 0

	for i := 0; i < len(digits); i++ {
		sum += pow(digits[i], len(digits))
	}

	return sum == input
}

func pow(i, k int) int {
	result := 1

	for ; k > 0; k-- {
		result *= i
	}

	return result
}