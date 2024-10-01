package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"
)

const (
    WIDTH  = 5
    HEIGHT = 5
)

type Position struct {
    x int
    y int
}

func main() {
    rand.Seed(time.Now().UnixNano())

    var maze [][]string
    for {
        maze = generateMaze()
        if isSolvableMaze(maze) {
            break
        }
    }

    playerPos := Position{0, 0}
    goalPos := Position{4, 4}
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Welcome to the Maze Adventure!")
    fmt.Println("Navigate through the maze to reach the goal (G).")
    fmt.Println("You can move using 'up', 'down', 'left', or 'right'. Enter 'quit' to exit.")

    for {
        printMaze(maze, playerPos)

        if playerPos == goalPos {
            fmt.Println("Congratulations! You've reached the goal!")
            break
        }

        fmt.Print("Enter your move: ")
        move, _ := reader.ReadString('\n')
        move = strings.TrimSpace(strings.ToLower(move))

        if move == "quit" {
            fmt.Println("Thanks for playing! Goodbye!")
            break
        }

        newPos := playerPos

        switch move {
        case "up":
            newPos.y--
        case "down":
            newPos.y++
        case "left":
            newPos.x--
        case "right":
            newPos.x++
        default:
            fmt.Println("Invalid move. Use 'up', 'down', 'left', or 'right'.")
            continue
        }

        if isValidMove(newPos, maze) {
            playerPos = newPos
        } else {
            fmt.Println("That path is blocked! Try a different direction.")
        }
    }
}

func generateMaze() [][]string {
    maze := make([][]string, HEIGHT)
    for i := range maze {
        maze[i] = make([]string, WIDTH)
        for j := range maze[i] {
            maze[i][j] = " "
        }
    }

    maze[0][0] = "S"
    maze[HEIGHT-1][WIDTH-1] = "G"

    numObstacles := rand.Intn(WIDTH*HEIGHT/2) + WIDTH
    for i := 0; i < numObstacles; i++ {
        x := rand.Intn(WIDTH)
        y := rand.Intn(HEIGHT)

        if (x == 0 && y == 0) || (x == WIDTH-1 && y == HEIGHT-1) || maze[y][x] == "X" {
            continue
        }

        maze[y][x] = "X"
    }

    return maze
}

func isSolvableMaze(maze [][]string) bool {
    visited := make([][]bool, HEIGHT)
    for i := range visited {
        visited[i] = make([]bool, WIDTH)
    }

    return dfs(Position{0, 0}, maze, visited)
}

func dfs(pos Position, maze [][]string, visited [][]bool) bool {
    if pos.x < 0 || pos.x >= WIDTH || pos.y < 0 || pos.y >= HEIGHT {
        return false
    }

    if maze[pos.y][pos.x] == "X" || visited[pos.y][pos.x] {
        return false
    }

    if pos.x == WIDTH-1 && pos.y == HEIGHT-1 {
        return true
    }

    visited[pos.y][pos.x] = true

    directions := []Position{
        {0, -1},  // up
        {0, 1},   // down
        {-1, 0},  // left
        {1, 0},   // right
    }

    for _, dir := range directions {
        newPos := Position{pos.x + dir.x, pos.y + dir.y}
        if dfs(newPos, maze, visited) {
            return true
        }
    }

    return false
}

func printMaze(maze [][]string, playerPos Position) {
    for y := 0; y < HEIGHT; y++ {
        for x := 0; x < WIDTH; x++ {
            if playerPos.x == x && playerPos.y == y {
                fmt.Print("P ")
            } else {
                fmt.Print(maze[y][x], " ")
            }
        }
        fmt.Println()
    }
}

func isValidMove(pos Position, maze [][]string) bool {
    if pos.x < 0 || pos.x >= WIDTH || pos.y < 0 || pos.y >= HEIGHT {
        return false
    }
    return maze[pos.y][pos.x] != "X"
}
