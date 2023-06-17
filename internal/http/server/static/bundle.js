const token="yvpmta89As0nN3AF+mcrYVyqJtzY61W/rTzzhilL0G8=",bind=(e,t,n)=>{document.querySelector(e).addEventListener(t,n)},call=(e,t)=>{fetch("/function/"+e,config("POST",t)).then(e=>e.json()).then(t=>emit(e+" done",t)).catch(t=>emit(e+" error",t))},config=(e,t)=>({body:JSON.stringify(t),headers:{Accept:"application/json",Authorization:"Bearer "+token,"Content-Type":"application/json"},method:e}),emit=(e,t)=>{window.dispatchEvent(new CustomEvent(e,{detail:{output:t}}))},get=async(e,t)=>{const n=await fetch(e,config("GET",t)),s=await n.json();return s},html=(e,t)=>{document.querySelector(e).innerHTML=t},on=(e,t)=>{window.addEventListener(e,e=>{t(e.detail.output)})},render_app=()=>{html("#app",` 
    <div>
        <button id="btnStatus"> Status </button>
    </div>
    `),bind("#btnStatus","click",()=>{call("status")})};on("status done",()=>{render_app()}),on("status error",()=>{render_app()}),render_app();const render_stats=e=>{let t=[];const s=Object.keys(e)[0];Object.keys(e[s]).forEach(e=>{t.push(e)});let n=``;t.forEach((t)=>{const o=e.active_count[t],i=e.last_response_time_ms[t],a=e.total_count[t];n+=`
<span class="fn-name">`+t+`</span>
<div class="fn-stats">
    <div> active count </div> <div> `+o+` </div>
    <div> last response time (ms) </div> <div> `+i+` </div>
    <div> total count </div> <div> `+a+` </div>
</div>
`}),html("#stats",n)};on("status done",()=>{get("/stats").then(e=>{render_stats(e)})}),on("status error",e=>{})