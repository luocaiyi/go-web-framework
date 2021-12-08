package main

import "github.com/luocaiyi/go-web-framework/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
