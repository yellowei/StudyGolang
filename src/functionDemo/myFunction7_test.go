package functionDemo

import (
	"fmt"
	"testing"
)

// 变量的作用于

var name string = "tom" // 定义变量并初始化

// name2 := "tom"
// 上面这句报错， 是因为它等价于两句话：
// var name string  // 定义变量
// name = "tom"		// 变量赋值
// 在函数体外， 可以定义变量和初始化变量， 但不能写赋值语句

// 编译器的就近原则， 定义的局部变量和全局变量同名时，选取使用局部变量
func changeName () {

	name := "jack"

	fmt.Println("changeName", name)
}

func changeName2 () {

	name = "jack"

	fmt.Println("changeName2", name)
}

func TestMyFunction14(t *testing.T) {

	fmt.Println("TestMyFunction14", name) //tom

	changeName() //jack

	fmt.Println("TestMyFunction14", name) //tom

	changeName2() //jack

	fmt.Println("TestMyFunction14", name) //jack
}