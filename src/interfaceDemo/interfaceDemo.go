// 接口
// 多态和高内聚低耦合

package interfaceDemo

import (
	"fmt"
)

// 定义一个接口，接口里的方法不允许有方法体
type Usb interface {
	Start()
	Stop()
}

// 实现接口， 必须实现接口的所有方法
// 接口的实现都是隐式实现
// 是基于方法列表的一种实现方式
type Phone struct {

}

func (u Phone) Start()  {
	fmt.Println("手机插入了")
}


func (u Phone) Stop()  {
	fmt.Println("手机拔出了")
}



type Camera struct {

}

func (u Camera) Start()  {
	fmt.Println("相机插入了")
}


func (u Camera) Stop()  {
	fmt.Println("相机拔出了")
}


// 调用接口

type Computer struct {

}

func (c Computer) Work(usb Usb)  {
	usb.Start()
	usb.Stop()
}
