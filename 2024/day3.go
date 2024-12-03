package main

import (
    "fmt"
    "bufio"
    "os"
    "regexp"
)

func main() {
    muls := findMuls("day3Input.txt")
    fmt.Println("Muls found:", muls)
}

func findMuls(filename string) []string {
    file, _ := os.Open(filename)
    defer file.Close()

    // not certain on this
    mulRegEx := `mul\((\d{1,3}),\s*(\d{1,3})\)`
    regEx, _ := regexp.Compile(mulRegEx)

    var muls []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        mulsFound := regEx.FindAllString(line, -1)
        if len(mulsFound) > 0 {
            muls = append(muls, mulsFound...)
        }
    }
    return muls
}

func mulMath(muls []string) ([]int) {
    for _, mul := range muls {
    // todo: extract the two number values
    // todo: convert those from string to int
    // todo: perform the math
    // todo: return the multiplication (to be summed in main)
    }
}
