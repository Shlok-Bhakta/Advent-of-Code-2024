package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main2() {
	fileHandle, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	leftdata := make([]int, 0)
	rightdata := make([]int, 0)
	for scanner.Scan() {
		strng := scanner.Text()
		split := strings.Split(strng, "   ")
		leftint, _ := strconv.Atoi(split[0])
		rightint, _ := strconv.Atoi(split[1])
		leftdata = append(leftdata, leftint)
		rightdata = append(rightdata, rightint)

	}
	sort.Ints(leftdata)
	sort.Ints(rightdata)
	sum := 0
	for i := 0; i < len(leftdata); i++ {
		if leftdata[i] == rightdata[i] {
			continue
		} else if leftdata[i] > rightdata[i] {
			sum += leftdata[i] - rightdata[i]
		} else if leftdata[i] < rightdata[i] {
			sum += rightdata[i] - leftdata[i]
		}
	}
	fmt.Println(sum)
}
