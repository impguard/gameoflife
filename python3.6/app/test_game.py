import pytest

from app.game import State, cell_index, neighbor_indices, cell_xy, transition


@pytest.fixture()
def state():
    board = [
        False, False, False, False,
        False, False, False, False,
        False, False, False, False,
        False, False, False, False
    ]

    yield State(board, 4, 4)


def test_cell_xy_can_map_indices_to_appropriate_xy(state):
    assert cell_xy(state, 0) == (0, 0)
    assert cell_xy(state, 4) == (0, 1)
    assert cell_xy(state, 1) == (1, 0)
    assert cell_xy(state, 5) == (1, 1)
    assert cell_xy(state, 15) == (3, 3)


class TestCellIndex():
    def test_cell_index_finds_proper_indices(self, state):
        assert cell_index(state, 0, 0) == 0
        assert cell_index(state, 0, 1) == 4
        assert cell_index(state, 1, 0) == 1
        assert cell_index(state, 1, 1) == 5
        assert cell_index(state, 3, 3) == 15

    def test_cell_index_finds_proper_indices_with_wrapping(self, state):
        assert cell_index(state, -1, -1) == 15
        assert cell_index(state, -4, 0) == 0


class TestNeighborIndices():
    def test_neighbors_finds_common_neighbors(self, state):
        neighbor_cells = neighbor_indices(state, 1, 1)
        neighbor_set = set(neighbor_cells)

        assert neighbor_set == {
            0, 1, 2,
            4, 6,
            8, 9, 10
        }

    def test_neighbors_finds_wrapped_neighbors(self, state):
        neighbor_cells = neighbor_indices(state, 0, 0)
        neighbor_set = set(neighbor_cells)

        assert neighbor_set == {
            15, 12, 13,
            3, 1,
            7, 4, 5
        }


class TestTransition():
    @classmethod
    def state_with_n_neighbors(cls, n, self=True):
        board = [True] * n + [False] * (8 - n)
        board.insert(4, self)

        state = State(board, 3, 3)
        return state

    def test_underpopulation_no_neighbors(self, state: State):
        state = TestTransition.state_with_n_neighbors(0)
        next_state = transition(state)

        assert not next_state.board[4]

    def test_underpopulation_one_neighbor(self, state: State):
        state = TestTransition.state_with_n_neighbors(1)
        next_state = transition(state)

        assert not next_state.board[4]

    def test_live_two_neighbor(self, state: State):
        state = TestTransition.state_with_n_neighbors(2)
        next_state = transition(state)

        assert next_state.board[4]

    def test_live_three_neighbor(self, state: State):
        state = TestTransition.state_with_n_neighbors(3)
        next_state = transition(state)

        assert next_state.board[4]

    def test_overpopulation_four_neighbors(self, state: State):
        state = TestTransition.state_with_n_neighbors(4)
        next_state = transition(state)

        assert not next_state.board[4]

    def test_overpopulation_five_neighbors(self, state: State):
        state = TestTransition.state_with_n_neighbors(4)
        next_state = transition(state)

        assert not next_state.board[4]

    def test_overpopulation_six_neighbors(self, state: State):
        state = TestTransition.state_with_n_neighbors(6)
        next_state = transition(state)

        assert not next_state.board[4]

    def test_overpopulation_seven_neighbors(self, state: State):
        state = TestTransition.state_with_n_neighbors(7)
        next_state = transition(state)

        assert not next_state.board[4]

    def test_overpopulation_eight_neighbors(self, state: State):
        state = TestTransition.state_with_n_neighbors(8)
        next_state = transition(state)

        assert not next_state.board[4]

    def test_overpopulation_nine_neighbors(self, state: State):
        state = TestTransition.state_with_n_neighbors(9)
        next_state = transition(state)

        assert not next_state.board[4]

    def test_reproduction(self, state: State):
        state = TestTransition.state_with_n_neighbors(3, False)
        next_state = transition(state)

        assert next_state.board[4]
