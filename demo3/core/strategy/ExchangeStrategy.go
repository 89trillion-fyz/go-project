package strategy

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-project/demo3/global"
	"go-project/demo3/model"
	"time"
)

//兑换策略

type ExchangeStrategy interface {
	Exchange(model *model.CdkeyModel) error
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
func commonExchange(cdkeyModel *model.CdkeyModel, details *ExchangeDetails) error {
	cdkeyModel.AlreadyExchangeNum += 1
	cdkeyModel.ExchangeList = append(cdkeyModel.ExchangeList, model.Exchanger{User: details.User, ExchangeTime: details.ExchangeTime})
	byteJson, err := json.Marshal(cdkeyModel)
	if err != nil {
		return err
	}
	result := global.GB_REDIS.Set(cdkeyModel.Cdkey, string(byteJson), -1)
	fmt.Println("exchange result : =", result)
	return nil
}
func (s OnceStrategy) Exchange(cdkeyModel *model.CdkeyModel) error {
	if cdkeyModel.TotalExchangeNum-cdkeyModel.AlreadyExchangeNum > 0 && cdkeyModel.CdkeyUser == s.ExchangeDetails.User && time.Now().Before(time.Time(cdkeyModel.ExpireTime)) {
		//兑换逻辑
		return commonExchange(cdkeyModel, &s.ExchangeDetails)
	}
	return errors.New("礼包码无效")
}
func (s LimitStrategy) Exchange(cdkeyModel *model.CdkeyModel) error {
	if cdkeyModel.TotalExchangeNum-cdkeyModel.AlreadyExchangeNum > 0 && time.Now().Before(time.Time(cdkeyModel.ExpireTime)) {
		//兑换逻辑
		return commonExchange(cdkeyModel, &s.ExchangeDetails)
	}
	return errors.New("礼包码无效")
}

func (s AlwaysStrategy) Exchange(cdkeyModel *model.CdkeyModel) error {
	if time.Now().Before(time.Time(cdkeyModel.ExpireTime)) {
		//兑换逻辑
		return commonExchange(cdkeyModel, &s.ExchangeDetails)
	}
	return errors.New("礼包码无效")
}
