package api

import (
	"encoding/json"
	"fmt"
	"go-project/demo3/core/strategy"
	"go-project/demo3/global"
	"go-project/demo3/model"
	protoModel "go-project/demo3/proto"
	"go-project/demo3/utils"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/gin-gonic/gin"
)

const letterBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func initCdkey() (string, error) {
	cdkey := RandStringBytes(8)
	fmt.Println("initCdkey cdkey", cdkey)
	//判断redis是否生成过
	for {
		cdkey = RandStringBytes(8)
		if val, _ := global.GB_REDIS.Get(cdkey).Result(); val == "" {
			fmt.Println("之前没有生成过")
			return cdkey, nil
		}
	}
}
func CreateCdkey(c *gin.Context) {
	var cdkeyModel model.CdkeyModel
	_ = c.ShouldBindJSON(&cdkeyModel)
	fmt.Println("cdkeyModel : ", cdkeyModel)
	cdkey, err := initCdkey()
	fmt.Println("cdkey :", cdkey)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	//防止端上传参数错误
	if cdkeyModel.CdkeyType == model.ONCE {
		cdkeyModel.TotalExchangeNum = 1
		if cdkeyModel.CdkeyUser == "" {
			utils.FailWithMessage("请填写指定用户名", c)
		}
	}
	cdkeyModel.Cdkey = cdkey
	byteJson, err := json.Marshal(cdkeyModel)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	str, err := global.GB_REDIS.Set(cdkey, string(byteJson), -1).Result()
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("set str ", str)
	utils.Result(utils.SUCCESS, cdkey, "ok", c)
}
func GetCdkModel(cdkey string) (model.CdkeyModel, error) {
	var cdkeyModel model.CdkeyModel
	cdkeyModelStr, err := global.GB_REDIS.Get(cdkey).Result()
	if err != nil {
		return cdkeyModel, err
	}
	if err := json.Unmarshal([]byte(cdkeyModelStr), &cdkeyModel); err != nil {
		return cdkeyModel, err
	}
	fmt.Println("cdkeyModelStr", cdkeyModelStr, "cdkeyModel", cdkeyModel)
	return cdkeyModel, nil
}
func GetCdkeyDetails(c *gin.Context) {
	cdkey := c.Query("cdkey")
	if err := utils.VerfiyQuery(cdkey); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	cdkeyModel, err := GetCdkModel(cdkey)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.Result(utils.SUCCESS, cdkeyModel, "ok", c)

}

func VerifyCdkey(c *gin.Context) {
	cdkey := c.Query("cdkey")
	userId := c.Query("userId")

	//校验cdkey是否合法
	match, err := regexp.MatchString("^[A-Z0-9]{8}$", cdkey)
	if err != nil {
		c.ProtoBuf(http.StatusOK, &protoModel.GeneralReward{Code: utils.ERROR, Msg: err.Error()})
		return
	}
	if !match {
		c.ProtoBuf(http.StatusOK, &protoModel.GeneralReward{Code: utils.ERROR, Msg: "礼品码格式错误"})
		return
	}
	cdkeyModel, err := GetCdkModel(cdkey)
	if cdkeyModel.IsEmpty() {
		c.ProtoBuf(http.StatusOK, &protoModel.GeneralReward{Code: utils.ERROR, Msg: "礼包码不存在"})
		return
	}
	fmt.Println("VerifyCdkey cdkeymodel", cdkeyModel)
	//执行兑换流程
	strategyObj := strategy.NewStrategyContext(cdkeyModel.CdkeyType, strategy.ExchangeDetails{User: userId, ExchangeTime: model.LocalTime(time.Now())})
	fmt.Printf("strategyObj ==is of type %T \n", strategyObj.ExchangeStrategy)
	var user model.User
	if user, err = strategyObj.ExchangeStrategy.Exchange(&cdkeyModel); err != nil {
		//TODO 异常问题
		c.ProtoBuf(http.StatusOK, &protoModel.GeneralReward{Code: utils.ERROR, Msg: err.Error()})
		return
	}
	changeMap := make(map[uint32]uint64)
	for _, content := range cdkeyModel.Contents {
		changeMap[uint32(content.ContentType)] = content.Count
	}
	pb := &protoModel.GeneralReward{
		Code:    utils.SUCCESS,
		Msg:     "ok",
		Changes: changeMap,
		Balance: user.ContentMap,
		Counter: nil,
		Ext:     "",
	}
	byteSlice, err := proto.Marshal(pb)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
	}
	fmt.Println("before rsp ", cdkeyModel, "byteSlice", byteSlice)
	c.ProtoBuf(http.StatusOK, pb)
}
