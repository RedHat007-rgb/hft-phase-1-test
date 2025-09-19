import { Module } from '@nestjs/common';
import { RedisGateway } from './redis.gateway';
import { RedisService } from './redis.service';

@Module({
  providers: [RedisGateway, RedisService],
})
export class RedisModule {}
