package initialize

import (
	"context"
	"fmt"
	"go-project/demo3/global"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongoDB(uri string, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}

	return client.Database(name), nil
}

func DbClient(dbName string) (*mongo.Database, error) {
	fmt.Println("DbClient ===", global.GB_GONFIG.Mongo.ApplyURI)
	return connectToMongoDB(global.GB_GONFIG.Mongo.ApplyURI, dbName, time.Duration(global.GB_GONFIG.Mongo.Timeout), global.GB_GONFIG.Mongo.PoolSize)
}
