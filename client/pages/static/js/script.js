let isRecording = false;
let recognition;

function record() {
  let txtarea = document.getElementById("voice-input");
  const recordBtn = document.getElementById("record");
  window.SpeechRecognition =
    window.SpeechRecognition || window.webkitSpeechRecognition;

  if (!isRecording) {
    recognition = new SpeechRecognition();
    recognition.interimResults = true;
    recognition.lang = "en-US";
    recognition.continuous = true;
    recognition.start();
    isRecording = true;
    recordBtn.innerHTML = "Stop";

    recognition.addEventListener("result", (e) => {
      const transcript = Array.from(e.results)
        .map((result) => result[0])
        .map((result) => result.transcript)
        .join("");
      console.log(transcript);
      txtarea.value = transcript;
    });

    recognition.addEventListener("end", () => {
      if (isRecording) {
        recognition.start();
      }
    });
  } else {
    recognition.stop();
    isRecording = false;
    recordBtn.innerHTML = "Record";
  }
}

function edit() {
  let txtarea = document.getElementById("voice-input");
  txtarea.disabled = !txtarea.disabled;
}

function send() {
  const data = document.getElementById("voice-input").value;
  const errorMsg = document.getElementById("error-msg");
  const gptResult = document.getElementById("gpt-result");
  gptResult.innerText = "Loading...";
  fetch("/api/", {
    method: "POST",
    headers: {
      "Content-Type": "text/plain",
    },
    body: data,
  })
    .then((res) => {
      return res.json();
    })
    .then((data) => {
      console.log(data);
      if (data.error) {
        gptResult.innerText = "Result will be displayed here";
        errorMsg.querySelector("p").innerText = data.error;
        errorMsg.showModal();
        return;
      }
      gptResult.innerText = data.message;
    })
    .catch((err) => {
      console.error(err.json().error);
      gptResult.innerText = "Result will be displayed here";
      errorMsg.innerHTML += "\nError: " + err;
      errorMsg.showModal();
    });
}
