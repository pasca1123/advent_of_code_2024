package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Provide file name")
	// fmt.Println("---------------------")
	// text, _ := reader.ReadString('\n')
	// text = strings.Replace(text, "\n", "", -1)
	startTime := time.Now()
	text := "full_input.txt"
	reportList := createListOfReports(text)
	correctReports := checkReports(reportList)
	elapsed := time.Since(startTime)
	fmt.Print("Correct reports: ")
	fmt.Print(correctReports)
	fmt.Print("\nTime: ")
	fmt.Print(elapsed)
	fmt.Print("\n")
}

func checkReports(reportList [][]string) int {
	var correctReports int = 0
	for _, report := range reportList {
		if isReportCorrect(report) {
			correctReports += 1
		} else if checkAdjustedReports(report) {
			correctReports += 1
		} else {
			fmt.Print("Incorrect")
		}
	}
	return correctReports
}

func checkAdjustedReports(report []string) bool {
	var correctReport bool
	for j, _ := range report {
		fmt.Print(report)
		fmt.Print("\n")
		copy_report := make([]string, len(report))
		copy(copy_report, report)
		smallerReport := append(copy_report[:j], copy_report[j+1:]...)
		fmt.Print(smallerReport)
		fmt.Print("\n")
		correctReport = isReportCorrect(smallerReport)
		fmt.Print(correctReport)
		fmt.Print("\n")
		if correctReport {
			return correctReport
		}
	}
	return correctReport
}

func isReportCorrect(report []string) bool {
	var correctReport bool = true
	var decInc []bool
	for i, _ := range report {
		if i == len(report)-1 {
			break
		}
		currentVal := convertStrToInt(report[i])
		nextVal := convertStrToInt(report[i+1])

		levelDiff := nextVal - currentVal
		absDiff := absValue(levelDiff)
		if 1 < absDiff && absDiff > 3 {
			correctReport = false
			break
		}

		if currentVal == nextVal {
			correctReport = false
			break
		} else if currentVal < nextVal {
			decInc = append(decInc, true)
		} else {
			decInc = append(decInc, false)
		}

	}
	allDec := checkAllDecInc(decInc)
	if correctReport && allDec {
		return true
	}
	return false
}

func checkAllDecInc(decInc []bool) bool {
	allTrue := true
	allFalse := true

	for _, val := range decInc {
		if val {
			allFalse = false
		} else {
			allTrue = false
		}
		if !allFalse && !allTrue {
			return false
		}
	}
	return true
}

func convertStrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Print(err)
	}
	return num
}

func absValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func createListOfReports(filename string) (reportList [][]string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		var report = strings.Split(line, " ")
		reportList = append(reportList, report)
	}
	return
}
