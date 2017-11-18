import java.util.Random


data class Game(val width: Int, val height: Int) {
    var board: List<Boolean> = BooleanArray(width * height).asList()
        private set

    constructor(width: Int, height: Int, board: List<Boolean>): this(width, height) {
        this.board = board
    }

    constructor(width: Int, height: Int, seed: Long = 0): this(width, height) {
        val random = Random(seed)
        val randomBoard = BooleanArray(width * height)

        for (y in 0..height-1) {
            for (x in 0..width - 1) {
                val index = cell_index(x, y)
                val value = random.nextInt(2) == 1

                randomBoard[index] = value
            }
        }

        board = randomBoard.asList()
    }

    fun cell_index(x: Int, y: Int): Int {
        val modx = ((x % width) + width) % width
        val mody = ((y % height) + height) % height

        return modx + mody * width
    }

    fun transition() {
        val newBoard = BooleanArray(width * height)

        for (y in 0..height-1) {
            for (x in 0..width-1) {
                val index = cell_index(x, y)
                val value = board[index]
            }
        }
    }

    override fun toString(): String {
        val result = StringBuilder()

        for (y in 0..height-1) {
            for (x in 0..width-1) {
                val index = cell_index(x, y)
                val string = if (board[index]) "o " else ". "

                result.append(string)
            }
            result.append("\n")
        }

        return result.subSequence(0, result.length - 1).toString()
    }
}

fun main(args: Array<String>) {
    val game = Game(3, 3, 0)
    println(game.cell_index(-1, 1))
    println("Hello, World!")
    println(game)
}
