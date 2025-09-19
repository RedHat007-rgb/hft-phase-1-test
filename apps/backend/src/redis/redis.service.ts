import { Injectable, OnModuleDestroy, OnModuleInit } from '@nestjs/common';
import redisClient from '@repo/redis-client';
import { RedisClientType } from 'redis';

@Injectable()
export class RedisService implements OnModuleInit, OnModuleDestroy {
  private client: RedisClientType;
  onModuleInit() {
    this.client = redisClient;
  }

  async onModuleDestroy() {
    await this.client.quit();
  }

  async subscribeMessage(channel: string, callback: (msg: string) => void) {
    await this.client.subscribe(channel, callback);
  }

  async redisSubscribe(channel: string, callback: (msg) => void) {
    await this.client.subscribe(channel, callback);
  }
}
