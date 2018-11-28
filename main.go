package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"github.com/goinggo/mapstructure"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}

func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {

	// type panic struct{
	//  error string
	// }


	// fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {

		type Inventory struct {
			ENV             string
			Container_Name  string
			Namespace_Name  string
			Container_Image string
			Host            string
			Others          string
		}
		s := `ENV : ivkprod
container_name : conggateway-grpc
namespace_name : cong
container_image : ccr.ccs.tencentyun.com/dhub.wallstcn.com/conggateway:4d447f32
host : 10.1.129.86
prefix : echo
file : asm_amd64.s
line : 523
message : [PANIC RECOVER] runtime error: index out of range goroutine 4019297 [running]:
github.com/labstack/echo/middleware.RecoverWithConfig.func1.1.1(0x13f1f58, 0x1000, 0xc000780000, 0x1509040, 0xc0001ea2a0)
/tmp/govendor/src/github.com/labstack/echo/middleware/recover.go:75 +0x11d
panic(0x10c84a0, 0x2236d00)
/usr/local/go/src/runtime/panic.go:513 +0x1b9
gitlab.wallstcn.com/cong/conggateway/business/collector.(*RealProcess).GetInfo(0xc000e77308, 0x0, 0x0, 0xc002dfa040, 0x1, 0x1, 0xc0007e5770, 0x3, 0x3, 0x7f9f05ef9aa0, ...)
/go/src/gitlab.wallstcn.com/cong/conggateway/business/collector/real.go:145 +0x5e26
gitlab.wallstcn.com/cong/conggateway/service/api/collector.HttpRealAPI(0x1509040, 0xc0001ea2a0, 0x2268c40, 0xc000c480c4, 0x48, 0xc000e77418)
/go/src/gitlab.wallstcn.com/cong/conggateway/service/api/collector/market_data.go:44 +0x220
gitlab.wallstcn.com/cong/conggateway/cache.(*SingleRequest).handler(0xc000a73a50, 0xc000c480c4, 0x48, 0xc0003aa270, 0x13f2fa0, 0x1509040, 0xc0001ea2a0)
/go/src/gitlab.wallstcn.com/cong/conggateway/cache/single_request.go:46 +0xdf
gitlab.wallstcn.com/cong/conggateway/cache.(*SingleRequest).ProcessSingleRequest(0xc000a73a50, 0xc000c480c4, 0x48, 0x1509040, 0xc0001ea2a0, 0x13f2fa0, 0xc000a55d98, 0xc000a55d40, 0xb, 0x2288a80)
/go/src/gitlab.wallstcn.com/cong/conggateway/cache/single_request.go:38 +0x14e
gitlab.wallstcn.com/cong/conggateway/cache.WapperURLCacheResponse.func1(0x1509040, 0xc0001ea2a0, 0x48c936, 0x5bfcf65e)
/go/src/gitlab.wallstcn.com/cong/conggateway/cache/url_cache.go:27 +0xc2
github.com/labstack/echo.(*Echo).add.func1(0x1509040, 0xc0001ea2a0, 0x2268c40, 0x0)
/tmp/govendor/src/github.com/labstack/echo/echo.go:440 +0x8a
gitlab.wallstcn.com/wscnbackend/ivankaway/middleware.LogUserActivity.func1(0x1509040, 0xc0001ea2a0, 0x2268c40, 0xe4f2da3e9b4d288)
/tmp/govendor/src/gitlab.wallstcn.com/wscnbackend/ivankaway/middleware/logger.go:46 +0x1b2
gitlab.wallstcn.com/wscnbackend/ivankaway/middleware.APIMetrics.func2.1(0x1509040, 0xc0001ea2a0, 0xe, 0x0)
/tmp/govendor/src/gitlab.wallstcn.com/wscnbackend/ivankaway/middleware/metrics.go:28 +0xb7
gitlab.wallstcn.com/cong/conggateway/service/middleware.UserInfo.func1(0x1509040, 0xc0001ea2a0, 0xc0000571b0, 0xf)
/go/src/gitlab.wallstcn.com/cong/conggateway/service/middleware/middleware.go:72 +0xcb1
gitlab.wallstcn.com/wscnbackend/ivankaway/middleware.RequestRealIp.func1(0x1509040, 0xc0001ea2a0, 0xc00005739a, 0x6)
/tmp/govendor/src/gitlab.wallstcn.com/wscnbackend/ivankaway/middleware/requestip.go:29 +0x205
gitlab.wallstcn.com/wscnbackend/ivankaway/middleware.ProcessGetBindData.func1(0x1509040, 0xc0001ea2a0, 0x20, 0x132465b)
/tmp/govendor/src/gitlab.wallstcn.com/wscnbackend/ivankaway/middleware/middleware.go:44 +0x3bf
github.com/labstack/echo/middleware.CORSWithConfig.func1.1(0x1509040, 0xc0001ea2a0, 0xc001378ae0, 0x14f4e20)
/tmp/govendor/src/github.com/labstack/echo/middleware/cors.go:113 +0x2ce
github.com/labstack/echo/middleware.GzipWithConfig.func1.1(0x1509040, 0xc0001ea2a0, 0x0, 0x0)
/tmp/govendor/src/github.com/labstack/echo/middleware/compress.go:92 +0x15c
github.com/labstack/echo/middleware.RecoverWithConfig.func1.1(0x1509040, 0xc0001ea2a0, 0x0, 0x0)
/tmp/govendor/src/github.com/labstack/echo/middleware/recover.go:82 +0xd8
github.com/labstack/echo.(*Echo).ServeHTTP.func1(0x1509040, 0xc0001ea2a0, 0xc000850e58, 0x130a920)
/tmp/govendor/src/github.com/labstack/echo/echo.go:531 +0x108
github.com/labstack/echo.(*Echo).ServeHTTP(0xc000850e00, 0x14f4e20, 0xc0023440e0, 0xc000236200)
/tmp/govendor/src/github.com/labstack/echo/echo.go:540 +0x231
net/http.serverHandler.ServeHTTP(0xc000a61520, 0x14f4e20, 0xc0023440e0, 0xc000236200)
/usr/local/go/src/net/http/server.go:2741 +0xab
net/http.(*conn).serve(0xc000d426e0, 0x14f6120, 0xc00266e700)
/usr/local/go/src/net/http/server.go:1847 +0x646
created by net/http.(*Server).Serve
/usr/local/go/src/net/http/server.go:2851 +0x2f5

goroutine 1 [IO wait]:
internal/poll.runtime_pollWait(0x7f9f05ef1a20, 0x72, 0x0)
/usr/local/go/src/runtime/netpoll.go:173 +0x66
internal/poll.(*pollDesc).wait(0xc00005ac18, 0x72, 0xc000c86500, 0x0, 0x0)
/usr/local/go/src/i

time : 2018-11-27T07:46:38.674148922Z
level : -`
		var logMap map[string]string
		logMap = make(map[string]string)
		ss := strings.Split(s, "\n")
		for i := 0; i < 5; i++ {
			//fmt.Println("---------------")
			sss := strings.Split(ss[i], " : ")
			logMap[sss[0]] = sss[1]
		}
		var Others string
		for i := 5; i < len(ss); i++ {
			Others = Others + ss[i] + "\n"
		}
		logMap["Others"] = Others
		var inventory Inventory
		mapstructure.Decode(logMap, &inventory)
		tmpl, err := template.ParseFiles("email.html")
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, inventory)
		if err != nil {
			panic(err)
		}
	}
}
func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}