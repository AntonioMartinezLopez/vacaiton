import { Controller, Post, Body, Logger } from '@nestjs/common';
import { OpenaiService } from './openai.service';
import { RoundtripQueryDto } from './dto/roundtrip-query.dto';

@Controller('openai')
export class OpenaiController {
  private readonly logger = new Logger(OpenaiController.name);

  constructor(private readonly openaiService: OpenaiService) {}

  @Post('createTrip') /** /api/openai/createTrip */
  // @Body: body annotation --> automatically uses body from request
  create(@Body() roundtripQueryDto: RoundtripQueryDto) {
    this.logger.debug(JSON.stringify(roundtripQueryDto));
    return this.openaiService.create(roundtripQueryDto);
  }
}
