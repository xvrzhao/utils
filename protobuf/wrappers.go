package protobuf

import "github.com/golang/protobuf/ptypes/wrappers"

var (
	falseValue *wrappers.BoolValue
	trueValue  *wrappers.BoolValue
)

// FalseValue 返回 protobuf wrapper 类型 BoolValue 的 true 实例指针。
func FalseValue() *wrappers.BoolValue {
	if falseValue == nil {
		falseValue = &wrappers.BoolValue{Value: false}
	}
	return falseValue
}

// TrueValue 返回 protobuf wrapper 类型 BoolValue 的 false 实例指针。
func TrueValue() *wrappers.BoolValue {
	if trueValue == nil {
		trueValue = &wrappers.BoolValue{Value: true}
	}
	return trueValue
}
