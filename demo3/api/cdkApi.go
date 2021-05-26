package api

import (
	"encoding/json"
	"fmt"
	"go-project/demo3/core/strategy"
	"go-project/demo3/global"
	"go-project/demo3/model"
	"go-project/demo3/utils"
	"math/rand"
	"regexp"
	"time"

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
func GetCdkModel(c *gin.Context, cdkey string) model.CdkeyModel {
	var cdkeyModel model.CdkeyModel
	cdkeyModelStr, err := global.GB_REDIS.Get(cdkey).Result()
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return cdkeyModel
	}
	if err := json.Unmarshal([]byte(cdkeyModelStr), &cdkeyModel); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return cdkeyModel
	}
	fmt.Println("cdkeyModelStr", cdkeyModelStr, "cdkeyModel", cdkeyModel)
	return cdkeyModel
}
func GetCdkeyDetails(c *gin.Context) {
	cdkey := c.Query("cdkey")
	if err := utils.VerfiyQuery(cdkey); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	cdkeyModel := GetCdkModel(c, cdkey)
	utils.Result(utils.SUCCESS, cdkeyModel, "ok", c)

}

func VerifyCdkey(c *gin.Context) {
	cdkey := c.Query("cdkey")
	user := c.Query("user")

	//校验cdkey是否合法
	match, err := regexp.MatchString("^[A-Z0-9]{8}$", cdkey)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	if !match {
		utils.FailWithMessage("礼品码格式错误", c)
		return
	}
	cdkeyModel := GetCdkModel(c, cdkey)
	if cdkeyModel.IsEmpty() {
		utils.FailWithMessage("礼包码不存在", c)
		return
	}
	fmt.Println("VerifyCdkey cdkeymodel", cdkeyModel)
	//校验礼包码是否在有效期内
	if time.Time(cdkeyModel.ExpireTime).Before(time.Now()) || cdkeyModel.TotalExchangeNum-cdkeyModel.AlreadyExchangeNum == 0 {
		utils.FailWithMessage("礼包码已经失效", c)
		return
	}
	//执行兑换流程
	strategyObj := strategy.NewStrategyContext(cdkeyModel.CdkeyType, strategy.ExchangeDetails{User: user, ExchangeTime: model.LocalTime(time.Now())})
	if err := strategyObj.ExchangeStrategy.Exchange(&cdkeyModel); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.Result(utils.SUCCESS, cdkeyModel, "ok", c)
}
