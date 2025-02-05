package repository

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"note-echo/src/application"
	"note-echo/src/model"
)

// Collection 连接 MongoDB, 连接指定的文档集合
func Collection(ctx echo.Context) *mongo.Collection {
	// 从配置文件中读取连接配置
	uri := "mongodb://" + application.App.Mongo.Username + ":" + application.App.Mongo.Password + "@" + application.App.Mongo.Host + ":" + application.App.Mongo.Port + "/"

	// 连接 MongoDB 数据库
	client, err := mongo.Connect(ctx.Request().Context(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
	}

	// 连接配置文件指定的数据库和文档集合
	collection := client.Database(application.App.Mongo.Database).Collection(application.App.Mongo.Collection)

	return collection
}

// List 查询所有 Linux 命令
func List(ctx echo.Context) []model.Command {
	// 获取集合连接
	collection := Collection(ctx)

	// 查询所有 Linux 命令
	cursor, err := collection.Find(ctx.Request().Context(), bson.D{})
	if err != nil {
		log.Println(err)
	}

	// 结构体数组
	var commandArray []model.Command
	// 返回值 cursor 相当于一个指针, 需要 Next() 遍历一个一个获取数据
	for cursor.Next(ctx.Request().Context()) {
		command := model.Command{}
		err := cursor.Decode(&command)
		if err != nil {
			log.Println(err)
		}
		commandArray = append(commandArray, command)
	}

	return commandArray
}
