package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	filename := "full_input.txt"
	sLeft, sRight := create2ListsFromFile(filename)
	fmt.Print(len(sLeft))
	fmt.Print("\n")
	copy_sRight := make([]int, len(sRight))
	copy(copy_sRight, sRight)
	sumDiffLists := calcSumOfDiff(sLeft, copy_sRight)
	fmt.Print(sumDiffLists)
	fmt.Print("\n")
	similarityScore := calcSimilarityScore(sLeft, sRight)
	fmt.Print(similarityScore)
	fmt.Print("\n")
}
func calcSimilarityScore(sLeft []int, sRight []int) int {
	uniqueMap := countUniqueOccurances(sRight)
	var similarityScore []int
	for _, v := range sLeft {
		val, ok := uniqueMap[v]
		if ok {
			val *= v
			similarityScore = append(similarityScore, val)
		}
	}
	sum := sumSlice(similarityScore)
	return sum
}

func countUniqueOccurances(listToCheck []int) map[int]int {
	m := make(map[int]int)
	for _, v := range listToCheck {
		val, ok := m[v]
		if ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}
	return m
}
func calcSumOfDiff(sLeft []int, sRight []int) int {
	start := time.Now()
	sort.Ints(sLeft)
	sort.Ints(sRight)
	var distance []int

	for _, v := range sLeft {
		smallest_val := sRight[0]
		diff := v - smallest_val
		diff_abs := absValue(diff)
		distance = append(distance, diff_abs)
		sRight = removeByValue(sRight, smallest_val)
	}
	sumDistances := sumSlice(distance)
	elapsed := time.Since(start)
	fmt.Print(elapsed)
	fmt.Print("\n")

	return sumDistances
}

func sumSlice(slice []int) int {
	result := 0
	for _, v := range slice {
		result += v
	}
	return result
}

func absValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func removeByValue(inputSlice []int, valueToRemove int) []int {
	for i, v := range inputSlice {
		if v == valueToRemove {
			return append(inputSlice[:i], inputSlice[i+1:]...)
		}
	}
	return inputSlice
}

func create2ListsFromFile(filename string) (sLeft []int, sRight []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		split := strings.Fields(line)

		leftVal := getIntFromStr(split[0])
		rightVal := getIntFromStr(split[1])

		sLeft = append(sLeft, leftVal)
		sRight = append(sRight, rightVal)
	}
	return
}

func getIntFromStr(s string) (outputInt int) {
	outputInt, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		// ... handle error
		panic(err)
	}
	return
}
