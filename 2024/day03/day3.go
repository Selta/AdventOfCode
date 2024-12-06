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

    mulRegEx := `mul\((\d{1,3}),\s*(\d{1,3})\)|do\(\)|don't\(\)`
    regEx, _ := regexp.Compile(mulRegEx)

    var muls []string

    toRetain := true // Initially, mul instructions are enabled

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        matches := regEx.FindAllString(line, -1)
	for _, match := range matches {
	    if match == "do()" {
		toRetain = true
	    } else if match == "don't()" {
		toRetain = false
            } else {
		if toRetain {
		    muls = append(muls, match)
	        }
	    }
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
