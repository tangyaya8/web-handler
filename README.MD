# WEB-HANDLER
是一个基于go语言实现的简单路由框架
- style: go_module
- env :go1.12.3 darwin/amd64
- IDE :goland

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

