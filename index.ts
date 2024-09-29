import { config } from "dotenv";
import { OpenAI } from "openai";
import { Elysia } from "elysia";
import { cors } from "@elysiajs/cors";

config(); // Load .env file

async function Bot(query: any) {
  const ai = new OpenAI({ apiKey: process.env.API_KEY });
  const res = await ai.chat.completions.create({
    model: "gpt-3.5-turbo",
    messages: [{ role: "user", content: query }],
  });
  return res.choices[0].message.content;
}

const app = new Elysia();

app
  .use(cors())
  .get("/:query", async ({ params: { query } }: any) => {
    const text = query.replace("-", " ");
    const res = await Bot(text);
    return res;
  })
  .listen(5000);
