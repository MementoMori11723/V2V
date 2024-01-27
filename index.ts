import {config} from "dotenv"
import {OpenAI} from "openai"
import {Elysia} from "elysia"
config()
async function Bot() {
    const ai = new OpenAI({apiKey:process.env.API_KEY})
    const res = await ai.chat.completions.create({
        model:"gpt-3.5-turbo",
        messages:[{role:"user",content:"What is the capital of india"}]
    })
    return res.choices[0].message.content
}
const app = new Elysia()
    .get("/",() => "Hello World!")
    .listen(5000)