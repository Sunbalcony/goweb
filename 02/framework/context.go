package framework

import (
	"context"
	"net/http"
	"sync"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	//是否超时标记位
	hasTimeout bool
	//写保护机制
	writerMux *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
	}

}

//基础函数

// WriteMux 加锁保护
func (ctx *Context) WriteMux() *sync.Mutex {
	return ctx.writerMux

}

// GetRequest 获取请求
func (ctx *Context) GetRequest() *http.Request {
	return ctx.request

}

// GetResponse 获取响应
func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter

}


func (ctx *Context) SetHandler() {

}
