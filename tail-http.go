package main

import (
	"flag"
	"fmt"
	"github.com/hpcloud/tail"
	"net/http"
)

func FileMsg(filename string)(Text string){
	tails, err := tail.TailFile(filename, tail.Config{
		Follow:    false,
		Location: &tail.SeekInfo{Offset: -2000, Whence: 2},
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
	fmt.Println("请求:", r)
	fmt.Printf("路径:%v,文件内容:%v\n", r.URL.Path, Text)
}

func main() {
	var host,port string
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.StringVar(&port, "p", "28050", "端口号，默认28050")
	flag.Parse()
	address := host + ":" + port

	fmt.Println("服务已启动,", host, port)
	http.HandleFunc("/", say)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}
