const bind = (sel, evt, fn) => {
    query(sel).addEventListener(evt, fn)
};

const call = (name, data) => {
    let cfg = { 
        body: JSON.stringify(data),
        headers: config.headers,
        method: config.method
    };
    fetch(name, cfg)
    .then( (res) => res.json() )
    .then( (data) => { emit(name + " done", data) })
    .catch( (err) => { emit(name + " error", err) })
};

const config = {
    headers: {
        "accept-encoding": "gzip, deflate",
        "content-type": "application/json"
    },
    method: "POST"
};

const emit = (evt, data) => {
    window.dispatchEvent(new CustomEvent(evt, { detail: { output: data } })) 
};

const on = (evt, fn) => {
    window.addEventListener(evt, (e) => { fn(e.detail.output) })
};

const query = (sel) => {
    return document.querySelector(sel)
};
