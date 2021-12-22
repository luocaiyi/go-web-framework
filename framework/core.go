package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	// 二级路由
	router map[string]*Tree
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	// 初始化路由
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{router: router}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// Group 初始化Group
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouterByRequest(request *http.Request) ControllerHandler {
	// uri 和 method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

// ServeHttp 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 封装自定义 context
	ctx := NewContext(request, response)

	// 寻找路由
	router := c.FindRouterByRequest(request)
	if router == nil {
		ctx.Json(http.StatusNotFound, "not found")
		return
	}

	// 调用路由函数，如果返回 err 代表存在内部错误，返回500状态码
	if err := router(ctx); err != nil {
		ctx.Json(http.StatusInternalServerError, "inner error")
		return
	}
}
