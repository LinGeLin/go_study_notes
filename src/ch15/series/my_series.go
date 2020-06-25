package series

import "fmt"

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

// 开头是小写，则包外无法访问
func square(n int) int {
	return n*n
}

func GetFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-1] + ret[i-2])
	}
	return ret
}
