package main

import (
	"flag"
	"fmt"
	"github.com/hpcloud/tail"
	"net/http"
)

var (
	host,port string
	help,quiet bool
)

func FileMsg(filename string)(Text string){
	tails, err := tail.TailFile(filename, tail.Config{
		Follow:    false,
		Location: &tail.SeekInfo{Offset: -10000, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	for msg := range tails.Lines{
		Text = Text + msg.Text + "\n"
	}
	return Text
}

func say(w http.ResponseWriter, r *http.Request){
	Text := FileMsg(r.URL.Path)
	w.Write([]byte(Text))
	if ! quiet{
		fmt.Println("请求:", r)
		fmt.Printf("路径:%v,内容:%v\n", r.URL.Path, Text)
	}
}

func web() {
	fmt.Println("服务已启动,", host, port)
	http.HandleFunc("/", say)
	address := host + ":" + port
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}

func main() {
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.StringVar(&port, "p", "28050", "端口号，默认28050")
	flag.BoolVar(&help, "help", false, "帮助页")
	flag.BoolVar(&quiet, "quiet", false, "不输出控制台信息")
	flag.Parse()

	if help {
		fmt.Printf(
			"tail-http v0.2\thttps://github.com/LIU-LIU-LIU/tail-http\n" +
			"帮助页:\n" +
			"运行参数:\n" +
			"-h 指定监听地址，缺省值:localhost\n" +
			"-p 指定监听端口,缺省值:28050\n" +
			"-help 输出帮助页后停止程序\n" +
			"-queit 设置静默模式，不会输出请求和返回信息\n" +
			"使用说明:\n" +
			"通过http请求此服务时，服务会把请求信息的uri部分用作服务器内的文件路径。\n" +
			"也就是如果我想查看/opt/nginx/logs/access.log日志的时候，应该访问地址:http://localhost:28050/opt/nginx/logs/access.log\n" +
			"你会收到此日志文件的后10,000字节的内容\n" +
			"如果此文件小于10,000字节，服务将不会输出内容\n")
		return
	}
	web()
}
