package ArrayDemo

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 输出A-Z
func TestMyArray5(t *testing.T) {

	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}
	fmt.Printf("myChars= %c\n,",myChars)
}

// 输出A-Z
func TestMyArray6(t *testing.T) {

	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = byte( int('A') + i)
	}
	fmt.Printf("myChars= %c\n,",myChars)
}

// 求出数组中的最大值
func TestMyArray7(t *testing.T) {

	var max = 0
	var maxIndex = -1
	var intArr = [5]int{-100, -200, -300, -400, -500}
	for i := 0; i < len(intArr); i++ {
		if i == 0 {
			max = intArr[i]
			maxIndex = i
		}
		if intArr[i] > max {
			max = intArr[i]
			maxIndex = i
		}
	}
	if maxIndex >= 0 {
		fmt.Printf("找到最大值max= %v，下标为index= %v \n,",max, maxIndex)
	}

}

// 求数组的和与平均值
func TestMyArray8(t *testing.T) {
	var avg,sum float64
	var arr = [5]int{-1001, -200, -300, -400, -500}
	for _, v := range arr {
		sum += float64(v)
	}
	avg = sum / float64(len(arr))
	fmt.Printf("和= %v，平均值x= %v \n,",sum, avg)
}

// 随机生成五个数， 反转打印
func TestMyArray9(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	var intArr [5]int
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Printf("%v\n", intArr)

	for i := 4; i >= 0; i-- {
		fmt.Printf("%v ", intArr[i])
	}
	fmt.Printf("\n")
}

// 随机生成五个数， 反转数组
func TestMyArray10(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	var intArr [5]int
	var lenArr = len(intArr)
	for i := 0; i < lenArr; i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Printf("%v\n", intArr)

	// 交换数组
	for i := 0; i < lenArr / 2; i++ {
		temp := intArr[i]
		intArr[i] = intArr[lenArr - 1 - i]
		intArr[lenArr - 1 - i] = temp
	}
	fmt.Printf("%v\n", intArr)
}