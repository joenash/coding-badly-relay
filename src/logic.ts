import { GameState } from "./interfaces/GameState";
import { getValidMove } from "./modules/getValidMove";

export function info() {
  console.log("INFO");
  const response = {
    apiversion: "1",
    author: "",
    color: "#F22F46",
    head: "trans-rights-scarf",
    tail: "do-sammy",
  };
  return response;
}

export function start(gameState: GameState) {
  console.log(`${gameState.game.id} START`);
}

export function end(gameState: GameState) {
  console.log(`${gameState.game.id} END\n`);
}

export function move(gameState: GameState) {
  let possibleMoves: { [key: string]: boolean } = {
    up: true,
    down: true,
    left: true,
    right: true,
  };

  console.log(gameState);

  const snakesOnBoard = gameState.board.snakes
    .slice(1)
    .map((snake) => ({ Id: snake.id, Move: getValidMove(gameState, snake.id) }));

  // this doesn't work because it selects a move (correctly) if a different snake would die moving up, we can't pass
  // 'up' as a default to all snakes -- we only tried to do that to brute force a different problem that I've
  // forgotten. GLHF

  for (const move in possibleMoves) {
    // @ts-expect-error Go modules not typed.
    const didWeDie = go_NEXT_BOARD_STATE_ELIMINATION_CAUSE(
      JSON.stringify(gameState),
      JSON.stringify([{ Id: gameState.you.id, Move: move }, ...snakesOnBoard])
    );

    console.log(didWeDie);

    if (didWeDie) {
      possibleMoves[move] = false;
    }
  }

  // TODO: Step 4 - Find food.
  // Use information in gameState to seek out and find food.

  // Finally, choose a move from the available safe moves.
  // TODO: Step 5 - Select a move to make based on strategy, rather than random.
  const safeMoves = Object.keys(possibleMoves).filter(
    (key) => possibleMoves[key]
  );
  const response = {
    move: safeMoves.shift(),
  };

  console.log(`${gameState.game.id} MOVE ${gameState.turn}: ${response.move}`);
  return response;
}
