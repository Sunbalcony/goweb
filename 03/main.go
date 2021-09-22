package main

import (
	"core/framework"
	"fmt"
	"net/http"
)

//所有的框架代码都存放在 framework 文件夹中，而所有的示例业务代码都存放在 framework 文件夹之外
func main() {
	core := framework.NewCore()
	registerRouter(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
