package exception

import (
	"errors"
	"fmt"
)

// Exception 例外情况错误信息实现
type Exception interface {
	// Code 获取错误码
	Code() int
	// Print 打印错误信息
	Print()
	// Supplement 补充错误信息
	Supplement(name string, content interface{}) Exception
	// Feedback 获取可以用于反馈给用户对信息
	Feedback() string
	error
}

func newException(code int, text string, feedback string) Exception {
	return &exception{
		code: code,
		err: errors.New(text),
		supplements: []supplement{},
		feedback: feedback,
	}
}

type exception struct {
	code 	int 					// 错误代码
	err 	error					// 原生错误信息
	stack 	string					// 堆栈信息
	supplements []supplement		// 异常补充信息
	feedback string					// 用户反馈信息
}

// Feedback 获取可以用于反馈给用户对信息
func (slf *exception) Feedback() string {
	return app.feedbackFormat(slf.feedback)
}

// Supplement 对例外情况进行补充说明
func (slf *exception) Supplement(name string, content interface{}) Exception {
	supplement := newSupplement(name, content)
	slf.supplements = append(slf.supplements, supplement)
	if app.supplementMiddlewares != nil {
		for i := 0; i < len(app.supplementMiddlewares); i++ {
			app.supplementMiddlewares[i](slf, supplement)
		}
	}
	return slf
}

// Error 获取例外情况的错误信息
func (slf *exception) Error() string {
	return slf.packet()
}

// Print 打印例外情况的错误信息
func (slf *exception) Print() {
	fmt.Println(slf.packet())
}

// Code 获取例外情况的错误码
func (slf *exception) Code() int {
	return slf.code
}

// packet 错误信息封包
func (slf *exception) packet() string {
	var result = "Exception  >> " + slf.err.Error() + "\n"
	if len(slf.supplements) > 0 {
		result += "Supplement >>\n"
		for i := 0; i < len(slf.supplements); i++ {
			result += slf.supplements[i].stack + "\n"
		}
	}
	result += "Detailed   >>\n" + slf.stack + "\n\t..."
	return result
}