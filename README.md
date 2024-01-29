# V2V

A voice to voice bot that gets responce from chatgpt.

## Requirments

- Bun or node.js
- Dotenv
- Open AI
- Elysia and ELysiajs/core

## Code

First we need to create an api using Bun and elysia

``` typescript
import {config} from "dotenv"
import {OpenAI} from "openai"
import {Elysia} from "elysia"
import {cors} from "@elysiajs/cors"
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
    .use(cors())
    .get("/:id",async ({params:{id}}:any) => {
        const text = id.replace("-"," ")
        const res = await Bot(text)
        return res
    })
    .listen(5000)
```

Next we need to create a html file which will be our main code

``` html
    <!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>V2V</title>
    <script defer>
        var speech = true; 
        window.SpeechRecognition = window.SpeechRecognition || window.webkitSpeechRecognition; 
        const recognition = new SpeechRecognition(); 
        recognition.interimResults = true; 
        const words = document.querySelector('.words');
        recognition.addEventListener('result', e => { 
        const transcript = Array.from(e.results) 
                .map(result => result[0]) 
                .map(result => result.transcript) 
                .join('') 
        document.getElementById("p").innerHTML = transcript;
        });
        if (speech == true) { 
            recognition.start(); 
            recognition.addEventListener('end', recognition.start); 
        }
        const clicked = async () => {
            txtbox = document.getElementById("p")
            const text = txtbox.innerHTML
            txtbox.innerHTML = await convert(text)
        }

        const convert = async(text) => {
            let result
            let url = text.replaceAll(" ","-")
            console.log(url)
            await fetch(`http://localhost:5000/${url}-under-3-lines`)
            .then(res => res.text())
            .then(data => {
                result = data
            })
            console.log(result)
            const uttrence = new SpeechSynthesisUtterance(result)
            uttrence.rate = 1
            console.log(uttrence)
            speechSynthesis.speak(uttrence)
            return result
        }
    </script>
</head>
<body>
    <div style="text-align: center;">
        <h1>Voice to Voice</h1>
        <div class="words" contenteditable>
            <p id="p"></p>
        </div>
        <button onclick="clicked()">Send</button>
    </div>
</body>
</html>
```

And there you have it, you have successfully created Voice to Voice Bot.

> Note : There are a few issues with speechSynthesis (it may or may not work properly).
