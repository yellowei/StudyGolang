package myError

import (
	"errors"
	"fmt"
	"testing"
)

func testPanic() int {

	defer func() {
		err := recover()//从panic中恢复
		if err != nil {
			fmt.Println(err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2

	return res
}

func TestMyError1(t *testing.T) {

	res := testPanic()

	fmt.Println(res)
}

// 自定义错误
func readConfig(name string) error {
	if name == "config.ini" {
		 return nil
	}else {
		return errors.New("读取文件错误。。。。")
	}
}

func TestMyError2(t *testing.T) {

	err := readConfig("config1.ini")

	if err != nil {
		panic(err.Error())
	}

}
