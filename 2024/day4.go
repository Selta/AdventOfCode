package main

import (
    "fmt"
    "bufio"
    "os"
    "regexp"
    "strings"
    "strconv"
)

func main() {

}

func getPuzzle(filename string) []string {
    file, _ := os.Open(filename)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
    }    
    return puzzle
}

func findXmas(puzzle []string) (int) {
    searchDirections := [][]int{
	{0, 1}, // left to right
	{0, -1}, // right to left
	{1, 0}, // down
        {-1, 0}, // up
	{1, -1}, // down diag left
	{1, 1}, // down diag right
	{-1, -1}, // up diag left
        {-1, 1}  // up diag right

	xMax, yMax := len(puzzle), len(puzzle[0]) // set max row and columns
        xPos, yPos := 0, 0 // start at top left cell
	
	for xPos < xMax && yPos < yMax {
	    // if cell == X, search for Ms on any axis
	      // if M (can be multiple, so loop each found M here), search for A along same axis as M (can only be one)
	         // if A, search for S along same axis A (can only be one)
		    // if S along same axis found, counter++
	    // else move right and loop
	}


}
