package httpResponse

import "iw/internal/model"

// AttackLogRep 战斗日志响应
type AttackLogRep struct {
	TS                 int64  `json:"ts"`
	UID                string `json:"userId"`
	UName              string `json:"name"`
	UProfile           string `json:"profile"`
	UGold              int    `json:"gold"`
	UWood              int    `json:"wood"`
	UScore             int    `json:"score"`
	OScore             int    `json:"oscore,omitempty"`
	UGem               int    `json:"gem"`
	Result             int    `json:"result"`
	Revenge            int    `json:"revenge,omitempty"`
	LogId              string `json:"logid,omitempty"`
	RevengeType        int    `json:"rType"`
	DestroyBattleScore int    `json:"dstBS"`
	TimeBattleScore    int    `json:"timBS"`
	KillBattleScore    int    `json:"kilBS"`
	LostBattleScore    int    `json:"losBS"`
	CombatBattleScore  int    `json:"cbtBS"`
	TotalBattleScore   int    `json:"ttlBS"`
	LevelBattleScore   int    `json:"lvlBS"`
	AdvanceBP          int    `json:"advanceBP,omitempty"` //是否购买battlePass 0否 1是
}

type WeatherResponse struct {
	Expire      int64               `json:"expire"`      // 客户端数据过期时间戳
	WeatherInfo []model.WeatherData `json:"weatherInfo"` // 天气信息
}

type InfRankResponse struct {
	List    []*model.UserInfo `json:"list"`
	Self    *model.RankSelf   `json:"self"`
	Stage   int               `json:"stage,omitempty"`
	MaxPage int32             `json:"maxPage"`
}

type PointInfoResponse struct {
	List []*model.UserInfo `json:"list"`
	Self *model.RankSelf   `json:"self"`
}

// -----------------------

type PVPSelfRank struct {
	Rank  int32 `json:"ranking"`
	Score int32 `json:"score"`
}

type PVPRankResponse struct {
	List    []*model.UserInfo `json:"list"`
	Self    *PVPSelfRank      `json:"self"`
	Stage   int               `json:"stage,omitempty"`
	MaxPage int32             `json:"maxPage"`
}

// -----------------------

type RedeemCodeClaim struct {
	UserId  string `json:"user_id"`
	Cdkey   string `json:"cdkey"`
	AddTime int64  `json:"addtime"`
}
