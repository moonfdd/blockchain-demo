package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	d, err := ViewUrl("http://c.gb688.cn/bzgk/gb/showGb?type=online&hcno=CEF016D5EB296398BE3AB89CAE46B5FA")
	if err != nil {
		fmt.Println("失败", err)
		return
	}
	fmt.Println("成功", string(d))
	os.WriteFile("gb.html", d, 0666)
}

func ViewUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "访问失败！", err)
		return nil, err
	}
	defer resp.Body.Close()
	var d []byte
	d, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取失败")
		return nil, err
	}

	return d, nil
}
