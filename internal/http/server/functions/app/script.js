
bind("#btn-secure", "click", () => {
    call("secure", {})
});

on("secure done", (data) => {
    query("#div-secure").innerHTML = JSON.stringify(data)
});

