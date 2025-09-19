package main

import (
	"log"

	"github.com/go-redis/redis"
)
func NewRedisClient(redisUrl string) *redis.Client{
	opt,err:=redis.ParseURL(redisUrl);
	if(err!=nil){
		log.Fatal("error connecting to redis...")
	}

	client:=redis.NewClient(opt)
	if err:=client.Ping().Err();err!=nil{
		log.Fatalf("redis ping eror",err)
	}
	return client
}


// func PublishMessage(client *redis.Client,channel string,payload []byte){
// 	if err:=client.Publish(context.Background(),channel,payload).Err();err!=nil{
// 		log.Printf("redis publish error %v,err")
// 	}
	
// }


func PublishMessage(client *redis.Client,channel string,payload []byte){
	// if err:=client.Publish(channel,payload).Err(); err!=nil{
	// 	log.Printf("redis publish error %v",err)
	// }

	err:=client.Publish(channel,payload).Err();
	if err!=nil{
		log.Printf("Redis publish error %v",err)
	}
}