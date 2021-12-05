package mySwitch

import (
	"fmt"
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