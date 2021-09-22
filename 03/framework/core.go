package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler

}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("core serveHTTP")
	ctx := NewContext(r, w)

	//一个简单的路由选择器，这里写死为测试路由foo

	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")
	router(ctx)

}
