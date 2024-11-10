const CACHE_NAME = "echo-flow-cache-v1",
  urlsToCache = [
    "/",
    "/assets/favicon.ico",
    "/assets/manifest.json",
    "/assets/close.png",
  ];
self.addEventListener("install", (e) => {
  e.waitUntil(
    caches
      .open(CACHE_NAME)
      .then(
        (e) => (
          console.log("Opened cache and caching assets"), e.addAll(urlsToCache)
        ),
      )
      .catch((e) => {
        console.error("Caching failed:", e);
      }),
  );
}),
  self.addEventListener("fetch", (e) => {
    e.respondWith(
      caches
        .match(e.request)
        .then(
          (t) =>
            t ||
            fetch(e.request).then((t) => {
              if (!t || 200 !== t.status || "basic" !== t.type) return t;
              let n = t.clone();
              return (
                caches.open(CACHE_NAME).then((t) => {
                  t.put(e.request, n);
                }),
                t
              );
            }),
        )
        .catch((e) => {
          console.error("Fetch error:", e);
        }),
    );
  }),
  self.addEventListener("activate", (e) => {
    let t = [CACHE_NAME];
    e.waitUntil(
      caches.keys().then((e) =>
        Promise.all(
          e.map((e) => {
            if (!t.includes(e))
              return console.log("Deleting old cache:", e), caches.delete(e);
          }),
        ),
      ),
    );
  });
