package serde

import (
	"syscall/js"

	"github.com/BattlesnakeOfficial/rules"
)

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

func BoardFromJSValue(gamejs js.Value) rules.BoardState {
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
