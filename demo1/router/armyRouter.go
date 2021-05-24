package router

import (
	"demo1/api"
	"github.com/gin-gonic/gin"
)

func InitArmyRouter(Router *gin.RouterGroup) {
	ArmyRouter := Router.Group("army")
	{
		//1）输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
		ArmyRouter.GET("findByRarityAndLock", api.FindByRarityAndLock)
		//2）输入士兵id获取稀有度
		ArmyRouter.GET("findRarityById", api.FindRarityById)
		//3）输入士兵id获取战力
		ArmyRouter.GET("findQualityById", api.FindQualityById)
		//4）输入cvc获取所有合法的士兵
		//ArmyRouter.GET("findByCvc")
		//5）获取每个阶段解锁相应士兵的json数据
		ArmyRouter.GET("findByLock", api.FindByLock)
	}

}
