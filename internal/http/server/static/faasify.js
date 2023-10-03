/*
  ___ ___  _ __ ___  _ __ ___   ___  _ __
 / __/ _ \| '_ ` _ \| '_ ` _ \ / _ \| '_ \
| (_| (_) | | | | | | | | | | | (_) | | | |
 \___\___/|_| |_| |_|_| |_| |_|\___/|_| |_|
*/
const $ = (sel) => {
    return document.querySelector(sel);
};

const attr = (ele) => {
    let values = {};
    Array.from(ele.attributes).forEach((n) => { values[n.nodeName] = n.nodeValue });
    return values;
}

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

const switchClosed = (sel) => {
    let kv = attr($(sel));
    $(sel).removeAttribute('open');
    $(sel).setAttribute('closed', '');
};

const switchOpen = (sel) => {
    let kv = attr($(sel));
    $(sel).removeAttribute('closed');
    $(sel).setAttribute('open', '');
};

const toggle = (sel) => {
    let kv = attr($(sel));
    if ('open' in kv) {
        switchClosed(sel);
        return;
    }
    switchOpen(sel);
};

/*
  __       _
 / _|     | |__   __ _ _ __
| |_ _____| '_ \ / _` | '__|
|  _|_____| |_) | (_| | |
|_|       |_.__/ \__,_|_|
*/
class FaasifyBar extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['color', 'percent', 'type'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let color = '';
        let style_inner = '';
        let style_outer = '';
        if ('color' in kv) {
            color = `background:${kv.color}`;
        }
        if ('percent' in kv) {
            style_inner = `--percent:${kv.percent};--size:${kv.size};${color}`;
        }
        if ('size' in kv) {
            style_outer = `--percent:100;--size:${kv.size}`;
        }
        if (kv.type == 'horizontal') {
            this.innerHTML  = `<div class="f-bar-h" style="${style_outer}"></div>`;
            this.innerHTML += `<div class="f-bar-h f-bar-p" style="${style_inner}"></div>`;
        } else {
            this.innerHTML  = `<div class="f-bar-v" style="${style_outer}"></div>`;
            this.innerHTML += `<div class="f-bar-v f-bar-p" style="${style_inner}"></div>`;
        }
    }
}

