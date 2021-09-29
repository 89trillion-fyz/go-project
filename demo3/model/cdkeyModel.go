package model

import (
	"reflect"
)

type CDKEY_TYPE uint32
type CONTENT_TYPE uint32

const (
	ONCE   CDKEY_TYPE = iota + 1 //制定用户一次性消耗
	LIMIT                        //不指定用户限制兑换次数
	ALWAYS                       //不限用户不限次数兑换

)
const (
	GLOD CONTENT_TYPE = iota + 1 //金币
	DIAMOND
)

type Content struct {
	ContentType CONTENT_TYPE `json:"content_type" binding:"required,oneof=1 2"` //物品类型
	Count       uint64       `json:"count"  binding:"required"`                 //数量
}
type Exchanger struct {
	User         string    `json:"user" binding:"required"`
	ExchangeTime LocalTime `json:"exchangeTime" binding:"required"`
}
type CdkeyModel struct {
	CdkeyType          CDKEY_TYPE  `json:"cdkeyType" binding:"oneof=1 2 3"`
	CdkeyUser          string      `json:"cdkeyUser"`                      //指定用户一次性
	Cdkey              string      `json:"cdkey"`                          //礼品码
	CreateTime         LocalTime   `json:"createTime" binding:"required" ` //创建时间
	Creator            string      `json:"creator" binding:"required"`     //创建人
	Desc               string      `json:"desc" binding:"required"`        //礼品描述
	Contents           []Content   `json:"contents" binding:"required,dive,required"`
	ExpireTime         LocalTime   `json:"expireTime" binding:"required"`             //有效期
	TotalExchangeNum   int         `json:"totalExchangeNum" binding:"required,min=0"` //兑换次数
	AlreadyExchangeNum int         `json:"alreadyExchangeNum"`                        //已经兑换次数
	ExchangeList       []Exchanger `json:"exchangeList"`                              //领取列表

}

func (c CdkeyModel) IsEmpty() bool {
	return reflect.DeepEqual(c, CdkeyModel{})
}
