package empty_interface

import (
	"testing"
	"fmt"
)

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }
	// if i, ok := p.(string); ok {
	// 	fmt.Println("String", i)
	// 	return 
	// }
	// fmt.Println("Unknow Type")
	switch v:=p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("Hello World")
}

func returnMultiValues(n int)(int, int){
	return n+1, n+2
}

func TestReturnMultiValues(t *testing.T) {
	a, _ := returnMultiValues(5)
	t.Log(a)
}