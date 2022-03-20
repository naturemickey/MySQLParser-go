package tree

import (
	"reflect"
	"strings"
)

func Assemble(stat Stat, stats []Stat) {
	if len(stats) == 0 {
		return
	}

	object := reflect.ValueOf(stat)
	myref := object.Elem()
	typeOfType := myref.Type()

	j := 0
	complete := false
	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)

		fieldName := typeOfType.Field(i).Name

		if fieldName == "BaseStat" {
			continue
		}

		setMethodName := "Set" + strings.Title(fieldName)

		fieldType := field.Type()
		paramType := reflect.TypeOf(stats[j])

		if fieldType == paramType || paramType.AssignableTo(fieldType) {
			v := object.MethodByName(setMethodName)
			v.Call([]reflect.Value{reflect.ValueOf(stats[j])})
			j++
		}

		if j == len(stats) {
			complete = true
			break
		}
	}

	if !complete {
		panic("有些地方出错了。")
	}
}
