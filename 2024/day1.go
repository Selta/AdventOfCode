package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "math"
    "sort"
)

func main() {
    left, right, _ := parseLocations("day1Input.txt")
    sortLocations(left, right)
    distances := calcEachDistance(left, right)
    distance := calcDistance(distances)
    fmt.Println("Total Distance: ", distance)
    counts := countRepeats(left, right) 
    // fmt.Println("the counts are: ", counts)
    similarity := calcSimilarity(left, counts)
    fmt.Println("The similarity score is: ", similarity)
}

func parseLocations(filename string) ([]int, []int, error) {
    file, err := os.Open("day1Input.txt")
    if err != nil {
	return nil, nil, fmt.Errorf("Error reading input file", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var left, right []int

    for scanner.Scan() {
        curLine := scanner.Text()
	locations := strings.Fields(curLine)
        if len(locations) == 2 {
            num1, _ := strconv.Atoi(locations[0])
	    num2, _ := strconv.Atoi(locations[1])
	
	    left = append(left, num1)
	    right = append(right, num2)
	}
    }
    return left, right, nil
}

func sortLocations(left, right []int) {
    sort.Ints(left)
    sort.Ints(right)
    return
}

func calcEachDistance(left, right []int) ([]int) {
    // we're reasonably sure there are the same number of indxes but let's validate
    if len(left) != len(right) {
        fmt.Println("Locations have different lengths")
	return nil
    }
    var distances []int

    for i := 0; i < len(left); i++ {
	abs := math.Abs(float64(left[i] - right[i]))
	distances = append(distances, int(abs))
    }
    return distances  
}

func calcDistance(distances []int) (int) {
    distance := 0
    for _, val := range distances {
        distance += val
    }
    return distance
}

func countRepeats(left, right []int) ([]int) {
    countMap := make(map[int]int)
    for _, value := range right {
	countMap[value]++
    }
    var counts []int
    for _, value := range left {
	counts = append(counts, countMap[value])
   }
   return counts
}

func calcSimilarity(left, counts []int) (int) {
    similarityScore := 0
    for i, count := range counts {
        if count > 0 {
            similarityScore += (left[i]*count)
        }
    }
    return similarityScore
}
