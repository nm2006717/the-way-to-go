package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Print("First example with -1: ")
	ret1, err1 := mySqrt(-1)
	if err1 != nil {
		fmt.Println("Error! Return values are: ", ret1, err1)
	} else {
		fmt.Println("It's ok! Return values are: ", ret1, err1)
	}

	fmt.Print("Second example with 5: ")
	//you could also write it like this
	if ret2, err2 := mySqrt2(5); err2 != nil {
		fmt.Println("Error! Return values are: ", ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are: ", ret2, err2)
	}
	// named return variables:
	fmt.Println(mySqrt2(5))

}

func mySqrt(value float64) (float64, error) {
	if value < 0 {
		return math.NaN(), errors.New("参数不能为负值")
	}
	return math.Sqrt(value), nil
}

func mySqrt2(value float64) (ret float64, err error) {
	if value < 0 {
		ret = math.NaN()
		err = errors.New("参数不能为负值")
	}
	ret = math.Sqrt(value)
	err = nil
	return
}
