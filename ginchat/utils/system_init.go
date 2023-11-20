package utils

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"context"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app inited")

}

func InitMySQl() {
	//自定义日志模版 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,
		},
	)

	dsn := "root:88888888@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn),
		&gorm.Config{Logger: newLogger})
	fmt.Println("MySQl inited....")
	//user := models.UserBasic{}
	//DB.Find(user)
	//fmt.Println(user)
}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish...", msg)
	err = Red.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println("Publish err。。。。。", err)
	}
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	fmt.Println("Subscribe...", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println("Subscribe  err:", err)
		// Handle the error here, such as returning it or logging it
		return "", err
	}
	fmt.Println("Subscribe...", msg.Payload)

	return msg.Payload, nil // No error occurred, so return nil for the error.
}
