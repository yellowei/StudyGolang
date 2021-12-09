package ArrayDemo

import (
	"fmt"
	"testing"
)

func TestMyArray1(t *testing.T) {

	var intArr [3]int
	// 数组的地址 就是 第一个元素的地址
	// 数组的第二个元素的地址，等于 第一个元素的地址 加上 数据类型的字节数
	fmt.Printf("intArr的类型=%T, intArr地址=%p, intArr[0]地址=%p, intArr[1]地址=%p\n", intArr, &intArr, &intArr[0], &intArr[1])
	var int32Arr [3]int32
	fmt.Printf("int32Arr地址=%p, int32Arr[0]地址=%p, int32Arr[1]地址=%p\n", &int32Arr, &int32Arr[0], &int32Arr[1])
}


// 四种初始化数组的方式
func TestMyArray2(t *testing.T) {

	var numArr0 = [3]int{1,2,3}
	fmt.Printf("%v\n", numArr0)

	var numArr1 [3]int = [3]int{1,2,3}
	fmt.Printf("%v\n", numArr1)

	var numArr2 = [...]int{1,2,3,4}
	fmt.Printf("%v\n", numArr2)

	var numArr3 = [...]int{2:100,3:200,300,400}
	fmt.Printf("%v\n", numArr3)
	var numArr4 = [...]int{2:100,6:200,0:300,4:400,7:500}
	fmt.Printf("%v\n", numArr4)
}

// for-range的使用
func TestMyArray3(t *testing.T) {

	var intArr = [...]int{1,2,3,4,5}

	for index, value := range intArr {

		fmt.Printf("intArr[%d]=%d\n", index, value)
	}
}