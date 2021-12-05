package myFor

import (
	"fmt"
	"math"
	"testing"
)

// 答应九九乘法表
func TestForSix(t *testing.T) {

	for i := 1; i <= 9; i ++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(" %v x %v = %v ", j, i, j*i)
		}
		fmt.Println()
	}
}

// 获取水仙花数(一个三位数， 各位的立方和等于自身)
func TestForSeven (t *testing.T) {

	for i := 100; i < 1000; i++ {

		// 数学中求出各位上数字 n/10^(site-1)%10
		var onesPlace = i / 1 % 10
		var tensPlace = i / 10 % 10
		var handredsPlace = i / 100 % 10
		//fmt.Printf("%d\n",handredsPlace)
		//fmt.Printf("%d\n",tensPlace)
		//fmt.Printf("%d\n",onesPlace)

		var total = math.Pow(float64(onesPlace), 3.0) +
			math.Pow(float64(tensPlace), 3.0) +
			math.Pow(float64(handredsPlace), 3.0)

		if int(total) == i {
			fmt.Printf("%f\n", total)
		}
	}
}
