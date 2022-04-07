globalThis.crypto = require('crypto').webcrypto;
require("../go/out/wasm_exec.js");
const wasm = require("./wasm");

import express from "express";
import { info, start, move, end } from "./logic";

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

wasm.run().then(() => {
    console.log('Wasm loaded')
    // @ts-expect-error Typings soon:tm:
    console.log('Adding in Go', go_ADD_STUFF(4, 5));
    // Start the Express server
    app.listen(port, () => {
        console.log(`Starting Battlesnake Server at http://0.0.0.0:${port}...`)
    })
})
