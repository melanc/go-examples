package main
 
import (
    "fmt"
	"net/rpc/jsonrpc"
    "github.com/melanc/go-practice/rpc/common"
)
 
func main() {
    // client和server的rpc类型保持一致
    // 此处都使用jsonrpc
    var client, err = jsonrpc.Dial("tcp", "127.0.0.1:1234")
    if err != nil {
        fmt.Println("连接不到服务器：", err)
    }
    var args = common.Args{40, 3}
    var result = common.Result{}
    fmt.Println("开始调用！")
	
	// 调用方式：Namepsace.Method()
	// Method前面不需要Struct名称引用
    err = client.Call("common.MathServ.Add", args, &result)
    if err != nil {
        fmt.Println("调用失败！", err)
    }
    fmt.Println("调用成功！结果：", result.Value)
}