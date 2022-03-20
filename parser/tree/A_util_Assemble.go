package tree

import (
	"reflect"
	"strings"
)

func Assemble(stat Stat, stats []Stat) {
	if len(stats) == 0 {
		return
	}

	reflectVal := reflect.ValueOf(stat)
	myref := reflectVal.Elem()
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
			method := reflectVal.MethodByName(setMethodName)
			method.Call([]reflect.Value{reflect.ValueOf(stats[j])})
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

func Children(stat Stat) []Stat {
	var res []Stat

	reflectVal := reflect.ValueOf(stat)
	myref := reflectVal.Elem()
	typeOfType := myref.Type()
	for i := 0; i < myref.NumField(); i++ {
		fieldName := typeOfType.Field(i).Name

		if fieldName == "BaseStat" {
			continue
		}

		getMethodName := strings.Title(fieldName)
		method := reflectVal.MethodByName(getMethodName)
		fieldVal := method.Call([]reflect.Value{})

		val := fieldVal[0].Interface()

		if val == nil || reflect.ValueOf(val).IsNil() {
			continue
		}

		switch val.(type) {
		case Stat:
			res = append(res, val.(Stat))
		}
	}
	return res
}
