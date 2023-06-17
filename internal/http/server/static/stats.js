
const render_stats = (obj) => {
    let names = [];
    const key = Object.keys(obj)[0];
     Object.keys(obj[key]).forEach( (name) => { 
        names.push(name);
    });
    let out = ``;
    names.forEach( (name, i) => {
        const active_count = obj["active_count"][name];
        const last_response_time_ms = obj["last_response_time_ms"][name];
        const total_count = obj["total_count"][name];
        out += `
<span class="fn-name">` + name + `</span>
<div class="fn-stats">
    <div> active count </div> <div> `+ active_count +` </div>
    <div> last response time (ms) </div> <div> `+ last_response_time_ms +` </div>
    <div> total count </div> <div> ` + total_count + ` </div>
</div>
`;
    })
    html('#stats', out);
}

on('status done', () => { 
    get('/stats').then((obj) => {
        render_stats(obj);
    })
})

on('status error', (err) => {

})

