package framework

import "net/http"

// Core 框架核心结构
type Core struct {
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	return &Core{}
}

// ServeHttp 框架核心结构实现Handler接口
func (c *Core) ServeHttp(response http.Response, request *http.Request) {
	// TODO
}
