package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handler        ControllerHandler
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

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux

}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request

}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHandler(handler ControllerHandler) {
	ctx.handler = handler
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

//

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()

}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()

}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()

}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)

}

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def

}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.URL.Query())
	}
	return map[string][]string{}

}
