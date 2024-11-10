# Echo Flow (V2V)

Echo Flow (also known as V2V) is an advanced voice-to-voice application developed in Go, leveraging Google’s Text-to-Speech API and OpenAI’s ChatGPT API for interactive, conversational experiences. Designed for accessibility and ease of use, Echo Flow allows users to have two-way audio interactions with AI.

## Features

- **Voice-to-Voice Conversations**: Speak your questions and get responses in real-time.
- **Powered by Google TTS**: Clear, natural-sounding speech synthesis for seamless conversations.
- **Enhanced AI Responses**: Uses ChatGPT for accurate and engaging conversational replies.
- **Dark Mode Support**: Toggle between light and dark themes.
- **Accessible UI**: Clean, mobile-friendly interface with large buttons for easy navigation.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/username/echo-flow.git
   cd echo-flow
   ```

2. **Set Up Environment Variables**:
   Add the following to a `.env` file or export them in your terminal:
   ```env
   TTS_API_KEY=<your_google_tts_api_key>
   GPT_API_KEY=<your_openai_api_key>
   ```
   or 

   ```bash
   export TTS_API_KEY=<your_google_tts_api_key>
   export GPT_API_KEY=<your_openai_api_key>
   ```
    you can also add port number to the env file or export it in the terminal
    ```env
    PORT=8080
    ```
    or
    ```bash
    export PORT=8080
    ```
    
3. **Run the Application**:
   ```bash
   make
   ```

## Usage

- **Record & Send**: Press the record button to speak your message. Upon release, Echo Flow will process your speech input and send it to ChatGPT.
- **Receive Response**: Listen to the AI-generated response using Google’s TTS. You can also toggle options like edit mode and dark mode.

## Dependencies

- **Go**: Latest stable version.
- **Google TTS API**: For text-to-speech functionalities.
- **OpenAI GPT API**: For conversational AI responses.

## License

MIT License.
