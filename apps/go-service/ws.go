package main

import (
	"context"
	"log"

	"github.com/gorilla/websocket"
)

func ConnectAndStream(ctx context.Context,wsUrl string,out chan<-[]byte){
	conn,_,err:=websocket.DefaultDialer.Dial(wsUrl,nil)
	if err!=nil{
		log.Printf("ws dial error %v",err)
	}

	for{
		_,msg,err:=conn.ReadMessage()
	if err!=nil{
		log.Printf("ws read error %v",err)
		_=conn.Close()
			break
	}

	select{
	case out <-msg:
	case <-ctx.Done():
		_=conn.Close()
		return
	}
	}



} 