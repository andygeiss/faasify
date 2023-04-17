
const ele = document.createElement('div');
ele.id = 'stats';
document.querySelector('body').appendChild(ele);

const get_fn_names = (obj) => {
    let names = [];
    const key = Object.keys(obj)[0];
    Object.keys(obj[key]).forEach( (name) => { 
        names.push(name) 
    });
    return names;
}

const get_fn_stats_as_html = (obj) => {
    let out = ``
    get_fn_names(obj).forEach( (name, i) => {
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
`
    })
    return out;
}

on('status done', (data) => { 
    get('/stats').then((obj) => { 
        html('#stats', get_fn_stats_as_html(obj))
    })
})

on('status error', (err) => {

})

