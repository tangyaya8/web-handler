package route

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_Dispatcher(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}
	webHandler := New()
	webHandler.GET("/hello/>name/>id", handle)
	webHandler.GET("/hello/>name/>id", handle)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func handle(writer http.ResponseWriter, request *http.Request, param *[]Param) {
	fmt.Println("我是处理业务逻辑的")
	fmt.Printf("参数%v", param)

}
