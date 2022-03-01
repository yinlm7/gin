package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/global"
)

func User(c *gin.Context) {
	fmt.Println(global.App.Config.App)
	fmt.Println(global.App)
	fmt.Println("aaa")
}

func Test(c *gin.Context) {
	method, _ := c.GetQuery("type")

	conn := global.App.PoolRedis.Get()
	defer conn.Close()

	var err error

	switch method {
	case "get":
		reply, err := conn.Do("GET", "conntest")
		fmt.Println(string(reply.([]byte)), err)
	case "set":
		reply, err := conn.Do("SET", "conntest", "zxcvb")
		fmt.Println(reply, err)
	}
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.String(200, "ok")
}

func MgoTest(c *gin.Context) {
	// 指定数据库名称
	db := global.App.PoolMongo.Database("demo")
	// 指定连接的集合
	col := db.Collection("demo")

	fmt.Println("mongoDbClient is ok", col)
	//// 做一个唯一查询
	//ret := col.FindOne(ctx,map[string]string{"name":"pika"})
	//// 获取的byte数据
	//raw,_ := ret.DecodeBytes()
	//// 转为字符串打印出来
	//fmt.Println(raw.String())

}
