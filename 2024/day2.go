package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
	"strconv"
	"strings"
)

// A report is a slice of levels. I know I don't have to do this but it was a rabbit hole of learning
type Report []int

func main() {
    reports, err := parseReports("day2Input.txt")
    if err != nil {
        fmt.Println("Error reading reports:", err)
	return
    }
    var safeReports, badReports, dampenedReports []Report
    for _, report := range reports {
	if inspectReport(report) {
	    safeReports = append(safeReports, report)
	} else {
	    badReports = append(badReports, report)
	}
    }
    fmt.Println("Safe Reports: ", len(safeReports))
    for _, report := range badReports {
	if dampenerCheck(report) {
	    dampenedReports = append(dampenedReports, report)
	}
    }
    fmt.Println("Safe and \"Dampened\" Reports: ", len(dampenedReports)+len(safeReports))
}

func parseReports(filename string) ([]Report, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("Error reading reports file: %v", err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    var reportList []Report
    for scanner.Scan() {
        line := scanner.Text()
        levelStrings := strings.Fields(line)
        var levels Report
        for _, levelStr := range levelStrings {
            level, err := strconv.Atoi(levelStr)
	    if err != nil {
	        return nil, fmt.Errorf("Error parsing level '%s' in report: %v", levelStr, err)
	    }
	    levels = append(levels, level)
        }
    reportList = append(reportList, levels)
    }
    return reportList, nil
}

func inspectReport(report Report) (bool) {
    // levels must be all increasing or all decreasing
    levelsIncreasing := true
    levelsDecreasing := true
    for i := 1; i < len(report); i++ {
       	if report[i] > report[i-1] {
	    levelsDecreasing = false
	} else if report[i] < report[i-1] {
	    levelsIncreasing = false
	} else if report[i] == report[i-1] {
	    return false
	}
    }
    // Condition 2: Check if the difference between adjacent levels is > 1 and< 3
    for i := 1; i < len(report); i++ {
        delta := int(math.Abs(float64(report[i] - report[i-1])))
	if delta < 1 || delta > 3 {
	    return false
	} 
    }
    if (levelsIncreasing && levelsDecreasing) {
    }
    return levelsIncreasing || levelsDecreasing
}

func dampenerCheck (report Report) (bool) {
    for i := 0; i < len(report); i++ {
	modifiedReport := append([]int(nil), report...)
	modifiedReport = append(modifiedReport[:i], modifiedReport[i+1:]...)
	if inspectReport(modifiedReport) {
	    return true
	}
    }
    return false
}
