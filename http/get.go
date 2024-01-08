package main

/*import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Printf("TCP连接错误，err: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Print(string(body))
}

func main() {
	apiUrl := "http://127.0.0.1:9090/get"
	data := url.Values{}
	data.Set("name", "zhangsan")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl) //解析相对或绝对的URI，返回一个url.URL类型的对象
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Print(string(b))
}
*/
