const $=e=>document.querySelector(e),bind=(e,t,n)=>{$(e).addEventListener(t,n)},call=(e,t)=>{let n={body:JSON.stringify(t),headers:config.headers,method:config.method};fetch(e,n).then(e=>e.json()).then(t=>{emit(e+" done",t)}).catch(t=>{emit(e+" error",t)})},config={headers:{"accept-encoding":"gzip, deflate","content-type":"application/json"},method:"POST"},emit=(e,t)=>{window.dispatchEvent(new CustomEvent(e,{detail:{output:t}}))},on=(e,t)=>{window.addEventListener(e,e=>{t(e.detail.output)})}