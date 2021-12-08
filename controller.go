package main

import (
	"context"
	"fmt"
	"github.com/luocaiyi/go-web-framework/framework"
	"net/http"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {
	// finish 负责通知结束
	finish := make(chan struct{}, 1)
	// panicChan 负责通知 panic 异常
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
	// 此处记得处理完毕后调用 cancel，告知 durationCtx 的后续 Context 结束
	defer cancel()

	go func() {
		// 增加异常处理
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// 这里做具体业务
		time.Sleep(10 * time.Second)
		ctx.Json(http.StatusOK, "ok")
		// 新的 goroutine 结束的时候通过一个 finish 通道告知父 goroutine
		finish <- struct{}{}
	}()

	select {
	// 监听 panic
	case <-panicChan:
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.Json(http.StatusOK, "panic")
	// 监听结束
	case <-finish:
		fmt.Println("finish")
	// 监听超时
	case <-durationCtx.Done():
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.Json(http.StatusInternalServerError, "time out")
		ctx.SetTimeout()
	}

	return nil
}
