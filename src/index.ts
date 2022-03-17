import express from "express";
import { readFileSync } from "fs";
import { info, start, move, end } from "./logic";
require("../go/out/wasm_exec.js");

async function run() {
    // @ts-expect-error Go is imported through the require statement.
    const go = new Go();
    const mod = await WebAssembly.compile(readFileSync('go/out/main.wasm'));
    const inst = await WebAssembly.instantiate(mod, go.importObject);
    console.log('go about to run')
    go.run(inst);
    await new Promise(r => setTimeout(r, 2000));
}

const app = express()
app.use(express.json())

const port = process.env.PORT || 8080

app.get("/", (req, res) => {
    res.send(info())
});

app.post("/start", (req, res) => {
    res.send(start(req.body))
});

app.post("/move", (req, res) => {
    const chosenMove = move(req.body);
    console.log(chosenMove.move);

    res.send(chosenMove)
});

app.post("/end", (req, res) => {
    res.send(end(req.body))
});

run().then(() => {
    console.log('Wasm loaded')
    // @ts-expect-error Typings soon:tm:
    console.log('Adding in Go', go_ADD_STUFF(4, 5));
    // Start the Express server
    app.listen(port, () => {
        console.log(`Starting Battlesnake Server at http://0.0.0.0:${port}...`)
    })
})
