package main

import "core/framework"

func registerRouter(core *framework.Core)  {
	core.Get("foo",FooControllerHandler)

}
