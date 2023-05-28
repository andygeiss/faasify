const token="3Cz7K8sR/wlj7Xv+kGxEzz0J2GOgPrmhPn2T30LrTh8=",bind=(e,t,n)=>{document.querySelector(e).addEventListener(t,n)},call=(e,t)=>{fetch("/function/"+e,config("POST",t)).then(e=>e.json()).then(t=>emit(e+" done",t)).catch(t=>emit(e+" error",t))},config=(e,t)=>({body:JSON.stringify(t),headers:{Accept:"application/json",Authorization:"Bearer "+token,"Content-Type":"application/json"},method:e}),emit=(e,t)=>{window.dispatchEvent(new CustomEvent(e,{detail:{output:t}}))},get=async(e,t)=>{const n=await fetch(e,config("GET",t)),s=await n.json();return s},html=(e,t)=>{document.querySelector(e).innerHTML=t},on=(e,t)=>{window.addEventListener(e,e=>{t(e.detail.output)})},ele=document.createElement("div");ele.id="stats",document.querySelector("body").appendChild(ele);const get_fn_names=e=>{let t=[];const n=Object.keys(e)[0];return Object.keys(e[n]).forEach(e=>{t.push(e)}),t},get_fn_stats_as_html=e=>{let t=``;return get_fn_names(e).forEach((n)=>{const o=e.active_count[n],i=e.last_response_time_ms[n],a=e.total_count[n];t+=`
<span class="fn-name">`+n+`</span>
<div class="fn-stats">
    <div> active count </div> <div> `+o+` </div>
    <div> last response time (ms) </div> <div> `+i+` </div>
    <div> total count </div> <div> `+a+` </div>
</div>
`}),t};on("status done",e=>{get("/stats").then(e=>{html("#stats",get_fn_stats_as_html(e))})}),on("status error",e=>{});const render=()=>{html("#app",` 
    <div>
        <button id="btnStatus"> call status </button>
    </div>
    `),bind("#btnStatus","click",()=>{call("status")})};on("status done",e=>{console.log(e),render()}),on("status error",e=>{console.error(e),render()}),render()