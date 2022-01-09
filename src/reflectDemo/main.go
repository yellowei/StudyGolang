package main

import (
	"fmt"
	"reflect"
	at "reflectDemo/aboutGoTest"
	"unsafe"
)

func main () {
	// 错误示范
	//fmt.Printf("testStudent.name: %v\n", at.TestStudent.name)

	if false {
		// 取地址访问
		fmt.Printf("testStudent.name: %v\n", GetPtrUnExportFiled(&at.TestStudent, "name"))
		if err := SetPtrUnExportFiled(&at.TestStudent, "name", "kangkang"); err != nil {
			fmt.Println("err: %v", err)
		}

		fmt.Printf("testStudent.name: %v\n", GetPtrUnExportFiled(&at.TestStudent, "name"))

	} else {

		// 非取地址
		fmt.Printf("testStudent.name: %v\n", GetUnExportFiled(at.TestStudent, "name"))

		if ret, err := SetUnExportedField(at.TestStudent, "name", "kangkang"); err != nil {
			fmt.Println("err: %v", err)
		} else {

			fmt.Printf("testStudent.name: %v\n", GetUnExportFiled(at.TestStudent, "name"))

			// 通过反射转换成为结构体类型
			v := ret.Interface().(at.Student)
			fmt.Printf("ret.name: %v\n", GetUnExportFiled(v, "name"))

		}
	}



}


// 取地址访问
func GetPtrUnExportFiled(ptr interface{}, fieldName string) reflect.Value {
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	return v
}

func SetPtrUnExportFiled(ptr interface{}, fieldName string, newFieldVal interface{}) (err error) {
	// 获取非导出字段反射对象
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	// 获取非导出字段可寻址反射对象
	// 与上面的区别是：这个是可寻址的
	v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()

	nv := reflect.ValueOf(newFieldVal)
	if v.Kind() != nv.Kind() {
		return fmt.Errorf("expected kind %v, got kind: %v", v.Kind(), nv.Kind())
	}
	v.Set(nv)
	return nil
}

// reflect 非取地址访问和修改非导出字段
func GetUnExportFiled(source interface{}, fieldName string) (ret reflect.Value) {
	v := reflect.ValueOf(source)

	vptr := reflect.New(v.Type()).Elem()
	vptr.Set(v)

	ret = vptr.FieldByName(fieldName)
	return
}

func SetUnExportedField(source interface{}, fieldName string, newFieldVal interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(source)

	vptr := reflect.New(v.Type()).Elem()
	vptr.Set(v)

	tv := vptr.FieldByName(fieldName)
	tv = reflect.NewAt(tv.Type(), unsafe.Pointer(tv.UnsafeAddr())).Elem()

	nv := reflect.ValueOf(newFieldVal)
	if tv.Kind() != nv.Kind() {
		return v, fmt.Errorf("expected kind %v, got kind: %v", v.Kind(), nv.Kind())
	}
	tv.Set(nv)
	return vptr, nil
}