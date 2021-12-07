package functionDemo

import (
	"fmt"
	"testing"
)

var(
	Fun1 = func (n1 ,n2 int) int {
		return n1 + n2
	}
)

func TestMyFunctionSix(t *testing.T) {
	fmt.Println("观察执行顺序")
}

// 匿名函数使用方式1
func TestMyFunctionSeven(t *testing.T) {

	res1 := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println(res1)
}

// 匿名函数使用方式2
func TestMyFunctionEight(t *testing.T) {

	// 将匿名函数赋值给变量，通过变量去使用
	a := func(n1, n2 int) int {
		return n1 + n2
	}

	res1 := a(10,20)

	fmt.Println(res1)
}

// 全局匿名函数使用方式
func TestMyFunction9(t *testing.T) {

	res1 := Fun1(10,20)

	fmt.Println(res1)
}

