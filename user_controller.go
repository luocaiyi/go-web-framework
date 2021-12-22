package main

import (
	"github.com/luocaiyi/go-web-framework/framework"
	"net/http"
)

func UserLoginController(ctx *framework.Context) error {
	ctx.Json(http.StatusOK, "ok, UserLoginController")
	return nil
}
