package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"runtime"
	"testing"

	"go-project/demo3/controller"
	model "go-project/demo3/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
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
	Router := controller.Routers()
	result := Post("/cdkey/createCdkey", "{\"cdkeyType\":1,\"cdkeyUser\":\"admin\",\"createTime\":\"2021-05-26 15:00:00\",\"creator\":\"admin\",\"desc\":\"兑换吗\",\"contents\":[{\"content_type\":1,\"count\":10},{\"content_type\":2,\"count\":20}],\"expireTime\":\"2021-05-26 19:00:00\",\"totalExchangeNum\":10}", Router)
	fmt.Println("result =", string(result))
}

//创建多次兑换
func TestCreateCdkey2(t *testing.T) {
	Router := controller.Routers()
	result := Post("/cdkey/createCdkey", "{\"cdkeyType\":1,\"cdkeyUser\":\"admin\",\"createTime\":\"2021-05-26 15:00:00\",\"creator\":\"admin\",\"desc\":\"兑换吗\",\"contents\":[{\"content_type\":1,\"count\":10},{\"content_type\":2,\"count\":11}],\"expireTime\":\"2021-05-30 19:00:00\",\"totalExchangeNum\":3}", Router)
	fmt.Println("result =", string(result))
}

//创建无限次兑换
func TestCreateCdkey3(t *testing.T) {
	Router := controller.Routers()
	result := Post("/cdkey/createCdkey", "{\"cdkeyType\":3,\"cdkeyUser\":\"admin\",\"createTime\":\"2021-05-26 15:00:00\",\"creator\":\"admin\",\"desc\":\"兑换吗\",\"contents\":[{\"content_type\":1,\"count\":10},{\"content_type\":2,\"count\":20}],\"expireTime\":\"2021-05-30 19:00:00\",\"totalExchangeNum\":3}", Router)
	fmt.Println("result =", string(result))
}
func TestGetCdkeyDetails(t *testing.T) {
	Router := controller.Routers()
	result := Get("/cdkey/getCdkeyDetails?cdkey=QEFI1VG8", Router)
	fmt.Println("result =", string(result))
}
func TestVerifyCdkey(t *testing.T) {
	Router := controller.Routers()
	result := Get("/cdkey/verifyCdkey?cdkey=VQ0DND1I&userId=admin", Router)
	fmt.Println("result ===", result)
	g := model.GeneralReward{}
	err := proto.Unmarshal(result, &g)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	fmt.Printf("result = %+v \n", &g)
}
func TestVerifyCdkey2(t *testing.T) {
	Router := controller.Routers()
	result := Get("/cdkey/verifyCdkey?cdkey=VRM66R6O&userId=6776af39-8503-49bd-92c1-acc303e5fea7", Router)
	fmt.Println(result)
	g := model.GeneralReward{}
	err := proto.Unmarshal(result, &g)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("result = %+v \n", &g)
}
func TestVerifyCdkey3(t *testing.T) {
	Router := controller.Routers()
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
	Router := controller.Routers()
	result := Post("/user/registerUser", "{ \"id\":\"\"}", Router)
	fmt.Println("result =", string(result))
	runtime.GC()
}

func TestLogin(t *testing.T) {
	Router := controller.Routers()
	result := Post("/user/login", "{ \"id\":\"6776af39-8503-49bd-92c1-acc303e5fea7\"}", Router)
	fmt.Println("result =", string(result))
}
