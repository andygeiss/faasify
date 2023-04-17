
const render = () => {
    html('#app', ` 
    <div>
        <button id="btnStatus"> call status </button>
    </div>
    `)
    bind('#btnStatus', 'click', () => { 
        call('status') 
    })
}

on('status done', (data) => {
    console.log(data)
    render()
})

on('status error', (data) => {
    console.error(data)
    render()
})

render()
