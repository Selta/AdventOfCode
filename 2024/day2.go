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
    var safeReports []Report
    for _, report := range reports {
        if inspectReport(report) {
	    safeReports = append(safeReports, report)
	}
    }
	fmt.Println("Safe Reports:", len(safeReports))
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

func inspectReport(report Report) bool {
    // levels are either increasing or decreasing
    isIncreasing := true
    isDecreasing := true
    isEqual := true
    for i := 1; i < len(report); i++ {
       	if report[i] > report[i-1] {
	    isDecreasing = false
	} else if report[i] < report[i-1] {
	    isIncreasing = false
	} else if report[i] == report[i-1] {
	    isEqual = true
	}
    }
    // If not all increasing nor all decreasing, it's not safe
    if !isIncreasing && !isDecreasing && isEqual  {
       	return false
    } else {
    }
    // Condition 2: Check if the difference between adjacent levels is > 1 and< 3
    for i := 1; i < len(report); i++ {
        delta := int(math.Abs(float64(report[i] - report[i-1])))
	if delta < 1 || delta > 3 {
	    return false
	} 
    }
    // If both conditions are satisfied, the report is safe
    return true
}