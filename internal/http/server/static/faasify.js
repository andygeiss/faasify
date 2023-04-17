// Code generated by faasify DO NOT EDIT
const token = "RvOYJFyqSIfZBz6wHqg5pDwyAWrkyBK8xO3niPhFOi0="

const bind = (element, event, fn) => {
    document.querySelector(element).addEventListener(event, fn)
}

const call = (name, data) => {
    fetch("/function/"+name, config('POST', data))
    .then((res) => res.json())
    .then((data) => emit(name + " done", data))
    .catch((err) => emit(name + " error", err))
}

const config = (method, data) => {
    return {
        body: JSON.stringify(data),
        headers: {
            'Accept': 'application/json', 
            'Authorization': 'Bearer ' + token,
            'Content-Type': 'application/json'
        },
        method: method
    }
}

const emit = (event, data) => {
    window.dispatchEvent(new CustomEvent(event, { detail: { output: data } }))
}

const get = async (path, data) => {
    const res = await fetch(path, config('GET', data));
    const obj = await res.json();
    return obj;
}

const html = (element, content) => {
    document.querySelector(element).innerHTML = content
}

const on = (event, fn) => {
    window.addEventListener(event, (evt) => { fn(evt.detail.output) }) 
}

