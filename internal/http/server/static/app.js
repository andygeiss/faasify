
const render_app = () => {
    html('#app', ` 
    <div>
        <button id="btnStatus"> Status </button>
    </div>
    `);
    bind('#btnStatus', 'click', () => { 
        call('status'); 
    });
}

on('status done', () => {
    render_app();
})

on('status error', () => {
    render_app();
})

render_app();
