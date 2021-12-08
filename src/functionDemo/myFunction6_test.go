package functionDemo

import (
	"fmt"
	"testing"
)

// defer的使用
// defer 的执行顺序是按照栈的先入后出的顺序出栈"执行"
// 当前函数执行完成之后，就开始出栈执行当前函数中defer
func sum (n1, n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

// 此例子说明，当defer入栈时，会把当前相关的变量的值拷贝入栈
// 也就是说，入栈时变量是多少，即使当前函数后续修改了变量的值，也不会影响已入栈的值
func sum2 (n1, n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}


func TestMyFunction12(t *testing.T) {

	res := 1

	defer fmt.Println("ok4 res=", res)

	res = sum(10, 20)

	defer fmt.Println("ok5 res=", res)
}

func TestMyFunction13(t *testing.T) {

	res := sum2(10, 20)
	fmt.Println("ok5 res=", res)
}