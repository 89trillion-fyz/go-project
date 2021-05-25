package test

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	router2 "go-project/demo2/router"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

// Get 根据特定请求uri，发起get请求返回响应
func Get(uri string, router *gin.Engine) []byte {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// Post 根据特定请求uri，发起post请求返回响应
func Post(uri string, bodyJson string, router *gin.Engine) []byte {
	// 构造post请求
	req := httptest.NewRequest("POST", uri, bytes.NewBufferString(bodyJson))
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}
func TestGetResult(t *testing.T) {
	Router := router2.Routers()
	result := Post("/getResult", "{\"str\":\"1&1\"}", Router)
	fmt.Println("result =", string(result))
}
