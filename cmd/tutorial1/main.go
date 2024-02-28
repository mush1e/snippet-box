package main

import (
	"errors"
	"fmt"
)

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by 0")
		return 0, 0, err
	}
	return numerator / denominator, numerator % denominator, err
}

func main() {
	fmt.Println("Hello world!!")
	var result, remainder, err = intDivision(4, 0)
	arr := []int32{1, 2, 3}
	switch {
	case err != nil:
		fmt.Println(err.Error())

	case result == 0:
		fmt.Println("The result is 0")

	default:
		fmt.Printf("The result is %v and the remainder is %v", result, remainder)
	}
	fmt.Println(cap(arr))
	arr = append(arr, 7)
	fmt.Println(len(arr), cap(arr))

	for i, v := range arr {
		fmt.Println(i, v)
	}

}
