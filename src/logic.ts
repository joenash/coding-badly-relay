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

export function move(gameState: GameState)  {
  let possibleMoves: { [key: string]: boolean } = {
    up: true,
    down: true,
    left: true,
    right: true,
  };

  console.log(gameState);

  const numberOfSnakesOnBoard = gameState.board.snakes.length;
  const all_moves = ["up", "down", "left", "right"];
  // copy all moves n times
  const moves_for_number_of_snakes: String[][] = Array(numberOfSnakesOnBoard).fill(all_moves);
  console.log(moves_for_number_of_snakes);

  // eliminate moves that aren't valid for a given snake
  const valid_moves_for_each_snake = gameState.board.snakes.map(snake => snake.id).map(snakeId => getValidMove(gameState, snakeId));
  for (let i = 0; i < numberOfSnakesOnBoard; i++) {
    const valid_moves_for_snake = valid_moves_for_each_snake[i];
    const moves_for_snake = moves_for_number_of_snakes[i];
    // select only the moves that are in the valid array
    moves_for_number_of_snakes[i] = moves_for_snake.filter(move => valid_moves_for_snake.includes(move));
  }

  console.log("------------");



  //// this doesn't work because it selects a move (correctly) if a different snake would die moving up, we can't pass
  //// 'up' as a default to all snakes -- we only tried to do that to brute force a different problem that I've
  //// forgotten. GLHF

  let minmax_tuples = [];
  let you_id = gameState.you.id;
  // find snake index from snake array based on id
  let you_index = gameState.board.snakes.findIndex(snake => snake.id === you_id);

  for (const move_tuple of cartesian(valid_moves_for_each_snake)) {
    // get ids from gameState.board.snakes
    const moves_with_ids = move_tuple.map((move, i) => ({ Id: gameState.board.snakes[i].id, Move: move }));
    console.log(moves_with_ids);
    // @ts-expect-error Go modules not typed.
    const values: String[] = go_NEXT_BOARD_STATE_ELIMINATION_CAUSE(
      gameState,
      moves_with_ids
    );

    // true if values[0] is not an empty string, an empty string means we lived
    const didWeDie = !!values[you_index];
    console.log(didWeDie);
    // look through values except you index for a true value
    const didAnyOtherSnakedie = values.some((value, index) => {
      if (index === you_index) {
        return false;
      }
      return !!value;
    });
    minmax_tuples.push({"self": didWeDie, "others": didAnyOtherSnakedie, "moves": move_tuple});
  }

  console.log(minmax_tuples);

  // x x H s s x x x x 
  // s H t x x x x x x 
  // s s s x x x x x x 
  // x x x x x x x x x 
  // x x x x x x x x x 
  // x x x x x x x x x 
  // x x x x x x x x x 
  // x x x x x x x x x 

  // find tuples where the other snake can certainly kill us

  let minmax_tuples_where_no_other_snake_dies = minmax_tuples.filter(tuple => tuple.others === false);

  let minmax_tuples_where_we_dont_die = minmax_tuples_where_no_other_snake_dies.filter(tuple => tuple.self === false);
  console.log(minmax_tuples_where_we_dont_die);
  // get move[0] out of the first one where we don't die
  let safe_move = minmax_tuples_where_we_dont_die[0].moves[you_index];

  const response = {
    move: safe_move,
  };

  //console.log(`${gameState.game.id} MOVE ${gameState.turn}: ${response.move}`);
  return response;
}
