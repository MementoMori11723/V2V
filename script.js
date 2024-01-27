// Speech to Text
window.SpeechRecognition = window.webkitSpeechRecognition
const recognition = new SpeechRecognition()

// text to speech
const uttrence = new SpeechSynthesisUtterance("Hello World")
uttrence.rate = 1
speechSynthesis.speak(uttrence)