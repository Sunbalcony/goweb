package framework

import (
	"context"
	"net/http"
	"sync"
)

type Context struct {
	request *http.Request
	responseWriter http.ResponseWriter
	ctx context.Context
	//是否超时标记位
	hasTimeout bool
	//写保护机制
	writerMux *sync.Mutex

}

func NewContext(r *http.Request,w http.ResponseWriter) *Context {
	return &Context{
		request: r,
		responseWriter: w,
		ctx: r.Context(),
		writerMux: &sync.Mutex{},

	}

}

