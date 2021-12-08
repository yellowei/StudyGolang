package myString

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 内建函数len()的使用
func TestMyString1(t *testing.T) {

	str := "hello中国"
	fmt.Println("str len=", len(str))

}

// 含有中文字符的字符串遍历
func TestMyString2(t *testing.T) {
	str2 := "hello中国"
	str3 := []rune(str2)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("%c\n", str3[i])
	}
}

// 字符串转整数
func TestMyString3(t *testing.T) {
	str1 := "123.1"
	int1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println("转换错误", err.Error())
	}else {
		fmt.Println(int1)
	}
}

// 整数转字符串
func TestMyString4(t *testing.T) {
	a := 12345
	str := strconv.Itoa(a)
	fmt.Printf("str= %v, type= %T\n", str, str)
}

// 字符串转 []byte
func TestMyString5(t *testing.T) {

	var str = "hello go"
	bytes := []byte(str)
	fmt.Printf("bytes= %v, type= %T\n", bytes, bytes)
}

// []byte转字符串
func TestMyString6(t *testing.T) {
	bytes := []byte{97,98,99}
	str := string(bytes)
	fmt.Printf("bytes= %v, type= %T\n", str, str)
}

// 10进制数转2，8，16进制数
func TestMyString7(t *testing.T) {

	str2 := strconv.FormatInt(123, 2)
	str8 := strconv.FormatInt(123, 8)
	str16 := strconv.FormatInt(123, 16)
	fmt.Printf("bytes= %v, type= %T\n", str2, str2)
	fmt.Printf("bytes= %v, type= %T\n", str8, str8)
	fmt.Printf("bytes= %v, type= %T\n", str16, str16)
}

// 判断字符串是否包含子串
func TestMyString8(t *testing.T) {

	str1 := "hello go"
	str2 := "go"
	isContain := strings.Contains(str1,str2)
	fmt.Printf("isContain= %v, isContain= %T\n", isContain, isContain)
}

// 返回字符串中包含子串的个数
func TestMyString9(t *testing.T) {

	str1 := "hello go"
	str2 := "go"
	count := strings.Count(str1,str2)
	fmt.Printf("count= %v, count= %T\n", count, count)
}


// 字符串比较
func TestMyString10(t *testing.T) {

	str1 := "Go"
	str2 := "go"
	// 区分大小写
	isEqual := str1 == str2
	fmt.Printf("isEqual= %v, isEqual= %T\n", isEqual, isEqual)
	// 不区分大小写
	isEqual   = strings.EqualFold(str1,str2)
	fmt.Printf("isEqual= %v, isEqual= %T\n", isEqual, isEqual)
}

// 获取子串在字符串中第一次出现的下标
func TestMyString11(t *testing.T) {
	str1 := "hello go go"
	str2 := "go"
	index := strings.Index(str1, str2)
	if index == -1 {
		fmt.Printf("str1不包含str2")
	}else {
		fmt.Printf("str2在str1中第一次出现的下标为 %v", index)
	}
	fmt.Printf("index= %v, index= %T\n", index, index)
}


// 获取子串在字符串中最后一次出现的下标
func TestMyString12(t *testing.T) {
	str1 := "hello go go"
	str2 := "go"
	index := strings.LastIndex(str1, str2)
	if index == -1 {
		fmt.Printf("str1不包含str2")
	}else {
		fmt.Printf("str2在str1中第一次出现的下标为 %v", index)
	}
	fmt.Printf("index= %v, index= %T\n", index, index)
}


// 替换字符串中指定的子串
func TestMyString13(t *testing.T) {
	str1 := "hello go go"
	str2 := "go语言"
	// 最后一个参数表示替换几个子串， -1 表示替换所有
	strNew := strings.Replace(str1, "go", str2, 1)
	fmt.Printf("strNew= %v, strNew= %T\n", strNew, strNew)
}

// 字符串分割
func  TestMyString14(t *testing.T) {

	strArr := strings.Split("hello,go,go", ",")
	fmt.Println(strArr)

	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
}

// 大小写转换
func TestMyString15 (t *testing.T) {
	strLower := strings.ToLower("GO")
	fmt.Println(strLower)
	strUpper := strings.ToUpper("go")
	fmt.Println(strUpper)
}

// 去掉字符串左右两边的空格
func TestMyString16(t *testing.T) {

	str := "  hello go  "
	str = strings.TrimSpace(str)
	fmt.Println(str)
}

// 去掉字符串左右两边指定的字符
func TestMyString17(t *testing.T) {

	str := "!' !a hello !go!aaa  !"
	str = strings.Trim(str, "!")
	fmt.Println(str)
}

// 去掉字符串左边指定的字符
func TestMyString18(t *testing.T) {

	str := "!' !a hello !go!aaa  !"
	str = strings.TrimLeft(str, "!") //只去掉左边
	fmt.Println(str)
}

// 去掉字符串右边指定的字符
func TestMyString19(t *testing.T) {

	str := "!' !a hello !go!aaa  !"
	str = strings.TrimRight(str, "!")//只去掉右边
	fmt.Println(str)
}

// 判断字符串是否以指定子串开头
func TestMyString20(t *testing.T) {
	str := "http://127.0.0.1"
	isBegin := strings.HasPrefix(str, "http")
	if isBegin {
		fmt.Printf("是否以http开头：%v\n", isBegin)
	}
}

// 判断字符串是否以指定子串结尾
func TestMyString21(t *testing.T) {
	str := "abcd.jpg"
	isEnd := strings.HasSuffix(str, ".jpg")
	if isEnd {
		fmt.Printf("是否以.jpg结尾：%v\n", isEnd)
	}
}

