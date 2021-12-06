package myBreak

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 随机数示例
func TestMyBreakOne (t *testing.T) {

	var count int = 0

	for {
		count++
		//生成随机数需要设置一个种子
		rand.Seed(time.Now().UnixNano())
		//生成随机数
		var val = rand.Intn(100)
		fmt.Println(val)
		if val == 99 {
			break
		}
	}

	fmt.Printf("一共用了%d次\n", count)


}


// 标签的使用
func TestMyBreakTwo(t *testing.T) {
	label1:
	for i := 0; i < 4; i++ {
		//label2:
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label1
			}
			fmt.Println("j=", j)
		}
	}
	fmt.Println("跳出来")
}

// 练习一
// 100以内的数求和， 求出当前和等于20的当前数
func TestMyBreakThree(t *testing.T) {

	sum := 0
	for i := 1; i < 100; i++ {

		sum += i
		fmt.Printf("%d ", i)
		if sum >= 20 {
			break
		}
	}
}

// 练习二
// 某人有100000元，当经过路口时需要缴费
// 当现金>50000时，每次缴费5%
// 当现金<=50000时，每次缴费1000
// 请计算该人可以通过几次路口
func TestMyBreakFour(t *testing.T) {

	totalMoney := 100000.0
	count := 0
	for {
		if totalMoney < 1000 {
			fmt.Printf("总共可以通过%d次\n", count)
			fmt.Printf("还剩下的余额为%.2f元\n", totalMoney)
			break
		}
		if totalMoney > 50000 {
			totalMoney -= totalMoney * 0.05
		}else {
			totalMoney -= 1000
		}

		count++
	}
}