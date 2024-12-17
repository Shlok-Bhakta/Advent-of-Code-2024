package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fileHandle, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	leftdata := make([]int, 0)
	rightdata := make(map[int]int)
	for scanner.Scan() {
		strng := scanner.Text()
		split := strings.Split(strng, "   ")
		leftint, _ := strconv.Atoi(split[0])
		rightint, _ := strconv.Atoi(split[1])
		leftdata = append(leftdata, leftint)
		rightdata[rightint]++
	}
	sort.Ints(leftdata)
	sum := 0
	for i := 0; i < len(leftdata); i++ {
		sum += leftdata[i] * rightdata[leftdata[i]]
	}
	fmt.Println(sum)
}
