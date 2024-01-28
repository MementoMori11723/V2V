import {config} from "dotenv"
import {OpenAI} from "openai"
import {Elysia} from "elysia"
config()
async function Bot(q:any) {
    const ai = new OpenAI({apiKey:process.env.API_KEY})
    const res = await ai.chat.completions.create({
        model:"gpt-3.5-turbo",
        messages:[{role:"user",content:q}]
    })  
    return res.choices[0].message.content
}
const app = new Elysia()
    .get("/:id",async ({params:{id}}:any) => {
        const text = id.replace("-"," ")
        const res = await Bot(text)
        return res
    })
    .listen(5000)