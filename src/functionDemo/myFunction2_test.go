package functionDemo

import (
	"fmt"
	"testing"
)

// 函数别名
type myNameGetSum func(int, int) int

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数别名
type myNameFunc func( func(int, int) int, int, int) int

type myNameFunc2 func( myNameGetSum, int, int) int

// 函数作为参数传递
func myFunc(funVar func(int, int) int, num1, num2 int) int {
	return funVar(num1, num2)
}

func myFunc2(funVar myNameGetSum, num1, num2 int) int {
	return funVar(num1, num2)
}

// 函数返回值指定返回
func getSumAndSub (n1, n2 int) (int, int) {
	return n1+n2, n1-n2
}
// 函数返回值命名后自动返回
func getSumAndSub2 (n1, n2 int) (sum int, sub int) {
	sum = n1+n2
	sub = n1-n2
	return
}

// 函数可变参数写法, 支持0到多个参数
func myFunc3 (args... int) (sum int) {
	for index,_ := range args {
		sum += args[index]
	}

	return
}

// 函数可变参数写法, 支持1到多个参数
func myFunc4 (n1 int ,args... int) (sum int) {
	sum = n1
	for index,_ := range args {
		sum += args[index]
	}
	return
}

// 引用传递示例
func swap(a * int, b * int) {
	*a,*b = *b,*a
}


 func TestMyFunctionFive(t *testing.T) {

	 res := myFunc2(getSum,1, 2)
	 fmt.Println(res)

	 res2 := myFunc3(1,2,3,4,5,6)
	 fmt.Println(res2)

	 res3 := myFunc4(1,2,3,4,5,6)
	 fmt.Println(res3)

	 a := 1
	 b := 2
	 fmt.Println(a,b)
	 swap(&a,&b)
	 fmt.Println(a,b)
 }
