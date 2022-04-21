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
export function cartesian<T>(sets: T[][]): T[][] {
  return sets.reduce<T[][]>(
    (results, ids) =>
      results
        .map(result => ids.map(id => [...result, id]))
        .reduce((nested, result) => [...nested, ...result]),
    [[]]
  );
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
  const youId = gameState.you.id;

  // @ts-expect-error Go modules not typed.
  const [move, score]: [String, number] = go_MINMAX(gameState, youId, 3);

  const response = {
    move: move,
  };

  console.log(`${gameState.game.id} MOVE ${gameState.turn}: ${response.move}`);
  return response;
}
