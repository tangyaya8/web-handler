# WEB-HANDLER
是一个基于go语言实现的简单路由框架
- style: go_module
- env :go1.12.3 darwin/amd64
- IDE :goland

## 使用方法：
在业务代码中使用route.New()新建路由对象
然后编写HTTP METHOD 逻辑代码：
注意：restful中的参数指定为```>你的参数名称```
```go
func Test_Dispatcher(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}
	webHandler := New()
	webHandler.GET("/hello/>name/>id", hello)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func hello(writer http.ResponseWriter, request *http.Request, param *[]Param) {
	fmt.Println("我是处理业务逻辑的")
	fmt.Printf("参数%v", param)

}
```

## 运行机制：
1. 在WebHandler初始化的时候拦截所有```/``路径下的请求
2. 初始化所有HTTP方法,为每个Path新建一个hash表,用来存储URI和路径以及参数之间对应关系
3. 用户请求到达经过ServerHTTP拦截器(WebHandler核心实现)，对应Path和URI之间的关系，并封装相关请求参数
4. 执行目标方法
## 特性：
支持restful风格的web请求

## 不足：
不能像SpringMVC那样做到方法入口参数绑定
应该还有BUG没发现。。

