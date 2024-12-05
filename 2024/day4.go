package main

import (
    "fmt"
    "bufio"
    "os"
    //"slices"
    //"regexp"
    //"strings"
    //"strconv"
)

func main() {
    puzzle := getPuzzle("day4Input.txt")
    xmasCount := findChar(puzzle, 'X', "XMAS")
    masCount := findChar(puzzle, 'A', "MAS")
    fmt.Println("Total XMAS found: ", xmasCount)
    fmt.Println("Total X-MAS found: ", masCount)
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

func findChar(puzzle []string, firstChar byte, findStr string) (int) {
    searchDirections := [][]int{
        {0, 1}, // To the right
	{0, -1}, // To the left
	{1, 0}, // down
        {-1, 0}, // up
	{1, -1}, // down diag to left
	{1, 1}, // down diag to right
	{-1, -1}, // up diag to left
        {-1, 1},  // up diag to right
    }
    xMax, yMax := len(puzzle), len(puzzle[0]) // set max row and columns
    //xPos, yPos := 0, 0 // start at top left cell
    xmasCount := 0

    for xPos :=0; xPos < xMax; xPos++ {
	for yPos := 0; yPos < yMax; yPos++ {    
	    if puzzle[xPos][yPos] == firstChar {
		if firstChar == 'X' {
	            for _, direction := range searchDirections {
	                if searchxMas(puzzle, findStr, xPos, yPos, direction, xMax, yMax) {
	                    xmasCount++
	                }
                    }
		} else if firstChar == 'A' {
			if cornerChk(puzzle, xPos, yPos, xMax, yMax) {
		            xmasCount++
			}
	            }
		}
            }
        }
    return xmasCount
}

func searchxMas(puzzle []string, findStr string, xPos, yPos int, direction []int, xMax, yMax int) bool {
    dX, dY := direction[0], direction[1]
    for i := 0; i < len(findStr); i++ {
	nextX := xPos + i*dX
	nextY := yPos + i*dY
	if nextX <0 || nextX >= xMax || nextY < 0 || nextY >= yMax {
	    return false
	}

	if puzzle[nextX][nextY] != findStr[i] {
	    return false
	}
    }
    return true
}

func cornerChk(puzzle []string, xPos, yPos, xMax, yMax int) bool {
    if xPos-1 < 0 || xPos+1 >= xMax || yPos-1 < 0 || yPos+1 >= yMax {
        return false
    }
    topLeft := puzzle[xPos-1][yPos-1]
    topRight := puzzle[xPos-1][yPos+1]
    botLeft := puzzle[xPos+1][yPos-1]
    botRight := puzzle[xPos+1][yPos+1]
    //topLeft := []int { {-1, -1} }
    //topRight := []int { {-1, 1} }
    //botLeft := []int { {1, -1} }
    //botRight := []int { {1, 1} }
    // if (topLeft M and botRight S && topRight M and botLeft S) || (topLeft M and botRight S && topRight S and botLeft M) || ...
    if (topLeft == 'M' && topRight == 'S' && botLeft == 'S' && botRight == 'M') || 
        (topLeft == 'M' && topRight == 'S' && botLeft == 'M' && botRight == 'S') ||
	(topLeft == 'M' && topRight == 'M' && botLeft == 'S' && botRight == 'S') ||
        (topLeft == 'S' && topRight == 'M' && botLeft == 'S' && botRight == 'M') ||
	(topLeft == 'S' && topRight == 'M' && botLeft == 'M' && botRight == 'S') ||
        (topLeft == 'S' && topRight == 'S' && botLeft == 'M' && botRight == 'M') {  
        return true
    }
    return false
}
