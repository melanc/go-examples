package main
 
import (
    "fmt"
    "net"
    "net/rpc"
	"net/rpc/jsonrpc"
    "github.com/melanc/go-practice/rpc/common"
)
 
func main() {
    srv := rpc.NewServer()
	// 注册方式：srv.RegisterName(Namespace, Pointer)
	srv.RegisterName("common.MathServ", new(common.MathService))
	
    fmt.Println("启动服务...")
    var address, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:1234") //定义TCP的服务承载地址
    listener, err := net.ListenTCP("tcp", address)               //监听TCP连接
    if err != nil {
        fmt.Println("启动失败！", err)
    }
    for {
        conn, err := listener.Accept() //如果接受到连接
        if err != nil {
            continue
        }
        fmt.Println("接收到一个调用请求...")
        srv.ServeCodec(jsonrpc.NewServerCodec(conn)) //让此rpc绑定到该Tcp连接上。
    }
    //fmt.Println("服务已停止!")
}