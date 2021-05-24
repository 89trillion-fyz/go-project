package api

import (
	"demo1/global"
	"demo1/model"
	response "demo1/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func FindByRarityAndLock(c *gin.Context) {
	rarity := c.Query("rarity")
	lock := c.Query("lock")
	fmt.Println("rarity", rarity, "lock", lock)
	armyArr := []model.Army{}
	jsondata := global.G_JSONDATA
	for _, army := range jsondata {
		if army.Rarity == rarity && army.UnlockArena == lock {
			armyArr = append(armyArr, army)
		}
	}
	response.Result(response.SUCCESS, armyArr, "ok", c)
}

func FindRarityById(c *gin.Context) {
	id := c.Query("id")
	fmt.Println("id", id)
	rarity := ""
	jsondata := global.G_JSONDATA
	for key, army := range jsondata {
		if key == id {
			rarity = army.Rarity
			break
		}
	}
	response.Result(response.SUCCESS, rarity, "ok", c)
}

func FindQualityById(c *gin.Context) {
	id := c.Query("id")
	fmt.Println("id", id)
	jsondata := global.G_JSONDATA
	quality := ""
	for key, army := range jsondata {
		if key == id {
			quality = army.Quality
			break
		}
	}
	response.Result(response.SUCCESS, quality, "ok", c)
}

func FindByLock(c *gin.Context) {
	lock := c.Query("lock")
	fmt.Println("lock", lock)
	jsondata := global.G_JSONDATA
	armyArr := []model.Army{}
	for _, army := range jsondata {
		if army.UnlockArena == lock {
			armyArr = append(armyArr, army)
		}
	}
	response.Result(response.SUCCESS, armyArr, "ok", c)
}
