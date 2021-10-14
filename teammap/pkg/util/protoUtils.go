package util

import (
	"log"
	"teammap/pkg/errorlog"
	"teammap/pkg/myerr"
	"teammap/protocol/protobuf/response"

	"google.golang.org/protobuf/proto"
)

// 用来放客户端展示的奖励的，扣的任何东西严禁放进去
type PerChange struct {
	JoinId uint32 // 获得的物品完整id，eg：1010102 1级士兵
	Amount int32  // 获得的物品数量
}

// 用来放客户端获得的英雄的信息
type HeroInf struct {
	Name  uint32 // 获得的英雄name，eg：101
	Star  uint32 // 获得的英雄star，eg：1
	Level uint32 // 获得的英雄level，eg：1
	Skill uint32 // 获得的英雄skill，eg：1
}

// 回复 错误码
func MakeErrorResponse(err *myerr.MyErr) []byte {
	errorlog.LogMyError(err)
	return MakeGeneralResponse(err.Code, err.Message)
}

// 回复通用消息
func MakeGeneralResponse(code int, msg string) []byte {
	general := &response.General{
		Code: int32(code),
		Msg:  msg,
	}

	bytes, err := proto.Marshal(general)
	if err != nil {
		// 出错
		log.Print("proto 转bytes 失败")
	}

	return bytes
}

// 回复 通用奖励消息
func MakeGeneralRewardResponse(
	code int,
	msg string,
	changesArr []*PerChange,
	balanceMap map[uint32]uint64,
	counterMap map[uint32]uint64,
	ext string) []byte {

	reward := &response.GeneralReward{
		Code:    int32(code),
		Msg:     msg,
		Changes: []*response.PerChange{},
		Balance: balanceMap,
		Counter: counterMap,
		Ext:     ext,
	}
	for _, data := range changesArr {
		reward.Changes = append(reward.Changes, &response.PerChange{
			JoinId: data.JoinId, Amount: data.Amount,
		})
	}

	bytes, err := proto.Marshal(reward)
	if err != nil {
		// 出错
		log.Print("proto 转bytes 失败")
	}

	return bytes
}

// 回复 通用状态消息
func MakeGeneralStateResponse(
	code int,
	msg string,
	counterMap map[uint32]uint64,
	stateMap map[int32]map[int32]uint32,
	ext string) []byte {

	state := &response.GeneralState{
		Code:    int32(code),
		Msg:     msg,
		Counter: counterMap,
		State:   convertStateDict(stateMap),
		Ext:     ext,
	}

	bytes, err := proto.Marshal(state)
	if err != nil {
		// 出错
		log.Print("proto 转bytes 失败")
	}

	return bytes
}

func MakeGenerateLayoutResponse(
	code int,
	msg string,
	balanceMap map[uint32]uint64,
	counterMap map[uint32]uint64,
	stateMap map[int32]map[int32]uint32,
	ext []uint32,
	changePos *response.SyncPos,
	bdInf *response.BdInformation,
	heroInf *response.HeroInformation,
) []byte {

	layoutResp := &response.GeneralLayout{
		Code:       int32(code),
		Msg:        msg,
		BdInf:      bdInf,
		ChangedPos: changePos,
		Ext:        ext,
		Balance:    balanceMap,
		HeroInf:    heroInf,
		Counter:    counterMap,
		State:      convertStateDict(stateMap),
	}

	bytes, err := proto.Marshal(layoutResp)
	if err != nil {
		// 出错
		log.Print("proto 转bytes 失败")
	}

	return bytes
}

// map 转 CounterDetail数组
func convertStateDict(stateDict map[int32]map[int32]uint32) []*response.StateDetail {

	var stateArray []*response.StateDetail

	// 外层循环 stateType -> map
	for k, innerMap := range stateDict {
		stateArray = append(
			stateArray,
			&response.StateDetail{
				Type:  response.StateType(k),
				State: innerMap,
			})
	}

	return stateArray
}
