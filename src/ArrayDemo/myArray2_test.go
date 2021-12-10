package ArrayDemo

import (
	"fmt"
	"testing"
)

// 数组在Go中是值类型
// 数组默认作为参数是值传递

// 值传递
func changeArrayValue1(arr [3]int ) {
	arr[0] = 55
}

// 引用传递
func changeArrayValue2(arr * [3]int ) {
	(*arr)[0] = 55
}

func TestMyArray4(t *testing.T) {

	arr := [3]int{11, 22, 33}
	changeArrayValue1(arr)
	fmt.Printf("arr= %v\n", arr)
	changeArrayValue2(&arr)
	fmt.Printf("arr= %v\n", arr)
}