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
	uri := "mongodb://" +
		application.App.Mongodb.Username + ":" +
		application.App.Mongodb.Password + "@" +
		application.App.Mongodb.Host + ":" +
		application.App.Mongodb.Port + "/"

	// 连接 MongoDB 数据库
	client, err := mongo.Connect(ctx.Request().Context(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
	}

	// 连接配置文件指定的数据库和文档集合
	collection := client.Database(application.App.Mongodb.Database).Collection(application.App.Mongodb.Collection)

	return collection
}

// List 查询所有命令
func List(ctx echo.Context) []model.Command {
	// 获取集合连接
	collection := Collection(ctx)

	// 查询所有命令
	cursor, err := collection.Find(ctx.Request().Context(), bson.D{})
	if err != nil {
		log.Println(err)
	}

	// 命令结构体数组
	var commandArray []model.Command
	// 使用 cursor 指针遍历获取数据
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
