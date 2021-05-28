package initialize

import (
	"context"
	"fmt"
	"go-project/demo3/global"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB(uri string, timeout time.Duration, num uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		fmt.Println("err..", err.Error())
	}
	if !reflect.DeepEqual(client, mongo.Client{}) {
		global.GB_MONGO = client
	}
}