/*
  __       _                          _                _
 / _|     | |__   __ _ _ __       ___| |__   __ _ _ __| |_
| |_ _____| '_ \ / _` | '__|____ / __| '_ \ / _` | '__| __|
|  _|_____| |_) | (_| | | |_____| (__| | | | (_| | |  | |_
|_|       |_.__/ \__,_|_|        \___|_| |_|\__,_|_|   \__|
*/
class FaasifyBarChart extends HTMLElement {
    constructor() {
        super();
        this._innerHTML = this.innerHTML;
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['colors', 'size', 'values'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let bars = '';
        let colors = [];
        let values = [];
        kv.colors.split(',').forEach( (color) => {
            colors.push(color);
        })
        kv.values.split(',').forEach( (percent) => {
            values.push(percent);
        });
        values.forEach( (percent, index) => {
            let color = 'var(--color-accent-light)'
            if (colors[index] === '1') {
                color = 'var(--color-accent)'
            }
            bars += `<f-bar color="${color}" percent="${percent}" size="${kv.size}"></f-bar>`;
        })
        this.innerHTML = `<f-table cols="repeat(${values.length}, 1fr)">${bars}</f-table>`;
    }
}

/*
  __       _           _   _
 / _|     | |__  _   _| |_| |_ ___  _ __
| |_ _____| '_ \| | | | __| __/ _ \| '_ \
|  _|_____| |_) | |_| | |_| || (_) | | | |
|_|       |_.__/ \__,_|\__|\__\___/|_| |_|
*/
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

/*
  __                         _
 / _|       ___ __ _ _ __ __| |
| |_ _____ / __/ _` | '__/ _` |
|  _|_____| (_| (_| | | | (_| |
|_|        \___\__,_|_|  \__,_|
*/
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

/*
  __
 / _|       __ _  __ _ _   _  __ _  ___
| |_ _____ / _` |/ _` | | | |/ _` |/ _ \
|  _|_____| (_| | (_| | |_| | (_| |  __/
|_|        \__, |\__,_|\__,_|\__, |\___|
           |___/             |___/
*/
class FaasifyGauge extends HTMLElement {
    constructor() {
        super();
        this._innerHTML = this.innerHTML;
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['color', 'label', 'percent', 'radius'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let color = '';
        let label = '';
        let percent = '0';
        let radius = 40;
        let c = radius+10;
        let style = '';
        if ('color' in kv) {
            color = ` style="stroke:${kv.color}"`;
        }
        if ('label' in kv) {
            label = `<span>${kv.label}</span>`;
        }
        if ('percent' in kv) {
            percent = kv.percent;
        }
        if ('radius' in kv) {
            radius = kv.radius;
            c = eval(parseInt(radius)+10);
        }
        style = `--percent:${percent};--radius:${radius};--dasharray:calc(6.275px * var(--radius));`;
        this.innerHTML = label +
            `<svg style="${style}">` +
            `<circle r="${radius}" cx="${c}" cy="${c}"></circle>` +
            `<circle class="f-gauge-percent" r="${radius}" cx="${c}" cy="${c}"${color}></circle>` +
            `</svg>`;
    }
}

/*
  __                                               _                _
 / _|       __ _  __ _ _   _  __ _  ___        ___| |__   __ _ _ __| |_
| |_ _____ / _` |/ _` | | | |/ _` |/ _ \_____ / __| '_ \ / _` | '__| __|
|  _|_____| (_| | (_| | |_| | (_| |  __/_____| (__| | | | (_| | |  | |_
|_|        \__, |\__,_|\__,_|\__, |\___|      \___|_| |_|\__,_|_|   \__|
           |___/             |___/
*/
class FaasifyGaugeChart extends HTMLElement {
    constructor() {
        super();
        this._innerHTML = this.innerHTML;
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['colors', 'values'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let gauges = '';
        let colors = [];
        let values = [];
        kv.colors.split(',').forEach( (color) => {
            colors.push(color);
        })
        kv.values.split(',').forEach( (percent) => {
            values.push(percent);
        });
        values.forEach( (percent, index) => {
            let color = 'var(--color-accent-light)'
            if (colors[index] === '1') {
                color = 'var(--color-accent)'
            }
            let i = eval(index+1);
            let radius = eval(100 - (index*11));
            gauges += `<f-gauge class="f-gauge-chart-${i}" color="${color}" percent="${percent}" radius="${radius}"></f-gauge>`;
        })
        this.innerHTML = gauges;
    }
}

/*
  __       _                   _                                                  _
 / _|     (_)_ __  _ __  _   _| |_      _ __   __ _ ___ _____      _____  _ __ __| |
| |_ _____| | '_ \| '_ \| | | | __|____| '_ \ / _` / __/ __\ \ /\ / / _ \| '__/ _` |
|  _|_____| | | | | |_) | |_| | ||_____| |_) | (_| \__ \__ \\ V  V / (_) | | | (_| |
|_|       |_|_| |_| .__/ \__,_|\__|    | .__/ \__,_|___/___/ \_/\_/ \___/|_|  \__,_|
                  |_|                  |_|
*/
class FaasifyInputPassword extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['id', 'label'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let label = '';
        let iid = `input-${kv.id}`;
        let icon = `<span class="material-symbols-outlined">key</span>`;
        if ('label' in kv) {
            label = `<label for="${iid}">${kv.label}</label>`
        }
        let input = `<input id="${iid}" class="f-input-text" type="password"></input>`;
        this.innerHTML = `${label}${icon}${input}`;
    }
}

/*
  __       _                   _                                _
 / _|     (_)_ __  _ __  _   _| |_      ___  ___  __ _ _ __ ___| |__
| |_ _____| | '_ \| '_ \| | | | __|____/ __|/ _ \/ _` | '__/ __| '_ \
|  _|_____| | | | | |_) | |_| | ||_____\__ \  __/ (_| | | | (__| | | |
|_|       |_|_| |_| .__/ \__,_|\__|    |___/\___|\__,_|_|  \___|_| |_|
                  |_|
*/
class FaasifyInputSearch extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['id', 'label'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let label = '';
        let iid = `input-${kv.id}`;
        if ('label' in kv) {
            label = `<label for="${iid}">${kv.label}</label>`
        }
        let icon = `<span class="material-symbols-outlined">search</span>`;
        let search = `<input id="${iid}" class="f-input-search" type="search"></input>`;
        this.innerHTML = `${label}${icon}${search}`;
    }
}

/*
  __       _                   _        _            _
 / _|     (_)_ __  _ __  _   _| |_     | |_ _____  _| |_
| |_ _____| | '_ \| '_ \| | | | __|____| __/ _ \ \/ / __|
|  _|_____| | | | | |_) | |_| | ||_____| ||  __/>  <| |_
|_|       |_|_| |_| .__/ \__,_|\__|     \__\___/_/\_\\__|
                  |_|
*/
class FaasifyInputText extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['disabled', 'icon', 'id', 'label', 'value'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let disabled = '';
        let icon = '';
        let label = '';
        let value = '';
        let iid = `input-${kv.id}`;
        if ('disabled' in kv) {
            disabled = ` disabled`;
        }
        if ('icon' in kv) {
            icon = `<span class="material-symbols-outlined">${kv.icon}</span>`;
        }
        if ('label' in kv) {
            label = `<label for="${iid}">${kv.label}</label>`
        }
        if ('value' in kv) {
            value = kv.value
        }
        let input = `<textarea id="${iid}" class="f-input-text${disabled}"${disabled}>${value}</textarea>`;
        this.innerHTML = `${label}${icon}${input}`;
    }
}

