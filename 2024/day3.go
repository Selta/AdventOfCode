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
    muls := findMuls("day3Input.txt")
    fmt.Println("Muls found:", muls)
    results := mulMath(muls)
    total := sumResults(results)
    fmt.Println("the sum of multipliers is: ", total)
}

func findMuls(filename string) []string {
    file, _ := os.Open(filename)
    defer file.Close()

    mulRegEx := `mul\((\d{1,3}),\s*(\d{1,3})\)`
    regEx, _ := regexp.Compile(mulRegEx)

    var muls []string

    toRetain := true // Initially, mul instructions are enabled

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        startIdx := 0 // Start from the beginning of the line

        for startIdx < len(line) {
            dontIdx := strings.Index(line[startIdx:], "don't()") 
            doIdx := strings.Index(line[startIdx:], "do()")

            if dontIdx == -1 && doIdx == -1 {
                if toRetain {
                    mulsFound := regEx.FindAllString(line[startIdx:], -1)
                    if len(mulsFound) > 0 {
                        muls = append(muls, mulsFound...)
                    }
                }
                break
            }

            if dontIdx != -1 && (dontIdx < doIdx || doIdx == -1) {
                toRetain = false // Disable future mul instructions
                startIdx += dontIdx + 7 // Move past "don't()"
            } else if doIdx != -1 {
                toRetain = true // Enable future mul instructions
                startIdx += doIdx + 5 // Move past "do()"
            }

            if toRetain {
                mulsFound := regEx.FindAllString(line[startIdx:], -1)
                if len(mulsFound) > 0 {
                    muls = append(muls, mulsFound...)
                }
            }
	    startIdx++
        }
    }

    return muls
}

func mulMath(muls []string) ([]int) {
    var results []int
    for _, mul := range muls {
        mulNums := strings.TrimPrefix(mul, "mul(")
        mulNums = strings.TrimSuffix(mulNums, ")")
        factors := strings.Split(mulNums, ",")
    
        multicand, _ := strconv.Atoi(strings.TrimSpace(factors[0]))
        multiplier, _ := strconv.Atoi(strings.TrimSpace(factors[1]))

        result := multicand * multiplier
        results = append(results, result)
    }
    return results
}

func sumResults(results []int) (int) {
    total := 0
    for _, result := range results {
        total += result
    }
    return total
}
