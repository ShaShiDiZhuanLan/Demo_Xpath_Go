/**
Author： 沙振宇
CreateTime：2019-12-4
UpdateTime：2019-12-4
Info：	“利用 Xpath 读取 html 内容”
		“通过Github来演示Xpath如何使用”
 */
package main

import (
	"fmt"
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xpath"
	"net/http"
	"strings"
)

// 去除空格和换行符
func getReplace(str string)  string{
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}

func main() {
	urlPath := "https://github.com/"
	res, err := http.Get(urlPath)
	if err != nil {
		panic("failed to get : " + err.Error())
	}

	doc, err := libxml2.ParseHTMLReader(res.Body)
	if err != nil {
		panic("failed to parse HTML: " + err.Error())
	}
	defer doc.Free()

	nodes := xpath.NodeList(doc.Find(`//summary[@class="HeaderMenu-summary HeaderMenu-link px-0 py-3 border-0 no-wrap d-block d-lg-inline-block"]/text()`))

	fmt.Printf("nodes type: %T,len: %d\n\n", nodes, len(nodes))

	for i := 0; i < len(nodes); i++ {
		fmt.Printf("nodes type: %T,text: %s\n", nodes[i], getReplace(nodes[i].String()))
	}
}