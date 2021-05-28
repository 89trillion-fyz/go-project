package api

import (
	"context"
	"fmt"
	"reflect"

	"go-project/demo3/initialize"
	"go-project/demo3/model"
	"go-project/demo3/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err == nil {
		client, err := initialize.DbClient(model.DB_NAME)
		if err != nil {
			utils.FailWithMessage(err.Error(), c)
			return
		}
		filter := bson.D{{"id", user.Id}}
		var findUser model.User
		_ = client.Collection(model.C_NAME_USER).FindOne(context.TODO(), filter).Decode(&findUser)
		if reflect.DeepEqual(findUser, model.User{}) {
			fmt.Println("not register 。。。")
			//初始化
			findUser.Id = uuid.New().String()
		}
		_, err = client.Collection(model.C_NAME_USER).InsertOne(context.TODO(), findUser)
		if err != nil {
			utils.FailWithMessage(err.Error(), c)
		}
		utils.OkWithData(findUser, c)
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}
func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err == nil {
		client, err := initialize.DbClient(model.DB_NAME)
		if err != nil {
			utils.FailWithMessage(err.Error(), c)
			return
		}
		if user.Id == "" {
			utils.FailWithMessage("请输入用户id", c)
			return
		}
		fmt.Println("user", user)
		filter := bson.D{{"id", user.Id}}
		var findUser model.User
		_ = client.Collection(model.C_NAME_USER).FindOne(context.TODO(), filter).Decode(&findUser)
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
