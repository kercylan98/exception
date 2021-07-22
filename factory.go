package exception

import (
	"fmt"
	"runtime"
)

// Factory 例外信息工厂
type Factory interface {
	// Hit 命中例外情况，产生并返回一个例外信息
	Hit() Exception
	// Equal 比较两个例外是否相同
	Equal(exception Exception) bool
}

func newFactory(code int, text string, feedback ...string) Factory {
	var defaultFeedback = "server internal exception"
	if len(feedback) != 0 {
		defaultFeedback = feedback[0]
	}
	return &factory{
		code:     code,
		text:     text,
		feedback: defaultFeedback,
	}
}

type factory struct {
	code     int    // 对应例外情况错误码
	text     string // 对应例外情况文本
	feedback string // 用户反馈信息
}

func (slf *factory) Equal(exception Exception) bool {
	return slf.code == exception.Code()
}

func (slf *factory) Hit() Exception {
	var exp = newException(slf.code, slf.text, slf.feedback).(*exception)
	var skip = 1
	var lines string
	for {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		lines = lines + fmt.Sprintf("\n\tat %s:%d (Method %s)", file, line, runtime.FuncForPC(pc).Name())
		skip++
	}
	exp.stack = lines[1:]
	if app.hitMiddlewares != nil {
		for i := 0; i < len(app.hitMiddlewares); i++ {
			app.hitMiddlewares[i](exp)
		}
	}
	return exp
}
