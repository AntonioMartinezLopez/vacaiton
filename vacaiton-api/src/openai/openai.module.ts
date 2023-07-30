import { Module } from '@nestjs/common';
import { OpenaiService } from './openai.service';
import { OpenaiController } from './openai.controller';

// Module: basic information concerning the module
// Controllers: REST controller classes that are being used in this module+
// Providers: classes that are available in this module (imports and exports possible)

@Module({
  //imports
  controllers: [OpenaiController],
  providers: [OpenaiService],
  //exports
})
export class OpenaiModule {}
