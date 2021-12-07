package go_web_framework

import (
	"github.com/luocaiyi/go-web-framework/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr: ":8080",
	}
	server.ListenAndServe()
}
