package main

import "github.com/luocaiyi/go-web-framework/framework"

func registerRouter(core *framework.Core) {
	// 需求 1+2 : HTTP方法 + 静态路由匹配
	core.Get("/user/login", UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Post("", SubjectAddController)
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}
