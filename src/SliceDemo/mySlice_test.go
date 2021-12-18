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

// 切片拷贝
func TestMySlice5(t *testing.T) {

	var a = []int{1,2,3,4,5}
	var b = make([]int, 10)

	copy(b, a)

	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)


	// 拷贝时，目标切片容量可大可小都不会报错
	var c = []int{1,2,3,4,5}
	var d = make([]int, 1)

	copy(d, c)

	fmt.Printf("c=%v\n", c)
	fmt.Printf("d=%v\n", d)
}

// 切片与切片对应数组的关系
func TestMySlice6(t *testing.T) {

	var a = [5]int{1,2,3,4,5}
	var slice = a[:]
	var slice2 = slice
	slice2[0] = 10

	fmt.Printf("a=%v\n", a)
	fmt.Printf("slice=%v\n", slice)
	fmt.Printf("slice2=%v\n", slice2)
}

// 切片作为参数传递，修改元素时生效，因为发生的是引用传递
func changeSlice(slice []int) {
	slice[0] = 10//此处修改会修改实参
}
func TestMySlice7(t *testing.T) {

	var a = []int{1,2,3,4,5}
	fmt.Printf("a=%v\n", a)
	changeSlice(a)
	fmt.Printf("a=%v\n", a)
}

// string与切片
func TestMySlice8(t *testing.T) {

	str := "hello@yellowei"

	slice := str[6:]

	println(slice)
}

// string是不可变的， 修改字符串需要先转为切片操作后再转回
// 中文字符串转为[]rune
func TestMySlice9(t *testing.T) {

	str := "hello@yellowei"

	//str[0] = "H"  //这里编译报错
	strByte := []byte(str)
	strByte[0] = 'H'
	str = string(strByte)
	println(str)
}

func TestMySlice10(t *testing.T) {

	str := "你好深圳"

	//str[0] = "H"  //这里编译报错
	strRune := []rune(str)
	strRune[0] = '您'
	str = string(strRune)
	println(str)
}
