package req

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sylphbyte/pr"
	// 如有 pr 包可用，取消注释下一行
	// "github.com/sylphbyte/pr"
)

const (
	GetMethod  MethodType = "GET"
	PostMethod MethodType = "POST"

	FormType ContentType = "application/x-www-form-urlencoded"
	JsonType ContentType = "application/json"
)

type Response struct {
	Code int
	Body []byte
}

func takeResponse(resp *http.Response) (ret *Response, err error) {
	var bs []byte
	if bs, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	ret = &Response{
		Code: resp.StatusCode,
		Body: bs,
	}

	return
}

type ContentType string

type MethodType string

type Request struct {
	Method      MethodType
	ContentType ContentType
	Url         string
	Params      map[string]interface{}
	Header      http.Header
	Timeout     time.Duration
}

func NewRequest(method MethodType, contentType ContentType, url string, params map[string]interface{}, header http.Header, timeout time.Duration) *Request {
	if method == GetMethod {
		url = getRequestURL(url, params)
		params = nil
	}

	return &Request{Method: method, ContentType: contentType, Url: url, Params: params, Header: header, Timeout: timeout}
}

// ShowRequestLog 控制是否显示详细请求日志
var ShowRequestLog = false

// SetShowRequestLog 设置请求日志显示开关
func SetShowRequestLog(show bool) {
	ShowRequestLog = show
}

func DoRequest(name any, req *Request) (*http.Response, error) {
	return do(name, req.Method, req.ContentType, req.Url, req.Params, req.Header, req.Timeout)
}

func Auto(name any, method MethodType, contentType ContentType, url string, params map[string]interface{}, header http.Header, duration time.Duration) (*http.Response, error) {
	if method == GetMethod {
		return Get(url, params)
	}

	if contentType == FormType {
		return Form(name, url, params, header, duration)
	}

	return Json(name, url, params, header, duration)
}

func Get(url string, params map[string]interface{}) (*http.Response, error) {
	client := http.Client{}
	return client.Get(getRequestURL(url, params))
}

// getRequestURL 获取Get 请求
func getRequestURL(url string, params map[string]interface{}) string {
	queryString := queryParams(params, "")
	return fmt.Sprintf("%s?%s", url, queryString)
}

func Form(name any, url string, params map[string]interface{}, header http.Header, duration time.Duration) (*http.Response, error) {
	if header == nil {
		header = http.Header{}
	}
	header.Set("Content-Type", string(FormType))
	return do(name, PostMethod, FormType, url, params, header, duration)
}

func Json(name any, url string, params map[string]interface{}, header http.Header, duration time.Duration) (*http.Response, error) {
	if header == nil {
		header = http.Header{}
	}

	header.Set("Content-Type", "application/json;charset=utf-8")
	return do(name, PostMethod, JsonType, url, params, header, duration)
}

func do(name any, method MethodType, contentType ContentType, url string, params map[string]interface{}, header http.Header, duration time.Duration) (resp *http.Response, err error) {
	if ShowRequestLog {
		fmt.Println()
		pr.Cyan(">>>remote: request info %s\n", strings.Repeat("==", 50))
		pr.Cyan(">>>remote: name: %v\n", name)
		pr.Cyan(">>>remote: method: %s\n", method)
		pr.Cyan(">>>remote: url: %s\n", url)
		pr.Cyan(">>>remote: params: %+v\n", params)
		pr.Cyan(">>>remote: header: %+v\n", header)
	}
	timeStart := time.Now()

	req, err := makeRequest(method, contentType, url, params)
	if err != nil {
		return
	}

	req.Header = header
	client := http.Client{
		Timeout: duration,
	}

	resp, err = client.Do(req)

	useTime := time.Since(timeStart).Milliseconds()

	if ShowRequestLog {
		var code int
		var bodyString string
		if resp != nil {
			code = resp.StatusCode
			bs, _ := io.ReadAll(resp.Body)
			bodyString = string(bs)
			resp.Body = io.NopCloser(bytes.NewReader(bs))
		}
		pr.Cyan(">>>remote: got code: %d\n", code)
		pr.Cyan(">>>remote: got body: %s\n", bodyString)
		if useTime > 1000 {
			pr.Red(">>>remote: use time: %dms\n", useTime)
		} else {
			pr.Green(">>>remote: use time: %dms\n", useTime)
		}
		pr.Cyan(">>>remote: ended %s\n", strings.Repeat("==", 50))
		fmt.Println()
	}

	return
}

func makeRequest(method MethodType, typ ContentType, url string, params map[string]interface{}) (*http.Request, error) {
	return http.NewRequest(string(method), url, getData(typ, params))
}

func getData(typ ContentType, params map[string]interface{}) io.Reader {
	if typ == JsonType {
		js, _ := _json.Marshal(params)
		return bytes.NewReader(js)
	}

	return strings.NewReader(queryParams(params, ""))
}

func queryParams(params map[string]interface{}, format string) string {
	values := url.Values{}
	var nk, ret string
	for k, v := range params {
		if len(format) != 0 {
			nk = fmt.Sprintf(format, k)
		} else {
			nk = k
		}

		switch v.(type) {
		case string:
			values.Add(nk, v.(string))
			break
		case []byte:
			values.Add(nk, string(v.([]byte)))
			break
		case map[string]interface{}:
			ret += queryParams(v.(map[string]interface{}), nk+"[%s]")
			ret += "&"
		case int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
			values.Add(nk, fmt.Sprintf("%d", v))
		}
	}

	ret += values.Encode()
	return ret
}

func PostForm(name any, url string, data map[string]interface{}, header http.Header, requestTimeout time.Duration) (ret *Response, err error) {
	var resp *http.Response
	if resp, err = Form(name, url, data, header, requestTimeout); err != nil {
		return
	}

	defer resp.Body.Close()
	return takeResponse(resp)
}

func PostJson(name any, url string, params map[string]interface{}, header http.Header, duration time.Duration) (ret *Response, err error) {
	var resp *http.Response
	if resp, err = Json(name, url, params, header, duration); err != nil {
		return
	}

	defer resp.Body.Close()
	return takeResponse(resp)
}

func FastGet(url string, data map[string]interface{}) (ret *Response, err error) {
	var resp *http.Response
	if resp, err = Get(url, data); err != nil {
		return
	}

	defer resp.Body.Close()
	return takeResponse(resp)
}
