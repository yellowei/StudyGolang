package myTime

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)


func TestMyTime1 (t *testing.T) {
	// 1.获取当前时间
	tm := time.Now()
	fmt.Printf("tm= %v, type= %T\n", tm, tm)

	// 2.获取其它日期信息， 年月日，时分秒
	year, mon, day := tm.Date()
	hor, min, sec := tm.Clock()
	fmt.Printf("%v年%v月%v日%v时%v分%v秒\n",
		year, int(mon), day, hor, min, sec)

	// 3. 格式化日期和时间
	// 方式一 printf  Sprintf  直接输出
	fmt.Printf("%v-%v-%v %v:%v:%v\n",
		year, int(mon), day, hor, min, sec)

	// 方式二 time.Format 方法
	fmt.Println(tm.Format("2006/01/02 15:04:05"))
	fmt.Println(tm.Format("2006-01-02 15:04:05"))
	fmt.Println(tm.Format("2006"))
	fmt.Println(tm.Format("01"))
	fmt.Println(tm.Format("02"))
	fmt.Println(tm.Format("15"))
	fmt.Println(tm.Format("04"))
	fmt.Println(tm.Format("05"))

	// 时间常量
	//time.Nanosecond
	//time.Microsecond
	//time.Millisecond
	//time.Second
	//time.Minute
	//time.Hour



}

// 休眠，结合时间常量使用
func TestMyTime2 (t *testing.T) {

	var i = 0
	for {
		i++
		if i == 100 {
			break
		}

		fmt.Println(i)
		//time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 100)  // 0.1秒的表示，  不能用除法
	}
}


// 获取随机数
func TestMyTime3 (t *testing.T) {

	var unixNano = time.Now().UnixNano() // 纳秒时间戳
	var unix = time.Now().Unix() // 时间戳
	fmt.Println(unixNano, unix)
	rand.Seed(unixNano) // 设置种子的随机性， 随机事件纳秒级
	var randint = rand.Intn(100)
	fmt.Println(randint)

}

// 统计一段代码的执行时间
func TestMyTime4(t *testing.T) {

	fmt.Println("start")
	startTime := time.Now().Unix()
	str := "hello"
	for i := 0; i < 200000; i++ {
		str += strconv.Itoa(i)
	}
	//fmt.Printf("%v\n", str)
	endTime := time.Now().Unix()
	fmt.Println("end")

	duration := endTime - startTime
	fmt.Printf("%v\n", duration)
}
