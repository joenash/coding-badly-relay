package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
	"time"

	"github.com/BattlesnakeOfficial/rules"
)

func init() {
	// we have to declare our functions in an init func otherwise they aren't
	// available in JS land at the call time.
	js.Global().Set("go_ADD_STUFF", js.FuncOf(add))
	js.Global().Set("go_GAME_BOARD_FROM_JS", js.FuncOf(gameBoardFromJS))
	js.Global().Set("go_NEXT_BOARD_STATE_ELIMINATION_CAUSE", js.FuncOf(nextBoardStateEliminationCause))
}

func main() {
	fmt.Println("loading")
	wait()
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

func add(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(args[0].Int() + args[1].Int())
}

func gameBoardFromJS(this js.Value, args []js.Value) interface{} {
	boardStr := args[0].String()
	return boardFromJson(boardStr).Turn
}

func nextBoardStateEliminationCause(this js.Value, args []js.Value) interface{} {
	game := args[0]
	movesJS := args[1]

	board := boardFromJSValue(game)
	// check boards match

	var movesArr2 []rules.SnakeMove
	// movesJS is an array of js values of strings and strings like "move", unpack it in to a rules.SnakeMove
	for i := 0; i < movesJS.Length(); i++ {
		move := movesJS.Index(i)
		movesArr2 = append(movesArr2, rules.SnakeMove{
			ID:   move.Get("Id").String(),
			Move: move.Get("Move").String(),
		})
	}

	standard := rules.StandardRuleset{
		FoodSpawnChance:     5,
		MinimumFood:         3,
		HazardDamagePerTurn: 0,
	}

	nextBoardState, _ := standard.CreateNextBoardState(&board, movesArr2)

	me := nextBoardState.Snakes[0]

	return js.ValueOf([]interface{}{me.EliminatedCause, true})
}

func toInt(arg interface{}) int32 {
	return int32(arg.(float64))
}

func toPointArray(arr []interface{}) []rules.Point {
	var positions []rules.Point

	for _, f := range arr {
		fMap := f.(map[string]interface{})
		positions = append(positions, rules.Point{
			X: toInt(fMap["x"]),
			Y: toInt(fMap["y"]),
		})
	}

	return positions
}

func toPointArrayJS(arr js.Value) []rules.Point {
	var positions []rules.Point

	for i := 0; i < arr.Length(); i++ {
		f := arr.Index(i)
		positions = append(positions, rules.Point{
			X: int32(f.Get("x").Int()),
			Y: int32(f.Get("y").Int()),
		})
	}

	return positions
}

func boardFromJson(boardStr string) rules.BoardState {
	var object map[string]interface{}
	json.Unmarshal([]byte(boardStr), &object)

	board := object["board"].(map[string]interface{})
	food := board["food"].([]interface{})
	hazards := board["hazards"].([]interface{})

	snakes := board["snakes"].([]interface{})

	var typedSnakes []rules.Snake
	for _, s := range snakes {
		sMap := s.(map[string]interface{})

		typedSnakes = append(typedSnakes, rules.Snake{
			ID:     sMap["id"].(string),
			Health: toInt(sMap["health"]),
			Body:   toPointArray(sMap["body"].([]interface{})),
		})
	}

	foodPos := toPointArray(food)

	return rules.BoardState{
		Turn:    toInt(object["turn"]),
		Height:  toInt(board["height"]),
		Width:   toInt(board["width"]),
		Food:    foodPos,
		Hazards: toPointArray(hazards),
		Snakes:  typedSnakes,
	}
}

func boardFromJSValue(gamejs js.Value) rules.BoardState {
	board := gamejs.Get("board")
	food := board.Get("food")
	hazards := board.Get("hazards")

	snakes := board.Get("snakes")

	var typedSnakes []rules.Snake

	for i := 0; i < snakes.Length(); i++ {
		s := snakes.Index(i)
		typedSnakes = append(typedSnakes, rules.Snake{
			ID:     s.Get("id").String(),
			Health: int32(s.Get("health").Int()),
			Body:   toPointArrayJS(s.Get("body")),
		})
	}

	foodPos := toPointArrayJS(food)

	return rules.BoardState{
		Turn:    int32(gamejs.Get("turn").Int()),
		Height:  int32(board.Get("height").Int()),
		Width:   int32(board.Get("width").Int()),
		Food:    foodPos,
		Hazards: toPointArrayJS(hazards),
		Snakes:  typedSnakes,
	}
}
