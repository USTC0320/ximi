package ximi

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// StructCopy 结构体复制
// source 当前有值的结构体
// target 接受值的结构体
// fields 需要的设置的属性
func StructCopy(source interface{}, target interface{}, fields ...string) (err error) {
	sourceKey := reflect.TypeOf(source)
	sourceVal := reflect.ValueOf(source)

	targetKey := reflect.TypeOf(target)
	targetVal := reflect.ValueOf(target)

	if targetKey.Kind() != reflect.Ptr {
		err = fmt.Errorf("被覆盖的数据必须是一个结构体指针")
		return
	}

	targetVal = reflect.ValueOf(targetVal.Interface())

	// 存放字段
	fieldItems := make([]string, 0)

	if len(fields) > 0 {
		fieldItems = fields
	} else {
		for i := 0; i < sourceVal.NumField(); i++ {
			fieldItems = append(fieldItems, sourceKey.Field(i).Name)
		}
	}

	for i := 0; i < len(fieldItems); i++ {
		field := targetVal.Elem().FieldByName(fieldItems[i])
		value := sourceVal.FieldByName(fieldItems[i])
		if field.IsValid() && field.Kind() == value.Kind() {
			field.Set(value)
		}

	}
	return

}

func TypeChange(source string, splitStr string) []int {
	tempArr := strings.Split(source, splitStr)
	result := make([]int, 0)
	for _, item := range tempArr {
		newItem, _ := strconv.Atoi(item)
		result = append(result, newItem)
	}
	return result
}

// IsPathExists 判断文件是否存在
func IsPathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func IsEmpty(arg interface{}) bool {
	switch reflect.TypeOf(arg).Kind().String() {
	case "int":
		if arg == 0 {
			return true
		} else {
			return false
		}
	case "string":
		if arg == "" {
			return true
		} else {
			return false
		}
	case "int64":
		if arg == 0 {
			return true
		} else {
			return false
		}
	case "uint8":
		if arg == false {
			return true
		} else {
			return false
		}
	case "float64":
		if arg == 0.0 {
			return true
		} else {
			return false
		}
	case "byte":
		if arg == 0 {
			return true
		} else {
			return false
		}
	case "ptr":
		//反射判空逻辑
		if reflect.ValueOf(arg).IsNil() { //利用反射直接判空
			return true
		} else {
			return false
		}
	case "struct":
		if arg == nil {
			return true
		} else {
			return false
		}
	case "slice":
		s := reflect.ValueOf(arg)
		if s.Len() == 0 {
			return true
		} else {
			return false
		}
	case "array":
		s := reflect.ValueOf(arg)
		if s.Len() == 0 {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}
