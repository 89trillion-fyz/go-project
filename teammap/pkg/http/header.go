package app

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"teammap/pkg/myerr"
	"teammap/pkg/util"
)

type HttpHeader struct {
	Gd  string  `json:"gaid"`    // 谷歌广告id 或 Apple idfa
	Ud  string  `json:"uid"`     // 用户uuid
	Vc  int     `json:"cvc"`     // 客户端版本号
	Sv  float64 `json:"svc"`     // 系统版本号
	Dv  string  `json:"device"`  // 设备名称
	Nw  string  `json:"network"` // 网络类型 wifi/lte/3g/2g/other
	Sc  string  `json:"simcode"` // sim卡 code
	Lg  string  `json:"lang"`    // 设备语言
	Ls  string  `json:"ls"`      // 用户设置的游戏语言
	Pf  string  `json:"pf"`      // 平台
	IP  string  `json:"ip"`      // client ip
	Cty string  `json:"country"` // 设备ip对应的国家
	Ap  int     `json:"appid"`   // appId
}

var RemoteHeader = &HttpHeader{Pf: "remote", Ap: 246}

// 从ws建立连接的http请求cookie里提取header info
func (g *Gin) WsHeaderInfo() (*HttpHeader, *myerr.MyErr) {
	info := &HttpHeader{}

	basicInfo, err := g.C.Cookie("Basic-Info")
	//修复url unescape 导致base64失效
	basicInfo = strings.ReplaceAll(strings.TrimSpace(basicInfo), " ", "+")
	if err != nil || len(basicInfo) < 1 {
		return info, myerr.LACK_OF_HEADER
	}

	decode, err := base64.StdEncoding.DecodeString(basicInfo)
	if err != nil {
		return info, myerr.INVALID_HEADER
	}

	err = json.Unmarshal(decode, info)
	if err != nil {
		return info, myerr.INVALID_HEADER
	}

	info.IP = g.C.ClientIP()
	info.Cty = util.GetCountryCodeByIp(info.IP)
	info.Pf = strings.ToLower(info.Pf)
	info.Ap = util.GetAppIdByPlatform(info.Pf)

	return info, myerr.SUCCESS
}

func (g *Gin) TestWsHeaderInfo() (*HttpHeader, *myerr.MyErr) {
	cvc := g.C.Query("cvc")
	gid := g.C.Query("gaid")
	info := &HttpHeader{
		Ap: 246,
		Pf: "android",
		Dv: "macbook",
		Vc: util.Atoi(cvc),
		Gd: gid,
	}
	return info, myerr.SUCCESS
}
