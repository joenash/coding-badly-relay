package main

import (
	"official-go-wasm/minmax"
	"official-go-wasm/serde"
	"syscall/js"
	"time"

	"github.com/BattlesnakeOfficial/rules"
)

var (
	allMoves = []string{rules.MoveUp, rules.MoveDown, rules.MoveRight, rules.MoveLeft}
)

const (
	MAX_OPPONENT_SNAKES = 3
)

func main() {
	minmax.Tests()
	wait()
}

func init() {
	// we have to declare our functions in an init func otherwise they aren't
	// available in JS land at the call time.
	js.Global().Set("go_MINMAX", js.FuncOf(minMax))
}

func minMax(this js.Value, args []js.Value) interface{} {
	gameJSObject := args[0]
	youID := args[1].String()
	depth := args[2].Int()

	boardObject := serde.BoardFromJSValue(gameJSObject)
	bestMove, score := minmax.RealMinMax(&boardObject, youID, depth)
	jsBestMove := js.ValueOf(bestMove)
	jsScore := js.ValueOf(score)
	return []interface{}{jsBestMove, jsScore}
}

func wait() {
	done := make(chan bool)
	js.Global().Get("process").Call("on", "SIGTERM", js.FuncOf(func(js.Value, []js.Value) interface{} {
		done <- true
		return nil
	}))
	for {
		select {
		case <-done:
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}
