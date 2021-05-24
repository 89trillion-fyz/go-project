package test

import (
	"demo1/global"
	"demo1/router"
	response "demo1/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func init() {
	global.G_JSONPATH = "../config.army.model.json"
}

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
func TestArmyFindByRarityAndLock(t *testing.T) {
	//api.FindByRarityAndLock()
	Router := router.Routers()
	body := Get("/army/findByRarityAndLock?rarity=1&lock=0", Router)
	fmt.Printf("response:%v\n", string(body))
	// 解析响应，判断响应是否与预期一致
	response := &response.Response{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
}
func TestArmyFindRarityById(t *testing.T) {
	//api.FindByRarityAndLock()
	Router := router.Routers()
	body := Get("/army/findRarityById?id=10101", Router)
	fmt.Printf("response:%v\n", string(body))
	response := &response.Response{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
}
func TestFindQualityById(t *testing.T) {
	Router := router.Routers()
	body := Get("/army/findQualityById?id=10101", Router)
	fmt.Printf("response:%v\n", string(body))
	response := &response.Response{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
}
