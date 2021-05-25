package api

import (
	"demo1/global"
	"demo1/model"
	"demo1/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func FindByRarityAndLock(c *gin.Context) {
	rarity := c.Query("rarity")
	lock := c.Query("lock")
	if err := utils.VerfiyQuery(rarity, lock); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("rarity", rarity, "lock", lock)
	armySlice := make([]model.Army, 0)
	jsondata := global.G_JSONDATA
	for _, army := range jsondata {
		if army.Rarity == rarity && army.UnlockArena == lock {
			armySlice = append(armySlice, army)
		}
	}
	utils.Result(utils.SUCCESS, armySlice, "ok", c)
}

func FindRarityById(c *gin.Context) {
	id := c.Query("id")
	fmt.Println("id", id)
	if err := utils.VerfiyQuery(id); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	jsonData := global.G_JSONDATA
	rarity := jsonData[id]
	utils.Result(utils.SUCCESS, rarity, "ok", c)
}

func FindQualityById(c *gin.Context) {
	id := c.Query("id")
	fmt.Println("id", id)
	if err := utils.VerfiyQuery(id); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	jsonData := global.G_JSONDATA
	rarity := jsonData[id]
	quality := rarity.Quality
	utils.Result(utils.SUCCESS, quality, "ok", c)
}

func FindByLock(c *gin.Context) {
	lock := c.Query("lock")
	fmt.Println("lock", lock)
	if err := utils.VerfiyQuery(lock); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	jsonData := global.G_JSONDATA
	armyArr := make([]model.Army, 0)
	for _, army := range jsonData {
		if army.UnlockArena == lock {
			armyArr = append(armyArr, army)
		}
	}
	utils.Result(utils.SUCCESS, armyArr, "ok", c)
}
