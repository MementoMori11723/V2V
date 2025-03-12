"serviceWorker" in navigator &&
  window.addEventListener("load", function () {
    navigator.serviceWorker
      .register("/assets/js/worker.js")
      .then(function (e) {
        console.log("ServiceWorker registration successful");
      })
      .catch(function (e) {
        console.error("ServiceWorker registration failed:", e);
      });
  });
let isRecording = !1,
  recognition;
function record() {
  let e = document.getElementById("voice-input"),
    n = document.getElementById("record");
  (window.SpeechRecognition =
    window.SpeechRecognition || window.webkitSpeechRecognition),
    isRecording
      ? (recognition.stop(),
        (isRecording = !1),
        (n.innerHTML = '<i class="fas fa-microphone"></i> Record'))
      : (((recognition = new SpeechRecognition()).interimResults = !0),
        (recognition.lang = "en-US"),
        (recognition.continuous = !0),
        recognition.start(),
        (isRecording = !0),
        (n.innerHTML = '<i class="fas fa-stop-circle"></i> Stop'),
        recognition.addEventListener("result", (n) => {
          let i = Array.from(n.results)
            .map((e) => e[0])
            .map((e) => e.transcript)
            .join("");
          e.value = i;
        }),
        recognition.addEventListener("end", () => {
          isRecording && recognition.start();
        }));
}
function edit() {
  let e = document.getElementById("voice-input");
  (e.disabled = !e.disabled), e.classList.toggle("editable");
}
function send() {
  let e = document.getElementById("voice-input").value,
    n = document.getElementById("gpt-result"),
    i = document.getElementById("error-msg"),
    t = document.getElementById("audio"),
    r = document.getElementById("send");
  (r.disabled = !0),
    (n.innerText = "Loading..."),
    fetch("/api", {
      method: "POST",
      headers: { "Content-Type": "text/plain" },
      body: e,
    })
      .then((e) => e.json())
      .then((e) => {
        if (e.error) {
          (n.innerText = "Result will be displayed here"),
            (i.querySelector("p").innerText = e.error),
            (i.style.display = "block");
          return;
        }
        (n.innerText = html), (t.src = e.audioMessage);
      })
      .catch((e) => {
        (n.innerText = "Result will be displayed here"),
          (i.querySelector("p").innerText = `Error: ${e}`),
          (i.style.display = "block");
      })
      .finally(() => {
        r.disabled = !1;
      });
}
function closeErrorDialog() {
  document.getElementById("error-msg").style.display = "none";
}
function speak() {
  let e = document.getElementById("audio"),
    n = document.getElementById("speak"),
    i = document.getElementById("error-msg");
  if (!e.src) {
    (i.querySelector("p").innerText = "No audio available to play."),
      (i.style.display = "block");
    return;
  }
  e.paused
    ? (e.play(),
      (n.innerHTML = '<i class="fas fa-stop"></i> Stop'),
      (e.onended = () => {
        n.innerHTML = '<i class="fas fa-volume-up"></i> Speak';
      }))
    : (e.pause(), (n.innerHTML = '<i class="fas fa-volume-up"></i> Speak'));
}
