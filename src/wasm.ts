globalThis.crypto = require('crypto').webcrypto;
require("../go/out/wasm_exec.js");

import { readFileSync } from "fs";

async function run() {
    // @ts-expect-error Go is imported through the require statement.
    const go = new Go();
    const mod = await WebAssembly.compile(readFileSync('go/out/main.wasm'));
    const inst = await WebAssembly.instantiate(mod, go.importObject);
    console.log('go about to run')
    go.run(inst);
    await new Promise(r => setTimeout(r, 2000));
}

export { run };
