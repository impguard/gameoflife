from random import randint
from typing import Sequence, Tuple, NamedTuple, List


State = NamedTuple('State', [
    ('board', List[bool]),
    ('width', int),
    ('height', int)
])


def cell_index(state: State, x: int, y: int) -> int:
    modx = x % state.width
    mody = y % state.height

    return modx + mody * state.height


def neighbor_indices(state: State, x: int, y: int) -> Sequence[int]:
    neighbors = [
        (x - 1, y - 1), (x, y - 1), (x + 1, y - 1),
        (x - 1, y), (x + 1, y),
        (x - 1, y + 1), (x, y + 1), (x + 1, y + 1)
    ]

    return [cell_index(state, cell[0], cell[1]) for cell in neighbors]


def neighbor_values(state: State, x, y):
    neighbors = neighbor_indices(state, x, y)
    return [state.board[neighbor] for neighbor in neighbors]


def cell_xy(state: State, cell_index: int) -> Tuple[int, int]:
    y = cell_index // state.height
    x = cell_index - y * state.height

    return (x, y)


def transition(state: State) -> State:
    """ Transitions from one state to the next. """
    next_board = [False] * state.width * state.height

    for y in range(state.height):
        for x in range(state.width):
            index = cell_index(state, x, y)
            values = neighbor_values(state, x, y)
            if sum(values) == 2:
                next_board[index] = state.board[index]
            elif sum(values) == 3:
                next_board[index] = True

    return State(next_board, state.width, state.height)


def state_to_str(state: State) -> str:
    output = []
    for y in range(state.height):
        for x in range(state.width):
            index = cell_index(state, x, y)
            is_alive = state.board[index]
            cellsprite = 'o' if is_alive else '.'
            output.append(cellsprite)
        output.append('\n')

    return ''.join(output)


def clear_state_str(state: State) -> str:
    return '\b' * state.width * state.height + '\b' * state.height


def random_state(width: int, height: int) -> State:
    board = [False] * width * height
    state = State(board, width, height)

    for y in range(height):
        for x in range(width):
            index = cell_index(state, x, y)
            state.board[index] = randint(1, 2) == 1

    return state
