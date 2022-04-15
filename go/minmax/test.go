package minmax

import (
	"encoding/json"
	"fmt"
	"reflect"
	"syscall/js"

	"official-go-wasm/serde"

	"github.com/BattlesnakeOfficial/rules"
)

func fixtureToBoard(fixture string) (*rules.BoardState, string) {
	var unpacked map[string]interface{}
	// load unpacked from json
	json.Unmarshal([]byte(fixture), &unpacked)
	youID := unpacked["you"].(map[string]interface{})["id"].(string)
	// convert to js.Value
	gamejs := js.ValueOf(unpacked)
	board := serde.BoardFromJSValue(gamejs)
	return &board, youID
}

func Tests() {
	{
		json_string := `{"game":{"id":"66635854-5cc5-4a0c-98a3-db06f52bd15a","ruleset":{"name":"standard","version":"v1.0.31","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14,"hazardMap":"","hazardMapAuthor":"","royale":{"shrinkEveryNTurns":0},"squad":{"allowBodyCollisions":false,"sharedElimination":false,"sharedHealth":false,"sharedLength":false}}},"timeout":500,"source":"custom"},"turn":11,"board":{"height":11,"width":11,"snakes":[{"id":"gs_WmkdCxPYwbvS87krGhKj743J","name":"does this work lol (original)","latency":"152","health":97,"body":[{"x":2,"y":5},{"x":3,"y":5},{"x":4,"y":5},{"x":5,"y":5},{"x":5,"y":4}],"head":{"x":2,"y":5},"length":5,"shout":"","squad":"","customizations":{"color":"#888888","head":"default","tail":"default"}},{"id":"gs_qF7pWbKdmwPvgXDrKF38fFRT","name":"badly","latency":"202","health":89,"body":[{"x":3,"y":6},{"x":3,"y":7},{"x":3,"y":8}],"head":{"x":3,"y":6},"length":3,"shout":"","squad":"","customizations":{"color":"#f22f46","head":"default","tail":"default"}}],"food":[{"x":0,"y":4},{"x":9,"y":8},{"x":8,"y":1}],"hazards":[]},"you":{"id":"gs_qF7pWbKdmwPvgXDrKF38fFRT","name":"badly","latency":"202","health":89,"body":[{"x":3,"y":6},{"x":3,"y":7},{"x":3,"y":8}],"head":{"x":3,"y":6},"length":3,"shout":"","squad":"","customizations":{"color":"#f22f46","head":"default","tail":"default"}}}`
		board, youID := fixtureToBoard(json_string)
		countOfAliveSnakes := len(board.Snakes)
		snakeIds := buildSnakeIds(board, youID, countOfAliveSnakes)

		// repeat all moves count of snakes times
		possibleMovesForSnakes := buildPossibleMoves(countOfAliveSnakes, board, snakeIds, 3)

		if !reflect.DeepEqual([][]string{{"down", "right", "left"}, {"up", "down", "left"}}, possibleMovesForSnakes) {
			fmt.Println("possibleMovesForSnakes is not as expected", possibleMovesForSnakes)
			panic("tests failed")
		}
	}
	{
		json_string := `{"game":{"id":"e7005bc9-1ec7-4ad7-bb9b-f75e7f6d5f2a","ruleset":{"name":"standard","version":"?","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14,"royale":{},"squad":{"allowBodyCollisions":false,"sharedElimination":false,"sharedHealth":false,"sharedLength":false}}},"timeout":500,"source":"custom"},"turn":25,"board":{"width":11,"height":11,"food":[{"x":1,"y":0}],"hazards":[],"snakes":[{"id":"gs_r9hGXXcptJ4Kwq3PDqtGKgCY","name":"does this work lol (original)","body":[{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":4,"y":8},{"x":5,"y":8},{"x":5,"y":7},{"x":5,"y":6}],"health":98,"latency":63,"head":{"x":6,"y":9},"length":7,"shout":"","squad":""},{"id":"gs_BKVtBTjpBbCbmy4mcRjRB6J7","name":"badly","body":[{"x":7,"y":8},{"x":7,"y":7},{"x":7,"y":6},{"x":7,"y":5}],"health":77,"latency":318,"head":{"x":7,"y":8},"length":4,"shout":"","squad":""}]},"you":{"id":"gs_BKVtBTjpBbCbmy4mcRjRB6J7","name":"badly","body":[{"x":7,"y":8},{"x":7,"y":7},{"x":7,"y":6},{"x":7,"y":5}],"health":77,"latency":318,"head":{"x":7,"y":8},"length":4,"shout":"","squad":""}}`
		board, youID := fixtureToBoard(json_string)
		countOfAliveSnakes := len(board.Snakes)
		snakeIds := buildSnakeIds(board, youID, countOfAliveSnakes)

		// repeat all moves count of snakes times
		possibleMovesForSnakes := buildPossibleMoves(countOfAliveSnakes, board, snakeIds, 3)

		if !reflect.DeepEqual([][]string{{"up", "right", "left"}, {"up", "down", "right"}}, possibleMovesForSnakes) {
			fmt.Println("possibleMovesForSnakes is not as expected", possibleMovesForSnakes)
			panic("tests failed")
		}
	}

	{
		json_string := `{"game":{"id":"e7005bc9-1ec7-4ad7-bb9b-f75e7f6d5f2a","ruleset":{"name":"standard","version":"?","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14,"royale":{},"squad":{"allowBodyCollisions":false,"sharedElimination":false,"sharedHealth":false,"sharedLength":false}}},"timeout":500,"source":"custom"},"turn":25,"board":{"width":11,"height":11,"food":[{"x":1,"y":0}],"hazards":[],"snakes":[{"id":"gs_r9hGXXcptJ4Kwq3PDqtGKgCY","name":"does this work lol (original)","body":[{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":4,"y":8},{"x":5,"y":8},{"x":5,"y":7},{"x":5,"y":6}],"health":98,"latency":63,"head":{"x":6,"y":9},"length":7,"shout":"","squad":""},{"id":"gs_BKVtBTjpBbCbmy4mcRjRB6J7","name":"badly","body":[{"x":7,"y":8},{"x":7,"y":7},{"x":7,"y":6},{"x":7,"y":5}],"health":77,"latency":318,"head":{"x":7,"y":8},"length":4,"shout":"","squad":""}]},"you":{"id":"gs_BKVtBTjpBbCbmy4mcRjRB6J7","name":"badly","body":[{"x":7,"y":8},{"x":7,"y":7},{"x":7,"y":6},{"x":7,"y":5}],"health":77,"latency":318,"head":{"x":7,"y":8},"length":4,"shout":"","squad":""}}`
		board, youID := fixtureToBoard(json_string)
		best_move, _ := RealMinMax(board, youID, 3)
		if best_move != "right" {
			fmt.Println("best_move is not as expected", best_move)
			panic("tests failed")
		}
	}
}
