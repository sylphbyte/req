package req

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/sylphbyte/pr"
)

func TestRequestLog(t *testing.T) {
	pr.Enable()
	// 测试数据
	getURL := "https://httpbin.org/get"
	postURL := "https://httpbin.org/post"
	params := map[string]interface{}{
		"name":    "测试用户",
		"id":      12345,
		"isValid": true,
		"nested": map[string]interface{}{
			"foo": "bar",
		},
	}
	header := http.Header{}
	header.Set("X-Test-Header", "测试头部")
	timeout := 10 * time.Second

	// 测试1: 打开日志开关，测试 GET 请求
	fmt.Println("\n===== 测试1: ShowRequestLog=true, GET 请求 =====")
	SetShowRequestLog(true)
	resp, err := Get(getURL, params)
	if err != nil {
		t.Fatalf("GET 请求失败: %v", err)
	}
	resp.Body.Close()

	// 测试2: 打开日志开关，测试 POST JSON 请求
	fmt.Println("\n===== 测试2: ShowRequestLog=true, POST JSON 请求 =====")
	resp, err = Json(postURL, params, header, timeout)
	if err != nil {
		t.Fatalf("POST JSON 请求失败: %v", err)
	}
	resp.Body.Close()

	// 测试3: 打开日志开关，测试 POST FORM 请求
	fmt.Println("\n===== 测试3: ShowRequestLog=true, POST FORM 请求 =====")
	resp, err = Form(postURL, params, header, timeout)
	if err != nil {
		t.Fatalf("POST FORM 请求失败: %v", err)
	}
	resp.Body.Close()

	// 测试4: 关闭日志开关，测试请求
	fmt.Println("\n===== 测试4: ShowRequestLog=false, GET 请求 =====")
	SetShowRequestLog(false)
	resp, err = Get(getURL, params)
	if err != nil {
		t.Fatalf("ShowRequestLog=false GET 请求失败: %v", err)
	}
	resp.Body.Close()

	fmt.Println("完成所有测试！如果上方有彩色日志输出，且最后一个测试没有请求日志，说明功能正常。")
}

func TestRequestWithResponse(t *testing.T) {
	fmt.Println("\n===== 测试 Response 类型的请求方法 =====")

	// 测试数据
	getURL := "https://httpbin.org/get"
	postURL := "https://httpbin.org/post"
	params := map[string]interface{}{
		"name": "测试 Response 对象",
		"time": time.Now().Unix(),
	}
	header := http.Header{}
	header.Set("X-Test-Source", "req-test")
	timeout := 10 * time.Second

	// 打开日志
	SetShowRequestLog(true)

	// 1. 测试 FastGet
	fmt.Println("\n----- 测试 FastGet -----")
	resp, err := FastGet(getURL, params)
	if err != nil {
		t.Fatalf("FastGet 失败: %v", err)
	}
	fmt.Printf("FastGet 响应码: %d, 响应体长度: %d\n", resp.Code, len(resp.Body))

	// 2. 测试 PostJson
	fmt.Println("\n----- 测试 PostJson -----")
	resp, err = PostJson(postURL, params, header, timeout)
	if err != nil {
		t.Fatalf("PostJson 失败: %v", err)
	}
	fmt.Printf("PostJson 响应码: %d, 响应体长度: %d\n", resp.Code, len(resp.Body))

	// 3. 测试 PostForm
	fmt.Println("\n----- 测试 PostForm -----")
	resp, err = PostForm(postURL, params, header, timeout)
	if err != nil {
		t.Fatalf("PostForm 失败: %v", err)
	}
	fmt.Printf("PostForm 响应码: %d, 响应体长度: %d\n", resp.Code, len(resp.Body))

	// 关闭日志
	SetShowRequestLog(false)
}
