
bind("#btn-secure", "click", () => {
    call("secure", {});
});

on("secure done", (data) => {
    $("#div-secure").innerHTML = JSON.stringify(data);
});

