package exception

import (
	"fmt"
	"testing"
	"time"
)

var (
	TestException = Reg(1000, "exception information for testing", "用于测试的异常信息")
)

func TestException_Code(t *testing.T) {
	t.Log(TestException.Hit().Code())
}

func TestException_Error(t *testing.T) {
	t.Log(TestException.Hit().Error())
}

func TestException_Feedback(t *testing.T) {
	t.Log(TestException.Hit().Feedback())
}

func TestException_Print(t *testing.T) {
	TestException.Hit().Print()
}

func TestException_Supplement(t *testing.T) {
	exp := TestException.Hit()
	exp.Supplement("supplement", "补充内容").Print()
}

func TestBindFeedbackFormatter(t *testing.T) {
	BindFeedbackFormatter(func(feedback string) string {
		return fmt.Sprintf("%v %s", time.Now(), feedback)
	})

	t.Log(TestException.Hit().Feedback())
}

func TestUseHitMiddleware(t *testing.T) {
	UseHitMiddleware(func(exception Exception) {
		t.Log("hit exception!", exception)
	})
	TestException.Hit()
}

func TestUseSupplementMiddleware(t *testing.T) {
	UseSupplementMiddleware(func(exception Exception, supplement supplement) {
		t.Log(supplement.name, supplement.value)
		t.Log(supplement.stack)
		t.Log(exception)
	})
	TestException.Hit().Supplement("name", "value")
}
