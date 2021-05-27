package test

import (
	"bytes"
	"fmt"
	model "go-project/demo3/proto"
	"go-project/demo3/router"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/gin-gonic/gin"
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
func TestCreateCdkey(t *testing.T) {
	Router := router.Routers()
	result := Post("/cdkey/createCdkey", "{\"cdkeyType\":1,\"cdkeyUser\":\"admin\",\"createTime\":\"2021-05-26 15:00:00\",\"creator\":\"admin\",\"desc\":\"兑换吗\",\"contents\":[{\"content_type\":1,\"count\":10},{\"content_type\":2,\"count\":20}],\"expireTime\":\"2021-05-26 19:00:00\",\"totalExchangeNum\":10}", Router)
	fmt.Println("result =", string(result))
}

//创建多次兑换
func TestCreateCdkey2(t *testing.T) {
	Router := router.Routers()
	result := Post("/cdkey/createCdkey", "{\"cdkeyType\":2,\"cdkeyUser\":\"admin\",\"createTime\":\"2021-05-26 15:00:00\",\"creator\":\"admin\",\"desc\":\"兑换吗\",\"contents\":[{\"content_type\":1,\"count\":10},{\"content_type\":2,\"count\":20}],\"expireTime\":\"2021-05-30 19:00:00\",\"totalExchangeNum\":3}", Router)
	fmt.Println("result =", string(result))
}

//创建无限次兑换
func TestCreateCdkey3(t *testing.T) {
	Router := router.Routers()
	result := Post("/cdkey/createCdkey", "{\"cdkeyType\":3,\"cdkeyUser\":\"admin\",\"createTime\":\"2021-05-26 15:00:00\",\"creator\":\"admin\",\"desc\":\"兑换吗\",\"contents\":[{\"content_type\":1,\"count\":10},{\"content_type\":2,\"count\":20}],\"expireTime\":\"2021-05-30 19:00:00\",\"totalExchangeNum\":3}", Router)
	fmt.Println("result =", string(result))
}
func TestGetCdkeyDetails(t *testing.T) {
	Router := router.Routers()
	result := Get("/cdkey/getCdkeyDetails?cdkey=QEFI1VG8", Router)
	fmt.Println("result =", string(result))
}
func TestVerifyCdkey(t *testing.T) {
	Router := router.Routers()
	result := Get("/cdkey/verifyCdkey?cdkey=VQ0DND1I&userId=admin", Router)
	g := model.GeneralReward{}
	err := proto.Unmarshal(result, &g)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	fmt.Printf("result = %+v \n", &g)
}
func TestVerifyCdkey2(t *testing.T) {
	Router := router.Routers()
	result := Get("/cdkey/verifyCdkey?cdkey=VQ0DND3I&userId=431d6f79-9c8e-4f2e-84c5-431d3b315709", Router)
	g := model.GeneralReward{}
	err := proto.Unmarshal(result, &g)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("result = %+v \n", &g)
}
func TestVerifyCdkey3(t *testing.T) {
	Router := router.Routers()
	result := Get("/cdkey/verifyCdkey?cdkey=A8HQN2J0", Router)
	for i := 0; i < 10; i++ {
		result = Get("/cdkey/verifyCdkey?cdkey=A8HQN2J0", Router)
		g := model.GeneralReward{}
		err := proto.Unmarshal(result, &g)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func TestRegester(t *testing.T) {
	Router := router.Routers()
	result := Post("/user/registerUser", "{ \"id\":\"asdas\"}", Router)
	fmt.Println("result =", string(result))
}

func TestLogin(t *testing.T) {
	Router := router.Routers()
	result := Post("/user/login", "{ \"id\":\"492be816-8ebc-46eb-8625-c297258eb19b\"}", Router)
	fmt.Println("result =", string(result))
}
