package error_test

import (
	"testing"
	"errors"
	"strconv"
	"fmt"
)

var LessThanTwoError = errors.New("n should not be less than 2")
var LargerThanHunderedError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHunderedError
	}
	fibList := []int{1, 1}
	for i:=2; i<n; i++ {
		fibList = append(fibList, fibList[i-2] + fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(1001); err != nil {
		if err == LessThanTwoError {
			t.Error(err)
		} else if (err == LargerThanHunderedError) {
			t.Error(err)
		}
	} else {
		t.Log(v)
	}
}

func GetFibonacci1(str string) {
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Print(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

func GetFibonacci2(str string) {
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}

	if list,err = GetFibonacci(i); err != nil {
		fmt.Println("Error, err")
		return
	}
	fmt.Println(list)
}