package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	url := "https://movie.douban.com/review/15217513/"

	// 创建一个HTTP客户端，用于发送HTTP请求
	client := &http.Client{}

	// 创建一个GET请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 添加自定义的请求头
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	//req.Header.Add("Authorization", "Bearer YourAccessToken")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP request failed with status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 解析HTML
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
	}
	// 打印HTML内容
	htmlString := renderHTML(doc)
	fmt.Println(htmlString)
	// 定位评论数据的HTML元素并提取数据
	// 这部分需要根据实际页面结构来完成

	// 处理分页，如果需要的话

	// 存储数据到你的数据存储中
}

func renderHTML(n *html.Node) string {
	var buf bytes.Buffer
	renderNode(&buf, n)
	return buf.String()
}

func renderNode(w *bytes.Buffer, n *html.Node) {
	if n == nil {
		return
	}

	switch n.Type {
	case html.ElementNode:
		if n.DataAtom == atom.Script || n.DataAtom == atom.Style {
			// 跳过脚本和样式标签
			return
		}
		w.WriteString("<" + n.Data)
		for _, attr := range n.Attr {
			w.WriteString(fmt.Sprintf(` %s="%s"`, attr.Key, attr.Val))
		}
		if n.FirstChild == nil {
			w.WriteString("/>")
		} else {
			w.WriteString(">")
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				renderNode(w, c)
			}
			w.WriteString("</" + n.Data + ">")
		}

	case html.TextNode:
		w.WriteString(n.Data)

	case html.CommentNode:
		w.WriteString("<!--" + n.Data + "-->")

	case html.DoctypeNode:
		w.WriteString("<!DOCTYPE " + n.Data + ">")

	default:
		// 处理其他节点类型
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			renderNode(w, c)
		}
	}
}
