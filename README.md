# Coding Badly Relay
[![Deploy to DO](https://www.deploytodo.com/do-btn-blue.svg)](https://cloud.digitalocean.com/apps/new?repo=https://github.com/masonegger/coding-badly-relay/tree/main)

This is the Coding Badly Relay Snake. Each episode of Coding Badly, a new guest works on the snake, adding their own code and strategies. At the end of that episode, it's handed on, for the next person.

The Relay snake started out as the Battlesnake Javascript Starter. Where will it end up?

Check below for the schedule!

## Schedule

All shows are on the [Battlesnake Twitch](https://battlesnake.tv) at 7pm UTC. [Find your timezone](https://everytimezone.com/convert/utc/7pm).

| Date          | Guest | Final commit |
| ------------- | ----- | ------------ |
| 3rd February  |   coreyja    |       [ad5ba0d](https://github.com/joenash/coding-badly-relay/commit/ad5ba0d2076c312619089f68496bef29d484b3eb)       |
| 17th February | masonegger  |              |
| 3rd March     | geeh  |              |
| 17th March    |   nhcarrigan    |              |
| 7th April     |   penelope_zone    |              |
| 21st April    |       |              |
| 5th May       |   pachicodes   |              |
| 19th May      |       |              |

## Joining the show

Want to be a guest in the relay? Contact [Joe](https://twitter.com/jna_sh) or [Kevin](https://twitter.com/_phzn). Each guest will be building on the snake built in the show prior. You'll have approximately two weeks warning to know what terrible nonsense the last holder of the baton has gotten up to, to work out what you're going to do about it. Shows are live, broadcast on Twitch, via Streamyard. Every guest will be accompanied by at least one (usually both) of the hosts.

Coding Badly strives to be a low-pressure, fun environment to enjoy silly code that does useless things. That isn't to say we don't welcome good code, but it is to say that if this is your first time live coding on a stream, or your first time using Javascript, or your first time seeing Battlesnake: do not worry, this is the space for you.

All guests will be expected to abide by the Battlesnake Code of Conduct, both on air and in their submitted code.

## Technologies Used (so far)

- [JavaScript](https://www.javascript.com/)
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
./build
```
