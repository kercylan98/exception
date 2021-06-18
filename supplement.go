package exception

import (
	"fmt"
	"runtime"
)

// supplement 对例外信息对补充项
type supplement struct {
	name 	string				// 补充项名称
	value 	interface{}			// 补充项具体值（如果实现了Serialize接口将序列化显示）
	stack 	string				// 堆栈信息
}

func newSupplement(name string, value interface{}) supplement {
	sup := supplement{
		name:  name,
		value: value,
	}

	pc, file, line, ok := runtime.Caller(2)
	if ok {
		if serialize, ok := value.(Serialize); ok {
			sup.stack = fmt.Sprintf("\t%s (%v) at %s:%d (Method %s)", name, serialize.Serialize(), file, line, runtime.FuncForPC(pc).Name())
		}else {
			sup.stack = fmt.Sprintf("\t%s (%v) at %s:%d (Method %s)", name, value, file, line, runtime.FuncForPC(pc).Name())
		}
	}

	return sup
}

// Name 获取补充项的名称
func (slf *supplement) Name() string {
	return slf.name
}

// Value 获取补充项的值
func (slf *supplement) Value() interface{} {
	return slf.value
}

// Stack 获取该行补充项的堆栈信息
func (slf *supplement) Stack() string {
	return slf.stack
}