package route

import (
	"log"
	"net/http"
	"strings"
	. "webHandler/config"
	. "webHandler/handler"
)

type WebHandler struct {
	URL             string
	MethodHandler   http.Handler
	NotFountHandler http.Handler
}

func New() *WebHandler {
	webHandler := new(WebHandler)
	http.Handle("/", webHandler)
	return webHandler
}

var Mappings map[string]map[string]Node

type Node struct {
	Params  []Param
	Method  string
	handler Req
}

func init() {
	Mappings = map[string]map[string]Node{
		"POST":    nil,
		"PUT":     nil,
		"HEAD":    nil,
		"DELETE":  nil,
		"OPTIONS": nil,
	}
}

type Req func(writer http.ResponseWriter, request *http.Request, param *[]Param)

func (w *WebHandler) GET(path string, req Req) {
	//初始化方法
	w.initHandler("GET", path, req)
}

func (w *WebHandler) POST(path string, req Req) {
	//初始化方法
	w.initHandler("POST", path, req)

}
func (w *WebHandler) DELETE(path string, req Req) {
	//初始化方法
	w.initHandler("DELETE", path, req)

}

func (w *WebHandler) HEAD(path string, req Req) {
	//初始化方法
	w.initHandler("DELETE", path, req)

}
func (w *WebHandler) OPTIONS(path string, req Req) {
	//初始化方法
	w.initHandler("OPTIONS", path, req)

}
func (w *WebHandler) initHandler(method, path string, req Req) {
	s, param, err := ExtractURL(DealWithUrl(path))
	if err != nil {
		log.Println("提取参数错误")
	}
	if Mappings[method] == nil {
		Mappings[method] = make(map[string]Node)
	}
	if mapping, ok := Mappings[method][s]; ok {
		//有多个URL
		log.Fatalf("duplicate Mapping %#v", mapping)
	} else {
		Mappings[method][s] = Node{
			Params:  param,
			Method:  method,
			handler: req,
		}
	}
}

//核心调度器
func (w *WebHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	requestURI := request.RequestURI
	// /hello/1/  /hello/tangbaobao/2
	// /hello/
	method := request.Method
	for _, mapping := range Mappings {
		for url, node := range mapping {
			if strings.Contains(requestURI, url) {
				requestURI = strings.TrimLeft(requestURI, url)
				paramKey := strings.Split(strings.TrimRight(requestURI, UnixSeparator), UnixSeparator)
				if len(paramKey) == len(node.Params) {
					//URL对应成功
					//检查method是否对应：
					if method != node.Method {
						log.Print("Method is not allow")
						if w.MethodHandler != nil {
							w.MethodHandler.ServeHTTP(writer, request)
						} else {
							http.Error(writer,
								http.StatusText(http.StatusMethodNotAllowed),
								http.StatusMethodNotAllowed,
							)
						}
						return
					}
					//封装参数
					for index, value := range paramKey {
						node.Params[index].Value = value
					}
					node.handler(writer, request, &node.Params)
					return
				}
			}
		}
	}
	//找不到mapping
	if w.NotFountHandler != nil {
		w.NotFountHandler.ServeHTTP(writer, request)
	} else {
		http.Error(writer,
			http.StatusText(http.StatusNotFound),
			http.StatusNotFound,
		)
	}
}
