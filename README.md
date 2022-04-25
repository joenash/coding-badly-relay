# Coding Badly Relay
[![Deploy to DO](https://www.deploytodo.com/do-btn-blue.svg)](https://cloud.digitalocean.com/apps/new?repo=https://github.com/masonegger/coding-badly-relay/tree/main)

This is the Coding Badly Relay Snake. Each episode of Coding Badly, a new guest works on the snake, adding their own code and strategies. At the end of that episode, it's handed on, for the next person.

The Relay snake started out as the Battlesnake Javascript Starter. Where will it end up?

Check below for the schedule!

## Schedule

All shows are on the [Battlesnake Twitch](https://battlesnake.tv) at 6pm UTC. [Find your timezone](https://everytimezone.com/convert/pdt/11am).

| Date          | Guest | Final commit |
| ------------- | ----- | ------------ |
| 3rd February  |   coreyja    |       [ad5ba0d](https://github.com/joenash/coding-badly-relay/commit/ad5ba0d2076c312619089f68496bef29d484b3eb)       |
| 17th February | masonegger  |       [984aa9d](https://github.com/joenash/coding-badly-relay/commit/984aa9db4329779ddcc6ac615078c05d44da6948)       |
| 3rd March     | geeh  |       [e98e3fc](https://github.com/joenash/coding-badly-relay/commit/e98e3fc7a4b49eb8059561fdd1a5becafbfab815)       |
| 17th March    |   nhcarrigan    |        [72d2883](https://github.com/joenash/coding-badly-relay/commit/72d2883cfdcfe55bff2d37afcf6763574199925e)      |
| 7th April     |   penelope_zone    | [7f0c342](https://github.com/joenash/coding-badly-relay/commit/7f0c342abb68d5b8b2744f20398105d4fdd860e1)             |
| 21st April    |   robbles    |      [72233c9](https://github.com/joenash/coding-badly-relay/commit/72233c92fd022949137ea0e8b3e9015dcacbb72d)        |
| 5th May       |   Alec Swanson   |              |
| 19th May      |       |              |

## Joining the show

Want to be a guest in the relay? Contact [Joe](https://twitter.com/jna_sh) or [Kevin](https://twitter.com/_phzn). Each guest will be building on the snake built in the show prior. You'll have approximately two weeks warning to know what terrible nonsense the last holder of the baton has gotten up to, to work out what you're going to do about it. Shows are live, broadcast on Twitch, via Streamyard. Every guest will be accompanied by at least one (usually both) of the hosts.

Coding Badly strives to be a low-pressure, fun environment to enjoy silly code that does useless things. That isn't to say we don't welcome good code, but it is to say that if this is your first time live coding on a stream, or your first time using Javascript, or your first time seeing Battlesnake: do not worry, this is the space for you.

All guests will be expected to abide by the Battlesnake Code of Conduct, both on air and in their submitted code.

## Technologies Used (so far)

- [JavaScript](https://www.javascript.com/)
- [TypeScript](https://www.typescriptlang.org/)
- [Node.js](https://nodejs.dev/)
- [Express](https://expressjs.com/)
- [Golang](https://go.dev/)
- [WASM](https://webassembly.org/)
- [Battlesnake Rules Repo](https://github.com/BattlesnakeOfficial/rules)

## Handoffs

### coreyja

We decided to do some WASM! We compiled the Battlesnake Rules repo, written in Go, to WASM and are including it in our node snake.
We expose a Go function that can be used to check if a given set of moves will cause death for a starting board state. We return the elimination cause
and an empty string implies that the snake survived.

To compile the Go source to a WASM file there is a `go/build.sh` script. Run this from the `go` subdirectory to compile the `main.wasm` file and copy the Go Exec JS file.

```
cd go
./build.sh
```

> Adding `fmt.Println` breaks the connection between the Go and JS

### Mason Egger ([@masonegger](https://twitter.com/masonegger)) and Matt Cowley ([@mattipv4](https://twitter.com/MattIPv4))

Mason joined to help tackle the "How do we deploy this?" part of the coding relay.
After stumbling around yaml for a bit, human yaml validator Matt Cowley appeared to help.
We setup a GitHub action to build the Go WASM from the previous stream. We also setup
the Deploy to DigitalOcean button so anyone can deploy this battlesnake to [DigitalOcean App Platform](https://www.digitalocean.com/products/app-platform) with just a few clicks

**Notes for deploying** If you fork this repository, you'll need to change the URL in the Deploy to DigitalOcean Button in order for it to point to your github repository, as well as modify the `.do/deploy.template.yaml` to also point to your repository.

**Tricky Note** The github repo username/repo is CASE SENSITIVE, so ya. Don't fall down that rabbit hole.

`[![Deploy to DO](https://www.deploytodo.com/do-btn-blue.svg)](https://cloud.digitalocean.com/apps/new?repo=https://github.com/<YOUR_GITHUB_USERNAME_HERE>/coding-badly-relay/tree/main)`

Change the yaml as so in `.do/deploy.template.yaml`
```yaml
spec:
  name: coding-badly-relay-snake
  services:
  - environment_slug: node-js
    github:
      branch: main
      repo: <YOUR_GITHUB_USERNAME_HERE>/coding-badly-relay
      deploy_on_push: true
    name: coding-badly-relay-snake
    build_command: mkdir -p go/out && wget -O go/out/main.wasm https://github.com/joenash/coding-badly-relay/releases/download/latest/main.wasm && wget -O go/out/wasm_exec.js https://github.com/joenash/coding-badly-relay/releases/download/latest/wasm_exec.js
```

If you decide to modify the WASM files, you'll need to update the build command (replace joenash with your github username) to grab your files, as well as manually trigger the GitHub Action to build your own WASM files.

## nhcarrigan

We converted the project to use TypeScript. Aside from building the Go files, you'll now need to build the TypeScript files. The `package.json` is updated to handle the TS - run `npm run build` to compile the TS files, and `npm start` to run your server.

**NOTE: Currently the tests are not functional as the Go modules aren't loading correctly.**

The `go` modules also don't have proper type definitions, so you'll want to use `//@ts-expect-error` above any lines where you use them, to suppress the compiler error.

If you want to deploy your BattleSnake to DigitalOcean and are setting up a new account, you can get
$100 free credit for 60 days by going to [do.co/battlesnake](https://do.co/battlesnake).


# Penelope

Alrighty then.  Firstly: you need to be using a modern-ish node. `nvm use 16` works for me.

If you're using VS Code, you need to set Go in to WASM mode:

open .vscode/settings.json (cmd+shift+p and then "open workspace settings.json")
```json
{
    "go.toolsEnvVars": {"GOOS": "js", "GOARCH": "wasm"}
}
```

The entire build/run loop is now one script: `./script/server`. We run some tests
at the start of the server. If these fail they do print a nasty backtrace, but
they don't crash the server. BE CAREFUL.

In terms of actual behaviour: you'll find a minmax package in the go code that implements
minmax. Tweaking this can be as easy as playing with the heuristic function, but braver
programmers than I might also want to mess with depth, alpha beta pruning, and other options.

## Rob

The server should now support all game modes (including wrapped and constrictor) and respect hazards!
The minmax heuristic was also tweaked in a possibly no-op fashion.

Use Penelope's instructions above - the setup is unchanged. Don't forget to `npm install` before you run the server!
