package mySwitch

import (
	"fmt"
	"go/types"
	"testing"
)

// switch 模拟 if-else-if结构
func TestMySwitchOne(t *testing.T) {

	var score float64 = 90.0

	fmt.Println(score)

	switch  {
	case score >= 90:
		fmt.Println("成绩优秀")
	case score >= 80:
		fmt.Println("成绩优良")
	case score >= 70:
		fmt.Println("成绩良好")
	case score >= 60:
		fmt.Println("成绩及格")
	case score >= 0:
		fmt.Println("成绩不及格")

	}
}

// type-switch 的使用

func TestSwitchTwo(t *testing.T) {

	//var y int
	var y1 types.Nil
	y := y1
	var x interface{}
	x = y

	switch i := x.(type) {
	case nil:
		fmt.Printf("这是%T类型\n",i)
	case int:
		fmt.Println("这是int类型")
	case float64:
		fmt.Println("这是float64类型")
	case func(int) float64:
		fmt.Println("这是func(int) float64类型")
	case bool, string:
		fmt.Println("这是bool或者string类型")
	default:
		fmt.Printf("未知类型: %T\n", i)
	}

}