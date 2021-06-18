# 例外、错误信息封装（临时想法实现）

- 包含堆栈信息的错误信息；
- 支持对异常信息进行补充说明，得到更详细的过程中信息；

# 使用
> 建议任何异常都在事先进行声明，同时在使用过程中都应该对声明过的异常进行命中（Hit）操作。

> go get github.com/kercylan98/exception

### 定义例外信息
```
var (
    TestException = exception.Reg(1000, "exception information for testing", "用于测试的异常信息")
)
```

### 命中
```
TestException.Hit()
```

### 获取错误码
```
TestException.Hit().Code()
```

### 获取和打印错误详情
```
TestException.Hit().Error()
TestException.Hit().Print()
```
```
Exception  >> exception information for testing
Detailed   >>
	at /Users/kercylan/Coding.localized/Golang/exception/exception_test.go:26 (Method exception.TestException_Print)
	at /Users/kercylan/go/go1.16/src/testing/testing.go:1194 (Method testing.tRunner)
	at /Users/kercylan/go/go1.16/src/runtime/asm_arm64.s:1130 (Method runtime.goexit)
	...
```

### 对错误进行补充
```
TestException.Supplement("状态器ID", "8100393841").Print()
```
```
Exception  >> exception information for testing
Supplement >>
	状态器ID (8100393841) at /Users/kercylan/Coding.localized/Golang/exception/exception_test.go:31 (Method exception.TestException_Supplement)
Detailed   >>
	at /Users/kercylan/Coding.localized/Golang/exception/exception_test.go:30 (Method exception.TestException_Supplement)
	at /Users/kercylan/go/go1.16/src/testing/testing.go:1194 (Method testing.tRunner)
	at /Users/kercylan/go/go1.16/src/runtime/asm_arm64.s:1130 (Method runtime.goexit)
	...
```

### 获取用于对用户反馈的信息
```
TestException.Hit().Feedback()
```
> 当需要对该信息进行格式化时，可以事先使用exception.BindFeedbackFormatter函数进行绑定处理函数
> ```
> exception.BindFeedbackFormatter(func(feedback string) string {
>     return fmt.Sprintf("%v %s", time.Now(), feedback)
> }
> ```

### 在错误被命中时进行回调
```
exception.UseHitMiddleware(func(exception Exception) {
    fmt.Print("hit exception!", exception)
})
```

### 在对错误进行补充时进行回调
```
exception.UseSupplementMiddleware(func(exception Exception, supplement supplement) {
    fmt.Println(supplement.name, supplement.value)
    fmt.Println(supplement.stack)
    fmt.Println(exception)
})
TestException.Hit().Supplement("name", "value")
```
