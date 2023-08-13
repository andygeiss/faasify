document.addEventListener("fetch", (event) => {
    if (event.request.cache === "only-if-cached" && event.request.mode !== "same-origin") return;
    if (event.request.headers.get("Accept").includes("text/html")) {
        event.respondWith( fetch(event.request).then((res) => { return res }).catch( (err) => { return caches.match("offline")}) )
    }
});
