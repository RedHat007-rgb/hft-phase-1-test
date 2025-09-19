package main

import "os"

type Config struct{
	BinanceUrl string
	RedisUrl string
}  


func SetConfig() *Config{
	Binanceurl:=getEnv("BINANCE_URL","wss://stream.binance.com:9443/ws/btcusdt@ticker");
	return &Config{
		BinanceUrl: Binanceurl,
		RedisUrl: getEnv("REDIS_URL","redis://localhost:6379"),
	}
}


func getEnv(k,fallback string) string{
	v,ok:=os.LookupEnv(k);
	if(!ok){
		return fallback;
	}
	return v
}

