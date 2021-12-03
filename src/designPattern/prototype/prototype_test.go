package prototype

import (
	"fmt"
	"testing"
)

var manager *PrototypeManager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	fmt.Printf("Clone %p\n",t)
	fmt.Println(*t)
	fmt.Println(&t)
	fmt.Printf("Clone %p\n",&t)
	fmt.Printf("Clone %p\n",&tc)
	fmt.Println(tc)
	if tc == *t {
		fmt.Println("tc == *t")
	}
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")
	fmt.Printf("t1 p %p\n",&t1)
	t2 := t1.Clone()
	fmt.Printf("t2 p %p\n",&t2)
	t3 := t1
	fmt.Printf("t3 p %p\n",&t3)
	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}else {
		fmt.Println(&t1)
		fmt.Println(&t2)
		fmt.Printf("t1 %v\n",t1)
		fmt.Printf("t2 %v\n",t2)
	}
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t1").Clone()
	t1 := c.(*Type1)
	if t1.name != "type1" {
		t.Fatal("error")
	}

}

func init() {
	manager = NewPrototypeManager()

	t1 := &Type1{
		name: "type1",
	}
	fmt.Printf("%p\n",&t1)
	fmt.Println(&t1)
	manager.Set("t1", t1)

	t2 := manager.Get("t1")
	fmt.Println(&t2)
	if t1 == t2 {
		fmt.Println("manager t1")
	}
}