package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    puzzle := getPuzzle("day4Tst.txt")

    xmasCount := findX(puzzle, "all", 'X', "XMAS")

    xShapeMASCount := searchMAS(puzzle)

    fmt.Println("Total XMAS found: ", xmasCount)
    fmt.Println("Total MAS found in X shape: ", xShapeMASCount)
}

func getPuzzle(filename string) []string {
    file, _ := os.Open(filename)
    defer file.Close()
    var puzzle []string

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        puzzle = append(puzzle, line)
    }
    return puzzle
}

func findX(puzzle []string, dir string, firstChar byte, findStr string) int {
    searchStraight := [][]int{
        {0, 1},   // To the right
        {0, -1},  // To the left
        {1, 0},   // Down
        {-1, 0},  // Up
    }
    searchDiag := [][]int{
        {1, -1},  // Down diagonal to left
        {1, 1},   // Down diagonal to right
        {-1, -1}, // Up diagonal to left
        {-1, 1},  // Up diagonal to right
    }
    var searchDirections [][]int

    if dir == "all" {
        searchDirections = append(searchStraight, searchDiag...)
    } else if dir == "diag" {
        searchDirections = searchDiag
    }

    xMax, yMax := len(puzzle), len(puzzle[0]) // set max row and columns
    xmasCount := 0

    for xPos := 0; xPos < xMax; xPos++ {
        for yPos := 0; yPos < yMax; yPos++ {
            if puzzle[xPos][yPos] == firstChar {
                for _, direction := range searchDirections {
                    if searchXmas(puzzle, findStr, xPos, yPos, direction, xMax, yMax) {
                        xmasCount++
                    }
                }
            }
        }
    }
    return xmasCount
}

func searchXmas(puzzle []string, findStr string, xPos, yPos int, direction []int, xMax, yMax int) bool {
    dX, dY := direction[0], direction[1]
    for i := 0; i < len(findStr); i++ {
        nextX := xPos + i*dX
        nextY := yPos + i*dY

        if nextX < 0 || nextX >= xMax || nextY < 0 || nextY >= yMax {
            return false
        }

        if puzzle[nextX][nextY] != findStr[i] {
            return false
        }
    }
    return true
}

func searchMAS(puzzle []string) int {
    diagonalPairs := [][2][]int{
        {{1, -1}, {-1, 1}}, // Down-left & Up-right
        {{1, 1}, {-1, -1}}, // Down-right & Up-left
    }

    xMax, yMax := len(puzzle), len(puzzle[0])
    xShapeMASCount := 0

    // Traverse the grid, avoiding boundaries (ensure we can look at both diagonals)
    for xPos := 1; xPos < xMax-1; xPos++ {
        for yPos := 1; yPos < yMax-1; yPos++ {
            for _, pair := range diagonalPairs {
                // Check for "MAS" in one diagonal and "SAM" in the other to form an X
                if searchXmas(puzzle, "MAS", xPos, yPos, pair[0], xMax, yMax) &&
                   searchXmas(puzzle, "MAS", xPos, yPos, pair[1], xMax, yMax) {
                    xShapeMASCount++
                }
            }
        }
    }

    return xShapeMASCount
}
