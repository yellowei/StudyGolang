package functionDemo

import "fmt"

// 全局变量先于init()执行
var age = testAge()
var Name = 90

func testAge () int {
	fmt.Println("testAge()")
	return 80
}

func init()  {
	fmt.Println("init()")
}