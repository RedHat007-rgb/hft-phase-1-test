import {
  OnGatewayConnection,
  OnGatewayInit,
  WebSocketGateway,
} from '@nestjs/websockets';
import { RedisService } from './redis.service';
import { InternalServerErrorException } from '@nestjs/common';

@WebSocketGateway()
export class RedisGateway implements OnGatewayConnection, OnGatewayInit {
  constructor(private redisService: RedisService) {}
  afterInit() {
    console.log('websocket initialized...');
  }
  async handleConnection(client: WebSocket) {
    if (
      client.readyState === WebSocket.CLOSED ||
      client.readyState === WebSocket.CLOSING
    ) {
      return;
    }
    try {
      await this.redisService.subscribeMessage(
        'market.ticker.BTCUSDT.binance',
        (message) => {
          if (client.readyState === WebSocket.OPEN) {
            client.send(message);
          }
        },
      );
    } catch {
      throw new InternalServerErrorException(
        'error occured .... in redis gateway...',
      );
    }
  }
}
