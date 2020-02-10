package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
)

//任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

//一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   //任务名
	Command   string    `bson:"command"`   //shell命令
	Err       string    `bson:"err"`       //脚本错误
	Content   string    `bson:"content"`   //脚本输出
	TimePoint TimePoint `bson:"timePoint"` //执行时间点
}

func main() {

	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		record     *LogRecord
		logArr     []interface{}
		result     *mongo.InsertManyResult
		insertid   interface{}
		docld      objectid.ObjectID
	)

	//1.建立连接
	if client, err = mongo.Connect(context.TODO(), "mongodb://111.229.213.40:27017", clientopt.ConnectTimeout(5*time.Second)); err != nil {
		fmt.Println(err)
		return
	}

	//2.选择数据库
	database = client.Database("cron")

	//3.选择表
	collection = database.Collection("log")

	//4.插入记录
	record = &LogRecord{
		JobName:   "job8",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 20},
	}

	//5.批量插入多条
	logArr = []interface{}{record, record, record}

	//插入
	if result, err = collection.InsertMany(context.TODO(), logArr); err != nil {
		fmt.Println(err)
		return
	}

	for _, insertid = range result.InsertedIDs {
		//反射
		docld = insertid.(objectid.ObjectID)
		fmt.Println("自增ID：", docld.Hex())
		//snowflake(雪花算法)：推特开源，用来生成id
		//毫秒\微妙的当前时间+机器的ID+当前毫秒/微秒内的自增ID（每当毫秒变化了，会重置为0，继续自增）
	}

}
