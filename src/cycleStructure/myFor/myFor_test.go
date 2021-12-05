package myFor

import (
	"fmt"
	"testing"
)

func TestForStepOne(t *testing.T) {

	// 打印矩形
	for i := 1; i <= 4; i ++ {
		for j := 1; j <= 4; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func TestForStepTwo(t *testing.T) {

	// 打印半金字塔
	//for i := 1; i <= 4; i ++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	// 打印半金字塔
	for i := 1; i <= 4; i ++ {
		for j := 1; j <= 4; j++ {
			if j <= i {
				fmt.Print("*")
			}

		}
		fmt.Println()
	}
}



func TestForStepThree(t *testing.T) {

	//金字塔层数
	var level int = 20
	// 打印金字塔
	for i := 1; i <= level; i ++ {
		for j := 1; j <= (2 * level - 1); j++ {
			if j <= level - i || j >= (level + i)  {
				fmt.Print(" ")
			} else  {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func TestForStepThree_Another(t *testing.T) {

	//金字塔层数
	var level int = 20
	// 打印金字塔
	for i := 1; i <= level; i ++ {
		for j := 1; j <= (2 * level - i); j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= (2 * i - 1); j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func TestForStepFour(t *testing.T) {

	//金字塔层数
	var level int = 20
	// 打印金字塔
	for i := 1; i <= level; i ++ {
		for j := 1; j <= (2 * level - 1); j++ {
			if j == level - i || j == (level + i)  {
				fmt.Print(" ")
			} else  {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func TestForStepFive(t *testing.T) {

	//金字塔层数
	var level int = 20
	// 打印金字塔
	for i := 1; i <= level; i ++ {
		for j := 1; j <= (2 * level - 1); j++ {
			if j == level - i + 1 || j == level + i - 1 || i == level {
				fmt.Print("*")
			} else  {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}