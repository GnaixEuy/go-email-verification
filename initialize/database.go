package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MonGoDBClient *mongo.Client
var RedisClient *redis.Client

func initMongoClient(ctx context.Context) (err error) {
	conn := options.Client().ApplyURI(GetConfig().DbUrI)
	MonGoDBClient, err = mongo.Connect(ctx, conn)
	if err != nil {
		return err
	}
	err = MonGoDBClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB!")
	return err
}

func initRedisClient(ctx context.Context) (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     GetConfig().RedisUrI,
		Password: GetConfig().RedisPass,
		DB:       GetConfig().RedisDb,
	})
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println("Connected to Redis!")
	return err
}

func InitDBClient(ctx context.Context) {
	err := initMongoClient(ctx)
	if err != nil {
		panic(err)
	}
	err = initRedisClient(ctx)
	if err != nil {
		panic(err)
	}
}

func CloseDBClient(ctx context.Context) {
	_ = MonGoDBClient.Disconnect(ctx)
	_ = RedisClient.Close()
}
