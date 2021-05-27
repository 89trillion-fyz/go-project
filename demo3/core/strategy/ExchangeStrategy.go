package strategy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-project/demo3/global"
	"go-project/demo3/initialize"
	"go-project/demo3/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//兑换策略

type ExchangeStrategy interface {
	Exchange(model *model.CdkeyModel) (model.User, error)
}
type StrategyContext struct {
	ExchangeStrategy ExchangeStrategy
}

func NewStrategyContext(cdkType model.CDKEY_TYPE, details ExchangeDetails) StrategyContext {
	instnce := new(StrategyContext)
	switch cdkType {
	case model.ONCE:
		instnce.ExchangeStrategy = NewOnceStrategy(details)
	case model.LIMIT:
		instnce.ExchangeStrategy = NewLimitStrategy(details)
	case model.ALWAYS:
		instnce.ExchangeStrategy = NewAlwaysStrategy(details)
	}
	return *instnce
}

type ExchangeDetails struct {
	User         string          //兑换用户
	ExchangeTime model.LocalTime //兑换时间
}
type OnceStrategy struct {
	ExchangeDetails ExchangeDetails
}

type LimitStrategy struct {
	ExchangeDetails ExchangeDetails
}

type AlwaysStrategy struct {
	ExchangeDetails ExchangeDetails
}

func NewOnceStrategy(details ExchangeDetails) OnceStrategy {
	instance := new(OnceStrategy)
	instance.ExchangeDetails = details
	return *instance
}
func NewLimitStrategy(details ExchangeDetails) LimitStrategy {
	instance := new(LimitStrategy)
	instance.ExchangeDetails = details
	return *instance
}
func NewAlwaysStrategy(details ExchangeDetails) AlwaysStrategy {
	instance := new(AlwaysStrategy)
	instance.ExchangeDetails = details
	return *instance
}
func commonExchange(cdkeyModel *model.CdkeyModel, details *ExchangeDetails) (model.User, error) {
	fmt.Println("commonExchange details", details.User)
	client, err := initialize.DbClient(model.DB_NAME)
	if err != nil {
		return model.User{}, err
	}
	filter := bson.D{{"id", details.User}}
	fmt.Println("filter", filter)
	var findUser model.User
	_ = client.Collection(model.C_NAME_USER).FindOne(context.TODO(), filter).Decode(&findUser)
	if &findUser == nil {
		return model.User{}, errors.New("用户未注册")
	}
	//用户添加钻石数量金币数量
	fmt.Println("findUser.ContentMap", findUser.ContentMap, "11111", findUser.ContentMap == nil)
	for _, content := range cdkeyModel.Contents {
		if findUser.ContentMap == nil {
			findUser.ContentMap = make(map[uint32]uint64)
		}
		switch content.ContentType {
		case model.GLOD:
			findUser.ContentMap[uint32(model.GLOD)] += content.Count
		case model.DIAMOND:
			findUser.ContentMap[uint32(model.DIAMOND)] += content.Count
		}
	}
	fmt.Println("兑换后用户状态", findUser)
	insertResult, err := client.Collection(model.C_NAME_USER).UpdateOne(context.TODO(), bson.D{{"id", findUser.Id}},
		bson.D{{"$set", bson.D{
			{"contentMap", findUser.ContentMap},
		}}})
	if err != nil {
		return model.User{}, err
	}
	fmt.Println("兑换结果", insertResult)
	cdkeyModel.AlreadyExchangeNum += 1
	cdkeyModel.ExchangeList = append(cdkeyModel.ExchangeList, model.Exchanger{User: details.User, ExchangeTime: details.ExchangeTime})
	byteJson, err := json.Marshal(cdkeyModel)
	if err != nil {
		return model.User{}, err
	}
	result := global.GB_REDIS.Set(cdkeyModel.Cdkey, string(byteJson), -1)
	fmt.Println("exchange result : =", result)
	return findUser, nil
}
func (s OnceStrategy) Exchange(cdkeyModel *model.CdkeyModel) (model.User, error) {
	if cdkeyModel.TotalExchangeNum-cdkeyModel.AlreadyExchangeNum > 0 && cdkeyModel.CdkeyUser == s.ExchangeDetails.User && time.Now().Before(time.Time(cdkeyModel.ExpireTime)) {
		//兑换逻辑
		return commonExchange(cdkeyModel, &s.ExchangeDetails)
	}
	return model.User{}, errors.New("礼包码无效")
}
func (s LimitStrategy) Exchange(cdkeyModel *model.CdkeyModel) (model.User, error) {
	if cdkeyModel.TotalExchangeNum-cdkeyModel.AlreadyExchangeNum > 0 && time.Now().Before(time.Time(cdkeyModel.ExpireTime)) {
		//兑换逻辑
		return commonExchange(cdkeyModel, &s.ExchangeDetails)
	}
	return model.User{}, errors.New("礼包码无效")
}

func (s AlwaysStrategy) Exchange(cdkeyModel *model.CdkeyModel) (model.User, error) {
	if time.Now().Before(time.Time(cdkeyModel.ExpireTime)) {
		//兑换逻辑
		return commonExchange(cdkeyModel, &s.ExchangeDetails)
	}
	return model.User{}, errors.New("礼包码无效")
}
