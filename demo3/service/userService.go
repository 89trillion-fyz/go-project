package service

import (
	"fmt"
	"reflect"

	"go-project/demo3/handler"
	"go-project/demo3/model"
	"go-project/demo3/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err == nil {
		var findUser model.User
		_ = handler.NewMgo(model.DB_NAME, model.C_NAME_USER).FindOne("id", user.Id).Decode(&findUser)
		if reflect.DeepEqual(findUser, model.User{}) {
			fmt.Println("not register 。。。")
			//初始化
			findUser.Id = uuid.New().String()
		}
		handler.NewMgo(model.DB_NAME, model.C_NAME_USER).InsertOne(findUser)
		utils.OkWithData(findUser, c)
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}
func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err == nil {
		if user.Id == "" {
			utils.FailWithMessage("请输入用户id", c)
			return
		}
		fmt.Println("user", user)
		var findUser model.User
		_ = handler.NewMgo(model.DB_NAME, model.C_NAME_USER).FindOne("id", user.Id).Decode(&findUser)
		fmt.Println("findUser", findUser)
		if reflect.DeepEqual(findUser, model.User{}) {
			utils.FailWithMessage("用户还没有注册", c)
			return
		}
		utils.OkWithData(findUser, c)
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}
