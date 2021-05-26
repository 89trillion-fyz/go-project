package model

import (
	"reflect"
)

type CDKEY_TYPE int

const (
	ONCE   CDKEY_TYPE = iota + 1 //制定用户一次性消耗
	LIMIT                        //不指定用户限制兑换次数
	ALWAYS                       //不限用户不限次数兑换

)

type Content struct {
	Item  string `json:"item" binding:"required"`  //物品
	Count string `json:"count" binding:"required"` //数量
}
type Exchanger struct {
	User         string    `json:"user" binding:"required"`
	ExchangeTime LocalTime `json:"exchangeTime" binding:"required"`
}
type CdkeyModel struct {
	CdkeyType          CDKEY_TYPE  `json:"cdkeyType" binding:"required"`
	CdkeyUser          string      `json:"cdkeyUser"`                      //指定用户一次性
	Cdkey              string      `json:"cdkey"`                          //礼品码
	CreateTime         LocalTime   `json:"createTime" binding:"required" ` //创建时间
	Creator            string      `json:"creator" binding:"required"`     //创建人
	Desc               string      `json:"desc" binding:"required"`        //礼品描述
	Contents           []Content   `json:"contents" binding:"required" `
	ExpireTime         LocalTime   `json:"expireTime" binding:"required"`       //有效期
	TotalExchangeNum   int         `json:"totalExchangeNum" binding:"required"` //兑换次数
	AlreadyExchangeNum int         `json:"alreadyExchangeNum"`                  //已经兑换次数
	ExchangeList       []Exchanger `json:"exchangeList"`                        //领取列表

}

func (c CdkeyModel) IsEmpty() bool {
	return reflect.DeepEqual(c, CdkeyModel{})
}
