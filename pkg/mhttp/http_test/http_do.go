package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{}

func init() {
	client.Timeout = time.Minute * 2
}

func main()  {

	http.HandleFunc("/", TestDo)
	log.Printf("%v", http.ListenAndServe("locahost:8080", nil))
}

func TestDo(w http.ResponseWriter, req *http.Request)  {
	time.Sleep(time.Second * 2)

	// 验证 csrf
	if req.Header.Get("_csrf") != "123456" {
		http.Error(w, "无效的 csrf token", 400)
	}

	// 获取cookie
	cookie := req.Header.Get("Cookie")
	if cookie == "" {
		http.Error(w, "请登录后再操作", 401)
	}

	// 获取用户名，简单的利用 strings 去截取字符串，只是个简单示例，没有考虑那么多可能性。
	fmt.Printf("%v", cookie)
	index :=  strings.Index(cookie, "=")
	name := cookie[index+1:]
	fmt.Printf("%v", name)
	if name != "Chenpy" {
		http.Error(w, "当前用户没有权限操作", 401)
		return
	}
	io.WriteString(w, "hello"+name)
}

func clientGo(){
	req, err := http.NewRequest("POST", "http://localhost:8000", nil)
	ErrPrint(err)

	// 设置一个 csrf token, 服务器端会去验证这个
	req.Header.Set("_csrf", "123456")
	// 设置一个 Cookie, 服务器端会去验证这个
	req.Header.Set("Cookie", "name=BroQiang")
	// 一会写完服务器端可以尝试分别将上面两行注释去测试

	resp, err := client.Do(req)
	ErrPrint(err)

	defer resp.Body.Close()
	DataPrint(resp.Body)
	// 打印下状态码，看下效果
	fmt.Printf("返回的状态码是： %v\n", resp.StatusCode)
	fmt.Printf("返回的信息是： %v\n", resp.StatusCode)
}



func DataPrint(body io.ReadCloser) {
	bytes, err := ioutil.ReadAll(body)
	ErrPrint(err)

	// 这里要格式化再输出， 因为ReadAll 返回的是字节切片
	fmt.Printf("%s\n", bytes)
}

func ErrPrint(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}


