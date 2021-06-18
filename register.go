package exception

import (
	"log"
	"sync"
)

var app = &register{
	mutex: new(sync.Mutex),
	exceptions: map[int]Factory{},
	feedbackFormat: func(feedback string) string {
		return feedback
	},
}

// FeedbackFormatter 用于对用户反馈信息的格式化工具（可用于全球化等处理）
type FeedbackFormatter func(feedback string) string

// HitMiddleware 例外命中中间件，当任何例外信息被命中时都会传入该中间件
//
// 需要注意的是，例外信息是可以进行补充说明的，也就是该中间件无法捕捉到后续的补充说明
type HitMiddleware func(exception Exception)

// SupplementMiddleware 例外信息补充中间件，当任何例外信息被补充说明时都将传入该中间件
//
// 一个例外信息可能被多次传入
type SupplementMiddleware func(exception Exception, supplement supplement)

// register 错误信息注册机
//
// 维持单例模式运行，应当在任何exception包函数调用前进行初始化
type register struct {
	mutex 			*sync.Mutex						// 避免在init函数中使用协程或其他并发情况而导致问题的互斥锁
	exceptions 		map[int]Factory					// 记录特定例外情况是否已注册，同时记录错误码和例外工厂的映射
	feedbackFormat	FeedbackFormatter				// 用户反馈信息的格式化工具
	hitMiddlewares 	[]HitMiddleware					// 命中中间件
	supplementMiddlewares []SupplementMiddleware 	// 补充中间件
}

// Reg 根据错误码及错误信息注册全局例外信息，如果code重复将会中止程序运行
//
// 在注册例外信息的时候，允许添加用户反馈信息，以实现对用户对友好提示
//
// 如果设置多个feedback，默认取第一个
func Reg(code int, text string, feedback ...string) Factory {
	app.mutex.Lock()
	defer app.mutex.Unlock()

	if _, exist := app.exceptions[code]; exist {
		log.Fatalf("registry exception failed. exception code \"%d\" existed", code)
	}
	factory := newFactory(code, text, feedback...)
	app.exceptions[code] = factory

	return factory
}

// BindFeedbackFormatter 绑定用户反馈信息格式化工具
func BindFeedbackFormatter(formatter FeedbackFormatter) {
	app.mutex.Lock()
	defer app.mutex.Unlock()
	app.feedbackFormat = formatter
}

// UseHitMiddleware 采用命中中间件
func UseHitMiddleware(middleware ...HitMiddleware) {
	app.mutex.Lock()
	defer app.mutex.Unlock()
	if app.hitMiddlewares == nil {
		app.hitMiddlewares = []HitMiddleware{}
	}
	app.hitMiddlewares = append(app.hitMiddlewares, middleware...)
}

// UseSupplementMiddleware 采用补充中间件
func UseSupplementMiddleware(middleware ...SupplementMiddleware) {
	app.mutex.Lock()
	defer app.mutex.Unlock()
	if app.supplementMiddlewares == nil {
		app.supplementMiddlewares = []SupplementMiddleware{}
	}
	app.supplementMiddlewares = append(app.supplementMiddlewares, middleware...)
}