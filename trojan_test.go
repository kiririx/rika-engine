package main

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"
)

// TestProxy 测试代理是否正常, 通过请求https://www.google.com
func TestProxy(t *testing.T) {
	proxyHost := "127.0.0.1"
	proxyPort := "17890"

	proxyURL, err := url.Parse(fmt.Sprintf("http://%s:%s", proxyHost, proxyPort))
	if err != nil {
		t.Fatalf("解析代理 URL 失败: %v", err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("https://www.google.com")
	if err != nil {
		t.Fatalf("通过代理请求 Google 失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("预期状态码 200, 但得到 %d", resp.StatusCode)
	}

	t.Log("代理测试成功")
}