/*
  __       _                   _        _                    _
 / _|     (_)_ __  _ __  _   _| |_     | |_ ___   __ _  __ _| | ___
| |_ _____| | '_ \| '_ \| | | | __|____| __/ _ \ / _` |/ _` | |/ _ \
|  _|_____| | | | | |_) | |_| | ||_____| || (_) | (_| | (_| | |  __/
|_|       |_|_| |_| .__/ \__,_|\__|     \__\___/ \__, |\__, |_|\___|
                  |_|                            |___/ |___/ 
*/
class FaasifyInputToggle extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['id', 'label'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let label = '';
        let iid = `input-${kv.id}`;
        this.innerHTML = `
<label class="f-input-toggle-inner">
    <div class="f-input-toggle">
        <input class="f-input-toggle-state" id="${iid}" name="${iid}" type="checkbox" value="check"/>
        <div class="f-input-toggle-indicator"></div>
    </div>
    <div class="f-input-toggle-label">${kv.label}</div>
</label>`;
    }
}

/*
  __                             _
 / _|      _ __  _   _ _ __ ___ | |__   ___ _ __
| |_ _____| '_ \| | | | '_ ` _ \| '_ \ / _ \ '__|
|  _|_____| | | | |_| | | | | | | |_) |  __/ |
|_|       |_| |_|\__,_|_| |_| |_|_.__/ \___|_|
*/
class FaasifyNumber extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['color', 'icon', 'value'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let icon = '';
        let style = '';
        if ('color' in kv) {
            style = ` style="color:${kv.color}"`;
        }
        if ('icon' in kv) {
            icon = `<span class="material-symbols-outlined">${kv.icon}</span>`;
        }
        this.innerHTML = `${icon}<span class="f-number"${style}>${kv.value}</span>`;
    }
}

/*
  __       _        _     _
 / _|     | |_ __ _| |__ | | ___
| |_ _____| __/ _` | '_ \| |/ _ \
|  _|_____| || (_| | |_) | |  __/
|_|        \__\__,_|_.__/|_|\___|
*/
class FaasifyTable extends HTMLElement {
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
        this.innerHTML = `<div class="f-table-inner" style="${style}">${this._innerHTML}</div>`;
    }
}

/*
  __       _            _
 / _|     | |_ _____  _| |_
| |_ _____| __/ _ \ \/ / __|
|  _|_____| ||  __/>  <| |_
|_|        \__\___/_/\_\\__|
*/
class FaasifyText extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.render();
    }
    static get observedAttributes() {
        return ['color', 'icon', 'value'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        this.render();
    }
    render() {
        let kv = attr(this);
        let icon = '';
        let style = '';
        if ('color' in kv) {
            style = ` style="color:${kv.color}"`;
        }
        if ('icon' in kv) {
            icon = `<span class="material-symbols-outlined">${kv.icon}</span>`;
        }
        this.innerHTML = `${icon}<span class="f-text"${style}>${kv.value}</span>`;
    }
}

/*
  __       _   _ _   _
 / _|     | |_(_) |_| | ___
| |_ _____| __| | __| |/ _ \
|  _|_____| |_| | |_| |  __/
|_|        \__|_|\__|_|\___|
*/
class FaasifyTitle extends HTMLElement {
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
        this.innerHTML = `${icon}<span class="f-title">${kv.value}</span>`;
    }
}

/*
     _       __ _       _ _   _
  __| | ___ / _(_)_ __ (_) |_(_) ___  _ __
 / _` |/ _ \ |_| | '_ \| | __| |/ _ \| '_ \
| (_| |  __/  _| | | | | | |_| | (_) | | | |
 \__,_|\___|_| |_|_| |_|_|\__|_|\___/|_| |_|
*/
customElements.define('f-bar', FaasifyBar);
customElements.define('f-bar-chart', FaasifyBarChart);
customElements.define('f-button', FaasifyButton);
customElements.define('f-card', FaasifyCard);
customElements.define('f-gauge', FaasifyGauge);
customElements.define('f-gauge-chart', FaasifyGaugeChart);
customElements.define('f-input-password', FaasifyInputPassword);
customElements.define('f-input-search', FaasifyInputSearch);
customElements.define('f-input-text', FaasifyInputText);
customElements.define('f-input-toggle', FaasifyInputToggle);
customElements.define('f-number', FaasifyNumber);
customElements.define('f-table', FaasifyTable);
customElements.define('f-text', FaasifyText);
customElements.define('f-title', FaasifyTitle);

