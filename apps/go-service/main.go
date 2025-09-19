package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main(){
	cfg:=SetConfig()
	redisClient:=NewRedisClient(cfg.RedisUrl);
	defer redisClient.Close()

	out:=make(chan []byte,512)

	ctx,cancel:=context.WithCancel(context.Background());
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		ConnectAndStream(ctx,cfg.BinanceUrl,out);
	}()


	go func(){
		defer wg.Done()
		for{
			select{
			case <-ctx.Done():
					return
			case msg:=<-out:
				PublishMessage(redisClient,"market.ticker.BTCUSDT.binance",msg)
			}
		}
	}()

	stop:=make(chan os.Signal,1)
	signal.Notify(stop,syscall.SIGINT,syscall.SIGTERM)
	<-stop
	log.Println("shutting down....")
	cancel()
}
