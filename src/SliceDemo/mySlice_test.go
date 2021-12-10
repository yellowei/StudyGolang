package SliceDemo

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestMySlice1(t *testing.T) {

	var intArr [7]int = [7]int{1, 22, 33, 44, 55, 66, 77}

	// 定义一个切片
	// 取intArr的下标为1到下标为4-1的元素作为切片的元素
	var intSlice = intArr[2:3]
	fmt.Printf("slice的元素= %v， slice的容量= %v, slice的元素个数= %v\n", intSlice, cap(intSlice), len(intSlice))

}

// 通过make的方式创建切片
// 切片扩展时，容量 = 长度 * 2
func TestMySlice2(t *testing.T) {

	var intSlice []int = make([]int, 5, 5)

	fmt.Printf("%v\n", intSlice)
	fmt.Printf("slice的元素= %v， slice的容量= %v, slice的元素个数= %v\n", intSlice, cap(intSlice), len(intSlice))

	intSlice = append(intSlice, 1)

	fmt.Printf("%v\n", intSlice)
	fmt.Printf("slice的元素= %v， slice的容量= %v, slice的元素个数= %v\n", intSlice, cap(intSlice), len(intSlice))
}

// 切片扩展后的内存地址
func TestMySlice3(t *testing.T) {

	// 这个超大容量的切片理论上可以消耗将近8G的内存空间
	var intSlice []int = make([]int, 8 * 125 * 1000 * 1000, 8 * 125 * 1000 * 1000)
	for i := 0; i < len(intSlice); i++ {
		fmt.Printf("intSlice[%d] = %p\n", i, &intSlice[i])
	}

	intSlice = append(intSlice, 1)

	for i := 0; i < len(intSlice); i++ {
		fmt.Printf("intSlice[%d] = %p\n", i, &intSlice[i])
	}
}

// 切片扩展后的内存地址校验
func TestMySlice4(t *testing.T) {

	var intSlice []int = make([]int, 8 * 125 * 1000, 8 * 125 * 1000)
	for i := 0; i < len(intSlice); i++ {
		if i < len(intSlice) - 1 {
			ptr1 := strings.TrimLeft(fmt.Sprintf("%p", &intSlice[i]), "0x")
			ptr2 := strings.TrimLeft(fmt.Sprintf("%p", &intSlice[i+1]), "0x")
			ptr1Int, err1 := strconv.ParseInt(ptr1, 16, 0)
			if err1 != nil {
				fmt.Printf("%v\n", err1)
			}
			ptr2Int, err2 := strconv.ParseInt(ptr2, 16, 0)
			if err2 != nil {
				fmt.Printf("%v\n", err2)
			}
			res := ptr2Int - ptr1Int
			if res != 8 {
				fmt.Printf("%v\n", res)
			}

		}
	}

	intSlice = append(intSlice, 1)
	//fmt.Printf("slice的元素= %v， slice的容量= %v, slice的元素个数= %v\n", intSlice, cap(intSlice), len(intSlice))

	for i := 0; i < len(intSlice); i++ {
		if i < len(intSlice) - 1 {
			ptr1 := strings.TrimLeft(fmt.Sprintf("%p", &intSlice[i]), "0x")
			ptr2 := strings.TrimLeft(fmt.Sprintf("%p", &intSlice[i+1]), "0x")
			ptr1Int, err1 := strconv.ParseInt(ptr1, 16, 0)
			if err1 != nil {
				fmt.Printf("%v\n", err1)
			}
			ptr2Int, err2 := strconv.ParseInt(ptr2, 16, 0)
			if err2 != nil {
				fmt.Printf("%v\n", err2)
			}
			res := ptr2Int - ptr1Int
			if res != 8 {
				fmt.Printf("%v\n", res)
			}
		}
	}
}