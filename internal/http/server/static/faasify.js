const $ = (sel) => {
    return document.querySelector(sel);
};

const bind = (sel, evt, fn) => {
    $(sel).addEventListener(evt, fn);
};

const call = (name, data) => {
    let cfg = { 
        body: JSON.stringify(data),
        headers: config.headers,
        method: config.method
    };
    fetch(name, cfg)
    .then( (res) => res.json() )
    .then( (data) => { emit(name + ' done', data) })
    .catch( (err) => { emit(name + ' error', err) });
};

const config = { headers: { 'accept-encoding': 'gzip, deflate' }, method: 'POST' };

const emit = (evt, data) => {
    window.dispatchEvent(new CustomEvent(evt, { detail: { output: data } }));
};

const on = (evt, fn) => {
    window.addEventListener(evt, (e) => { fn(e.detail.output) });
};

const attr = (ele) => {
    let values = {};
    Array.from(ele.attributes).forEach((n) => { values[n.nodeName] = n.nodeValue });
    return values;
}

/* faasify components */

class FaasifyBar extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['label', 'percent', 'type'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        if (kv.type == 'horizontal') {
            this.innerHTML  = `<div class="f-bar-h" style="--percent:100;--size:${kv.size}"></div>`;
            this.innerHTML += `<div class="f-bar-h f-bar-p" style="--percent:${kv.percent};--size:${kv.size}"></div>`;
        } else {
            this.innerHTML  = `<div class="f-bar-v" style="--percent:100;--size:${kv.size}"></div>`;
            this.innerHTML += `<div class="f-bar-v f-bar-p" style="--percent:${kv.percent};--size:${kv.size}"></div>`;
        }
    }
}

class FaasifyButton extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['icon', 'label'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        if ('icon' in kv) {
            this.innerHTML = `<span class="material-symbols-outlined">${kv.icon}</span>`;
        }
        if ('label' in kv) {
            this.innerHTML += `<span>${kv.label}</span>`;
        }
    }
}

class FaasifyCard extends HTMLElement {
    constructor() {
        super();
        this._innerHTML = this.innerHTML;
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['cols', 'rows'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let style = '';
        if ('cols' in kv) {
            style += `grid-template-columns:${kv.cols};`;
        }
        if ('rows' in kv) {
            style += `grid-template-rows:${kv.rows};`;
        }
        this.innerHTML = `<div class="f-card-inner" style="${style}">${this._innerHTML}</div>`;
    }
}

class FaasifyNumber extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['icon', 'value'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let icon = '';
        if ('icon' in kv) {
            icon = `<span class="material-symbols-outlined">${kv.icon}</span>`;
        }
        this.innerHTML = `${icon}<span class="f-number">${kv.value}</span>`;
    }
}

class FaasifyTable extends HTMLElement {
    constructor() {
        super();
        this._innerHTML = this.innerHTML;
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['cols', 'id', 'rows'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let style = '';
        if ('cols' in kv) {
            style += `grid-template-columns:${kv.cols};`;
        }
        if ('rows' in kv) {
            style += `grid-template-rows:${kv.rows};`;
        }
        this.innerHTML = `<div class="f-table-inner" style="${style}">${this._innerHTML}</div>`;
    }
}

class FaasifyText extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['icon', 'value'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let icon = '';
        if ('icon' in kv) {
            icon = `<span class="material-symbols-outlined">${kv.icon}</span>`;
        }
        this.innerHTML = `${icon}<span class="f-text">${kv.value}</span>`;
    }
}

customElements.define('f-bar', FaasifyBar);
customElements.define('f-button', FaasifyButton);
customElements.define('f-card', FaasifyCard);
customElements.define('f-number', FaasifyNumber);
customElements.define('f-table', FaasifyTable);
customElements.define('f-text', FaasifyText);

