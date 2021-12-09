package myBuildIn

import (
	"fmt"
	"testing"
)

func TestMyBuildIn1(t *testing.T) {

	num1 := 100
	fmt.Printf("num1 type= %T, num1= %v, num1 address= %v\n", num1, num1, &num1)
}

func TestMyBuildIn2(t *testing.T) {

	// new()创建一个指向该类型的空值的指针
	num2 := new(int) // 指针
	fmt.Printf("num2 type= %T, num2= %v, num2 address= %v, *num2= %v\n", num2, num2, &num2, *num2)
	*num2 = 100
	fmt.Printf("num2 type= %T, num2= %v, num2 address= %v, *num2= %v\n", num2, num2, &num2, *num2)
}

func TestMyBuildIn3(t *testing.T) {

	// make()创建一个引用类型的指针
	amap := make(map[string]string)
	fmt.Printf("amap type= %T, amap= %v, amap address= %v\n", amap, amap, &amap)
	amap["first"] = "a1"
	amap["second"] = "b2"
	fmt.Printf("amap type= %T, amap= %v, amap address= %v\n", amap, amap, &amap)
}