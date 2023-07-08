import { Configuration, OpenAIApi } from "openai";
import "dotenv/config";

const configuration = new Configuration({
  apiKey: process.env.OPENAI_API_KEY,
});

async function testApi() {
  const openai = new OpenAIApi(configuration);
  const chat_completion = await openai.createChatCompletion({
    model: "gpt-3.5-turbo",
    messages: [{ role: "user", content: "What is the capital of germany?" }],
  });

  console.log(chat_completion.data);
  console.log(chat_completion.data.choices[0].message?.content);
}

testApi();
