package main

import (
	"context"
	"core/framework"
	"fmt"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	//负责通知结束
	panicChan := make(chan interface{}, 1)
	//负责通知panic异常

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()
	go func() {

		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// 具体业务逻辑
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		//新的goroutine结束的时候通过一个finish通道告知父goroutine
		finish <- struct{}{}
	}()
	//请求监听时增加锁机制
	select {
	//监听panic
	case p := <-panicChan:
		//边界场景
		//异常事件、超时事件触发时，需要往 responseWriter 中写入信息，这个时候如果有其他 Goroutine 也要操作 responseWriter，会不会导致 responseWriter 中的信息出现乱序？
		//我们要保证在事件处理结束之前，不允许任何其他 Goroutine 操作 responseWriter，这里可以使用一个锁（sync.Mutex）对 responseWriter 进行写保护
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "Panic")
		//监听结束事件
	case <-finish:
		fmt.Println("finish")
		//监听超时事件
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "TimeOut")
		//边界场景
		//超时事件触发结束之后，已经往 responseWriter 中写入信息了，这个时候如果有其他 Goroutine 也要操作 responseWriter， 会不会导致 responseWriter 中的信息重复写入？
		//我们可以设计一个标记，当发生超时的时候，设置标记位为 true，在 Context 提供的 response 输出函数中，先读取标记位；当标记位为 true，表示已经有输出了，不需要再进行任何的 response 设置了
		c.SetHasTimeout()

	}
	return nil

}
