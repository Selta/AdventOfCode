package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
    puzzle := getPuzzle("day5Input.txt")
    rules, updates := getRulesPages(puzzle)
    //fmt.Println("Rules: ", rules)
    //fmt.Println("updates: ", updates)
    sumMiddlePages, _ := validateAndSum(rules, updates)
    _, sumCorrectedPages := validateAndSum(rules,updates)
    fmt.Println("Sum of middle pages: ", sumMiddlePages)
    fmt.Println("Sum of corrected middle pages: ", sumCorrectedPages)

}

func getPuzzle(filename string) []string {
    file, err := os.Open(filename)
    if err != nil {
    panic(err)
    }
    defer file.Close()

    var puzzle []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    puzzle = append(puzzle, scanner.Text())
    }
    return puzzle
}

func getRulesPages(puzzle []string) (map[int][]int, [][]int) {
    rules := make(map[int][]int)
    var updates [][]int

    isRulesSection := true
    for _, line := range puzzle {
    line = strings.TrimSpace(line)
    if line == "" {
        isRulesSection = false
        continue
    }

    if isRulesSection {
        parts := strings.Split(line, "|")
        left, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
        right, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
        rules[left] = append(rules[left], right)
    } else {
        update := []int{}
        for _, numStr := range strings.Split(line, ",") {
        num, _ := strconv.Atoi(strings.TrimSpace(numStr))
        update = append(update, num)
        }
        updates = append(updates, update)
    }
    }
    return rules, updates
}

func validateAndSum(rules map[int][]int, updates [][]int) (int, int) {
    sum := 0
    sumCr := 0

    for _, update := range updates {
    if isValidUpdate(rules, update) {
        middleIndex := len(update) / 2
        sum += update[middleIndex]
    } else {
        correctOrder := reorderUpdate(rules, update)
        middleIndexCr := len(correctOrder)/2
        sumCr += correctOrder[middleIndexCr]
    }
    }
    return sum, sumCr
}

func isValidUpdate(rules map[int][]int, update []int) bool {
    position := make(map[int]int)
    for i, page := range update {
    position[page] = i
    }

    for x, dependents := range rules {
    if posX, existsX := position[x]; existsX {
        for _, y := range dependents {
        if posY, existsY := position[y]; existsY {
            // if you got this far in this mess of if and for, that means the order is bad
            if posX > posY {
            return false
            }
        }
        }
    }
    }
    return true
}

func reorderUpdate(rules map[int][]int, update []int) []int {
    // directed graph from the rules relevant to the update
    // https://www.codeglimpse.com/posts/exploring-directed-acyclic-graphs-in-golang
    // nodes == pages
    // an edge = page to page rule (ex 47|53)
     // assuming acyclic, as there shouldn't be both 47|53 amd 53|47 like rules
    graph := make(map[int][]int)
    inDegree := make(map[int]int)

    for _, page := range update {
        graph[page] = []int{}
        inDegree[page] = 0
    }

    // edges (page to page) based on rule
    for x, dependents := range rules {
        for _, y := range dependents {
            if containsChk(update, x) && containsChk(update, y) {
                graph[x] = append(graph[x], y)
                inDegree[y]++
            }
        }
    }

    // topological sort
    var sorted []int
    queue := []int{}

    // Find all nodes with in-degree 0
    for page, degree := range inDegree {
        if degree == 0 {
            queue = append(queue, page)
        }
    }

    // do the actual sorting in the dag
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        sorted = append(sorted, current)

        for _, neighbor := range graph[current] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }
    return sorted
}

func containsChk(slice []int, value int) bool {
    for _, v := range slice {
        if v == value {
            return true
        }
    }
    return false

}
