package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	router map[string]map[string]ControllerHandler //二级map
}

func NewCore() *Core {
	//定义二级map
	//定义路由map
	//[URI]handler
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}
	//将二级map写入一级map
	//[method][URI]handler
	router := map[string]map[string]ControllerHandler{}
	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter
	return &Core{router: router}
}

//`````````````````路由注册start```````````````````````````````

// Get method
func (c *Core) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = handler

}

//Post method
func (c *Core) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler
}

// Put method
func (c *Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = handler
}

// Delete method
func (c *Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["DELETE"][upperUrl] = handler
}

//`````````````````路由注册end```````````````````````````````

//路由匹配,如果没有匹配到则返回nil

func (c *Core) FindRouteByRequest(r *http.Request) ControllerHandler {
	//uri和method全部转换为大写，保证大小写不敏感
	uri := r.URL.Path
	method := r.Method
	upperMethod := strings.ToUpper(method)
	upperUri := strings.ToUpper(uri)

	//匹配路由map
	//查找第一层
	if methodHandlers,ok:=c.router[upperMethod];ok{
		//查找第二层map
		if handler,ok:=methodHandlers[upperUri];ok{
			return handler
		}
	}
	return nil

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
