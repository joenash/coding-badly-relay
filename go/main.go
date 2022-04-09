package main

import (
	"fmt"
	"math"
	"syscall/js"
	"time"

	"github.com/BattlesnakeOfficial/rules"
)

var (
	allMoves = []string{rules.MoveUp, rules.MoveDown, rules.MoveRight, rules.MoveLeft}
)

const (
	MINMAX_MINIMIZING   = iota
	MINMAX_MAXIMIZING   = iota
	MAX_OPPONENT_SNAKES = 3
)

func main() {
	fmt.Println("loading")
	wait()
}

func init() {
	// we have to declare our functions in an init func otherwise they aren't
	// available in JS land at the call time.
	js.Global().Set("go_MINMAX", js.FuncOf(minMax))
	tests()
}

func minMax(this js.Value, args []js.Value) interface{} {
	gameJSObject := args[0]
	youID := args[1].String()
	depth := args[2].Int()

	boardObject := boardFromJSValue(gameJSObject)
	bestMove, score := realMinMax(&boardObject, youID, depth)
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

func cartersianProductOfStringArrays(arrays ...[]string) [][]string {
	var result [][]string
	if len(arrays) == 1 {
		for _, a := range arrays[0] {
			result = append(result, []string{a})
		}
		return result
	}
	first := arrays[0]
	cartOfTail := cartersianProductOfStringArrays(arrays[1:]...)
	for _, head := range first {
		for _, tail := range cartOfTail {
			result = append(result, append([]string{head}, tail...))
		}
	}

	return result
}

type action struct {
	youMove    string
	otherMoves []string
}

func buildSnakeIds(board *rules.BoardState, youID string, countOfAliveSnakes int) []string {
	snakeIds := make([]string, countOfAliveSnakes)
	writtenCount := 1
	snakeIds[0] = youID
	for _, snake := range board.Snakes {
		if snake.ID != youID {
			snakeIds[writtenCount] = snake.ID
			writtenCount++
		}
	}

	return snakeIds
}

func realMinMax(board *rules.BoardState, youID string, depth int) (string, float32) {
	countOfAliveSnakes := len(board.Snakes)

	snakeIds := buildSnakeIds(board, youID, countOfAliveSnakes)

	// repeat all moves count of snakes times
	possibleMovesForSnakes := buildPossibleMoves(countOfAliveSnakes, board, snakeIds, depth)
	if depth == 3 {
		fmt.Println("realMinmax", possibleMovesForSnakes)
	}
	standard := rules.StandardRuleset{
		FoodSpawnChance:     5,
		MinimumFood:         3,
		HazardDamagePerTurn: 0,
	}

	movesProduct := cartersianProductOfStringArrays(possibleMovesForSnakes...)
	payOffTable := map[string]map[[MAX_OPPONENT_SNAKES]string]float32{}
	for _, moveList := range movesProduct {
		// simulate the game and get the new boards

		movesWithSnakeIds := []rules.SnakeMove{}
		for i, move := range moveList {
			movesWithSnakeIds = append(movesWithSnakeIds, rules.SnakeMove{
				ID:   snakeIds[i],
				Move: move,
			})
		}

		newBoard, err := standard.CreateNextBoardState(board, movesWithSnakeIds)
		if err != nil {
			panic(err)
		}
		youMove := moveList[0]
		var otherMoves [MAX_OPPONENT_SNAKES]string
		if len(moveList)-1 > MAX_OPPONENT_SNAKES {
			panic("too many snakes")
		}
		for i, move := range moveList[1:] {
			otherMoves[i] = move
		}
		if payOffTable[youMove] == nil {
			payOffTable[youMove] = make(map[[MAX_OPPONENT_SNAKES]string]float32)
		}
		if depth == 0 {
			scoreForState := heuristic(newBoard, youID)
			// initialize the payOffTable
			payOffTable[youMove][otherMoves] = scoreForState
		} else {
			if moveList[0] == "up" && moveList[1] == "right" && depth == 3 {
				you := findYou(newBoard, youID)
				fmt.Println("you", you)
				fmt.Println("elimination cause", you.EliminatedCause)
			}
			you := findYou(newBoard, youID)
			if you.EliminatedCause != "" {
				payOffTable[youMove][otherMoves] = -1.0
			} else {
				_, score := realMinMax(newBoard, youID, depth-1)
				payOffTable[youMove][otherMoves] = score
			}
		}
	}

	// print depth and playoff table
	if depth == 3 {
		fmt.Println("------------", depth, "--------------")
		for youMove, otherMoves := range payOffTable {
			for otherMove, score := range otherMoves {
				fmt.Println(youMove, otherMove, score)
			}
		}
	}

	// accross the inner dimension take the min, then take the max of that
	var bestMove string
	bestScore := float32(math.Inf(-1))
	for youMove, inner := range payOffTable {
		worstScore := float32(math.Inf(1))
		for _, score := range inner {
			if score < worstScore {
				worstScore = score
			}
		}
		if worstScore > bestScore {
			bestScore = worstScore
			bestMove = youMove
		}
	}

	return bestMove, bestScore
}

func buildPossibleMoves(countOfAliveSnakes int, board *rules.BoardState, snakeIds []string, depth int) [][]string {
	if depth == 3 {
		fmt.Println("buildPossibleMoves", countOfAliveSnakes, board, snakeIds, depth)
	}
	possibleMovesForSnakes := make([][]string, countOfAliveSnakes)
	for i := 0; i < countOfAliveSnakes; i++ {
		newAllMoves := make([]string, len(allMoves))
		copy(newAllMoves, allMoves)
		possibleMovesForSnakes[i] = newAllMoves
	}

	for i := 0; i < countOfAliveSnakes; i++ {
		possibleMovesForSnakes[i] = filterForValidMoves(board, snakeIds[i], possibleMovesForSnakes[i])
	}
	if depth == 3 {
		fmt.Println("possible moves for snakes: ", possibleMovesForSnakes)
	}
	return possibleMovesForSnakes
}

func findYou(board *rules.BoardState, youID string) *rules.Snake {
	for _, snake := range board.Snakes {
		if snake.ID == youID {
			return &snake
		}
	}
	panic("you not found")
}

func heuristic(board *rules.BoardState, youID string) float32 {
	// find the snake that is you
	you := findYou(board, youID)

	if you.EliminatedCause != "" {
		return -1.0
	}

	// if you are the longest snake return 0.8
	if len(board.Snakes) == 1 {
		return 0.8
	}
	for _, snake := range board.Snakes {
		if snake.ID != youID {
			if len(snake.Body) > len(you.Body) {
				return 0.0
			}
		}
	}
	return 0.8
}

func filterForValidMoves(board *rules.BoardState, snakeId string, possibleMoves []string) []string {
	// find the snake in the snakes array
	snakeIdx := -1

	for i, s := range board.Snakes {
		if s.ID == snakeId {
			snakeIdx = i
		}
	}
	snake := &board.Snakes[snakeIdx]
	if snake == nil {
		panic("snake not found")
	}
	head := snake.Body[0]
	validMoves := []string{}

	// eliminate top down left right
	for _, move := range possibleMoves {
		invalid := false
		if move == rules.MoveUp && head.Y == board.Height-1 {
			invalid = true
		}
		if move == rules.MoveDown && head.Y == 0 {
			invalid = true
		}
		if move == rules.MoveLeft && head.X == 0 {
			invalid = true
		}
		if move == rules.MoveRight && head.X == board.Width-1 {
			invalid = true
		}
		if !invalid {
			validMoves = append(validMoves, move)
		}
	}

	possibleMoves = validMoves
	validMoves = []string{}

	// eliminate neck moves
	neck := snake.Body[1]
	for _, move := range possibleMoves {
		invalid := false
		if move == rules.MoveUp && head.Y+1 == neck.Y {
			invalid = true
		}
		if move == rules.MoveDown && head.Y-1 == neck.Y {
			invalid = true
		}
		if move == rules.MoveLeft && head.X-1 == neck.X {
			invalid = true
		}
		if move == rules.MoveRight && head.X+1 == neck.X {
			invalid = true
		}

		if !invalid {
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
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
