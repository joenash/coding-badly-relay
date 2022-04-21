package minmax

import (
	"math"

	"github.com/BattlesnakeOfficial/rules"
)

const (
	MAX_OPPONENT_SNAKES = 3
)

func RealMinMax(ruleset rules.Ruleset, board *rules.BoardState, youID string, depth int) (string, float32) {
	countOfAliveSnakes := len(board.Snakes)

	snakeIds := buildSnakeIds(board, youID, countOfAliveSnakes)

	// this takes "all" moves (e.g. up down left right) and filters for each snake
	// the moves that won't take them off the board, and won't move them in to their neck
	// there is a possible bug here where that leaves 0 moves left, which will break everything
	// so do pay attention for that
	possibleMovesForSnakes := buildPossibleMoves(ruleset.Name(), countOfAliveSnakes, board, snakeIds, depth)

	// take the cartesian product of all moves that we can make, so that we
	// simulate all possible forward game states
	movesProduct := cartersianProductOfStringArrays(possibleMovesForSnakes...)

	// you can imagine this as a table where the first axis is moves you make
	// and the second axis is moves other snakes are making, and as we iterate
	// we fill it up with the results of the moves deeper in the tree
	// 0.3 here would be the actual score
	// opponent: | {left} | {right} | {up} | {down} |
	// you       |  x     |  x      |  x   |  x     |
	// {left}    |  0.3   | 0.3     |  0.3 |  0.3   |
	// {right}   |  0.3   | 0.3     |  0.3 |  0.3   |
	// {up}      |  0.3   | 0.3     |  0.3 |  0.3   |
	// {down}    |  0.3   | 0.3     |  0.3 |  0.3   |
	payOffTable := map[string]map[[MAX_OPPONENT_SNAKES]string]float32{}

	// for each row in the cartesian product, simulate it
	for _, moveList := range movesProduct {
		// simulate the game and get the new boards

		movesWithSnakeIds := zipTogetherMovesAndSnakeIDs(moveList, snakeIds)

		// run the simulation of the game with the given moves, this uses the battlesnake
		// rules repo
		newBoard, err := ruleset.CreateNextBoardState(board, movesWithSnakeIds)
		if err != nil {
			panic(err)
		}

		// split moves in to your moves and opponent moves
		youMove := moveList[0]
		var otherMoves [MAX_OPPONENT_SNAKES]string
		if len(moveList)-1 > MAX_OPPONENT_SNAKES {
			panic("too many snakes")
		}

		// staticcheck is wrong about this
		for i, move := range moveList[1:] {
			otherMoves[i] = move
		}

		if payOffTable[youMove] == nil {
			payOffTable[youMove] = make(map[[MAX_OPPONENT_SNAKES]string]float32)
		}

		// if we're at the leaves of the minmax tree, we run the heuristic over the board
		if depth == 0 {
			scoreForState := heuristic(newBoard, youID)
			// initialize the payOffTable
			payOffTable[youMove][otherMoves] = scoreForState
		} else {
			you := findYou(newBoard, youID)
			// if we're dead, we don't recurse
			if you.EliminatedCause != "" {
				payOffTable[youMove][otherMoves] = -1.0
			} else {
				_, score := RealMinMax(ruleset, newBoard, youID, depth-1)
				payOffTable[youMove][otherMoves] = score
			}
		}
	}

	// this is the "actual" minmax, it will find the "worst" an opponent
	// can do for a given move, and then pick the best move given all of that
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

func zipTogetherMovesAndSnakeIDs(moveList []string, snakeIds []string) []rules.SnakeMove {
	movesWithSnakeIds := []rules.SnakeMove{}
	for i, move := range moveList {
		movesWithSnakeIds = append(movesWithSnakeIds, rules.SnakeMove{
			ID:   snakeIds[i],
			Move: move,
		})
	}
	return movesWithSnakeIds
}
