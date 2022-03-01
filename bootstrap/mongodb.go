package bootstrap

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goAdmin/global"
	"strconv"
	"time"
)

type Database struct {
	Mongo *mongo.Client
}

var DB *Database

//初始化
func InitMongo() {
	DB = &Database{
		Mongo: SetConnect(),
	}
	global.App.PoolMongo = DB.Mongo
}

// 连接设置
func SetConnect() *mongo.Client {
	//uri := "mongodb://dcjt:dcjt123dalgurak@127.0.0.1:27017/admin"
	uri := global.App.Config.Mongodb.MgoUrl

	fmt.Println("mongodbUrl is ", uri)
	clientNum, _ := strconv.Atoi(global.App.Config.Mongodb.ClientNum)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(uint64(clientNum))) // 连接池
	if err != nil {
		fmt.Println("create poolMongodb is failed. err:", err)
	}
	return client
}
