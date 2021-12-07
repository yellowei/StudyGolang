package functionDemo

import (
	"fmt"
	"testing"
)

// 闭包是函数与变量的结合
// 闭包可以看做一个类， 类有成员函数和成员变量
// 累加器
func AddUpper() func (int) int {

	var n int = 10

	// 返回一个匿名函数，匿名函数引用外部变量n
	// 因此，匿名函数和外部变量n形成了一个整体， 构成闭包
	return func(x int) int {
		n += x
		return n
	}
}

// 累加器的使用
func TestMyFunction10(t *testing.T) {

	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(1)) // 12
	fmt.Println(f(1)) // 13
}
