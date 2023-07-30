import { Injectable } from '@nestjs/common';
import { RoundtripQueryDto } from './dto/roundtrip-query.dto';

@Injectable()
export class OpenaiService {
  create(createOpenaiDto: RoundtripQueryDto) {
    return 'This action adds a new openai';
  }
}
