package functionDemo

import (
	"fmt"
	"strings"
	"testing"
)

// 1)编写一个函数makeSuffix(suffix string) 可以接收一个文件后缀名（如.jpg），并返回一个闭包
// 2)调用闭包，可以传入一个文件名，如果该文件名没有后缀，则返回文件名.jpg，如果有后缀了，就直接返回文件名
// 3)要求使用闭包的方式完成
// 4)判断字符串尾部是否有子字符串，用strings.HasSuffix

func makeSuffix(suffix string) func (name string) string {

	return func(name string) string {
		
		// 如果name没有指定后缀名，则加上
		if !strings.HasSuffix(name, suffix) {
			return name + "." + suffix
		}
		return name
	}
}

func makeSuffix2(suffix, name string) string {
	// 如果name没有指定后缀名，则加上
	if !strings.HasSuffix(name, suffix) {
		return name + "." + suffix
	}
	return name
}

func TestMyFunction11(t *testing.T) {

	res1 := makeSuffix("jpg")("一张图片.png")
	fmt.Println(res1)

	jpgSuffix := makeSuffix("jpg")
	pngSuffix := makeSuffix("png")
	bmpSuffix := makeSuffix("bmp")
	fmt.Println(jpgSuffix("a"))
	fmt.Println(pngSuffix("a"))
	fmt.Println(bmpSuffix("a"))

	fmt.Println(makeSuffix2("jpg", "b"))
	fmt.Println(makeSuffix2("png", "b"))
	fmt.Println(makeSuffix2("bmp", "b"))
}