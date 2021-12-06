package functionDemo

import (
	"fmt"
	"testing"
)

// 递归测试
// n的值传递， 不会因为在被调用的函数体内改变而改变
func test (n int) {

	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}

func test2 (n int) {

	if n > 2 {
		n--
		test2(n)
	}else {
		//这句话的执行时机 体现了对函数形参的理解
		fmt.Println("n=", n)
	}

}

func TestFunctionOne(t *testing.T) {

	n := 10
	test2(n)
	fmt.Println("n=", n)
}


// 斐波那契数列
func fibonacciSequence (n int) int {
	if n <= 2 {
		return 1
	}else {
		res := fibonacciSequence(n-1) + fibonacciSequence(n-2)
		//fmt.Println(res)
		return res
	}
}

func TestMyFunctionTwo(t *testing.T) {

	fmt.Println(fibonacciSequence(1))
	fmt.Println(fibonacciSequence(2))
	fmt.Println(fibonacciSequence(3))
	fmt.Println(fibonacciSequence(4))
	fmt.Println(fibonacciSequence(5))
	fmt.Println(fibonacciSequence(6))
	fmt.Println(fibonacciSequence(7))
	fmt.Println(fibonacciSequence(8))
	fmt.Println(fibonacciSequence(9))
}

// f(1) = 3; f(n) = 2 * f(n - 1) + 1
func f (n int) int {
	if n == 1 {
		return 3
	}else {
		return 2 * f(n - 1) + 1
	}
}

func TestMyFunctionThree(t *testing.T) {

	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
	fmt.Println(f(5))
	fmt.Println(f(6))
	fmt.Println(f(7))
	fmt.Println(f(8))
	fmt.Println(f(9))
}

// 猴子吃桃问题
// 有一堆桃子， 第一天吃了其中的一半再多吃一个，第二天也如此， 直到第十天想吃时只剩一个了

// 正向思维
func f3 (n int) int  {
	if n >= 10  {
		return 1
	} else if n < 10 && n > 0 {
		return  2 * ( f3(n+1) + 1 )
	} else {
		return 0
	}
}

// 逆向思维
func  f2 (n int) int  {

	if n == 1 {
		return 1
	}else {
		return  2 * (f2(n-1) + 1)
	}
}

func TestMyFunctionFour(t *testing.T) {
	fmt.Println(f2(10))
	fmt.Println(f3(1))
}