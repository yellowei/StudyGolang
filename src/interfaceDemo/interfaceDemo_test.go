package interfaceDemo

import (
	"testing"
)


func Test (t *testing.T) {
	var phone Phone
	var camera Camera
	var computer Computer

	computer.Work(phone)
	computer.Work(camera)

}