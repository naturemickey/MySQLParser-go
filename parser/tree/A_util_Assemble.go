package tree

import (
	"reflect"
	"strings"
)

type BaseAssemble struct {
}

func (this *BaseAssemble) Assemble(stats []Stat) {
	if len(stats) == 0 {
		return
	}

	object := reflect.ValueOf(this)
	myref := object.Elem()
	typeOfType := myref.Type()

	j := 0
	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)

		fieldName := typeOfType.Field(i).Name

		setMethodName := "Set" + strings.Title(fieldName)

		if field.Type() == reflect.TypeOf(stats[j]) {
			v := object.MethodByName(setMethodName)
			v.Call([]reflect.Value{reflect.ValueOf(stats[j])})
			j++
		}

	}

	if j != len(stats)-1 {
		panic("有些地方出错了。")
	}
}
