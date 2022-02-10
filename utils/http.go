package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// APIError 微信支付 API v3 标准错误结构
type APIError struct {
	StatusCode int         // 应答报文的 HTTP 状态码
	Header     http.Header // 应答报文的 Header 信息
	Body       string      // 应答报文的 Body 原文
	Code       string      `json:"code"`             // 应答报文的 Body 解析后的错误码信息，仅不符合预期/发生系统错误时存在
	Message    string      `json:"message"`          // 应答报文的 Body 解析后的文字说明信息，仅不符合预期/发生系统错误时存在
	Detail     interface{} `json:"detail,omitempty"` // 应答报文的 Body 解析后的详细信息，仅不符合预期/发生系统错误时存在
}

// Error 输出 APIError
func (e *APIError) Error() string {
	var buf bytes.Buffer
	_, _ = fmt.Fprintf(&buf, "error http response:[StatusCode: %d Code: \"%s\"", e.StatusCode, e.Code)
	if e.Message != "" {
		_, _ = fmt.Fprintf(&buf, "\nMessage: %s", e.Message)
	}
	if e.Detail != nil {
		var detailBuf bytes.Buffer
		enc := json.NewEncoder(&detailBuf)
		enc.SetIndent("", "  ")
		if err := enc.Encode(e.Detail); err == nil {
			_, _ = fmt.Fprint(&buf, "\nDetail:")
			_, _ = fmt.Fprintf(&buf, "\n%s", strings.TrimSpace(detailBuf.String()))
		}
	}
	if len(e.Header) > 0 {
		_, _ = fmt.Fprint(&buf, "\nHeader:")
		for key, value := range e.Header {
			_, _ = fmt.Fprintf(&buf, "\n - %v=%v", key, value)
		}
	}
	_, _ = fmt.Fprintf(&buf, "]")
	return buf.String()
}

// APIResult 微信支付API v3 请求结果
type APIResult struct {
	Request  *http.Request  // 本次请求所使用的 HTTPRequest
	Response *http.Response // 本次请求所获得的 HTTPResponse
}

func Request(ctx context.Context, method string, url string, header http.Header, reqBody io.Reader, signBody interface{}) (*http.Response, error) {
	var (
		err     error
		request *http.Request
	)

	// Construct Request
	if request, err = http.NewRequestWithContext(ctx, method, url, reqBody); err != nil {
		return nil, err
	}

	// 增加请求头
	for key, values := range header {
		for _, v := range values {
			request.Header.Add(key, v)
		}
	}

	http := &http.Client{}

	result := &APIResult{
		Request: request,
	}

	result.Response, err = http.Do(request)

	if err != nil {
		return result.Response, err
	}
	// Check if Success
	if err = CheckResponse(result.Response); err != nil {
		return result.Response, err
	}
	return result.Response, nil
}

// CheckResponse 校验请求是否成功
//
// 当http回包的状态码的范围不是200-299之间的时候，会返回相应的错误信息，主要包括http状态码、回包错误码、回包错误信息提示
func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}
	slurp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("invalid response, read body error: %w", err)
	}
	_ = resp.Body.Close()

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(slurp))
	apiError := &APIError{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(slurp),
	}
	// 忽略 JSON 解析错误，均返回 apiError
	_ = json.Unmarshal(slurp, apiError)
	return apiError
}
