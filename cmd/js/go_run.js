let namespace = "namespace";

window[namespace] = window[namespace] || {};
window[namespace].listener = (event) => {
	console.log(event);
}



const go = new Go(); // Defined in wasm_exec.js
const WASM_URL = './wasm/wasm.wasm';

let wasm;

if ('instantiateStreaming' in WebAssembly) {
	WebAssembly.instantiateStreaming(fetch(WASM_URL),go.importObject).then( res=> {
        go.run(res.instance)
    })
} else {
	fetch(WASM_URL).then(resp =>
		resp.arrayBuffer()
	).then(bytes =>
		WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
			wasm = obj.instance;
			go.run(wasm);
		})
	)
}
