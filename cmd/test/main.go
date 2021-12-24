package main
import (
	"github.com/ddliu/go-httpclient"
	"time"
)

// httpclient.WithHeader 共享全局变量,并发写map 出现fatal error: concurrent map writes,且无法recover。
func main() {
		for i := 0; i <3;i++ {
			go do()
		}

	time.Sleep(1*time.Second)
}

func do()  {
	httpclient.
		WithHeader("1","1").
		WithHeader("2","2").
		WithHeader("3","3").
		WithHeader("4","4").Get("baidu.com")
}