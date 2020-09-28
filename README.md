# tail-http
> 通过http请求查看服务器上的日志

## 服务说明
1. 下载tail-http到服务器（目前不支持Windows）
2. 给予运行权限，`chmod +x tail-http`
3. 运行。
> 运行参数:  
-h 指定监听地址，缺省值:localhost  
-p 指定监听端口,缺省值:28050  
-help 输出帮助页后停止程序  
-queit 设置静默模式，不会输出请求和返回信息  

## 使用说明
通过http请求此服务时，服务会把请求信息的uri部分用作服务器内的文件路径。  
也就是如果我想查看/opt/nginx/logs/access.log日志的时候，应该访问地址:`http://localhost:28050/opt/nginx/logs/access.log`  
你会收到此日志文件的后2000字节的内容  
> 如果此文件小于8000字节，服务将不会输出内容

## 编译说明
下载go环境，  
运行:`go build tail-http.go`  
