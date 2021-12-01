package one

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Run() {
	firstSolution()
	secondSolution()
}

func getNumbers() []int {
	file, err := os.Open("./pkg/one/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	numbers := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return numbers
}

func firstSolution() {
	p := 0
	i := 0

	for k, v := range getNumbers() {
		if k == 0 {
			p = v
			continue
		}
		if p < v {
			i++
		}
		p = v
	}

	fmt.Println("First Solution: " + strconv.Itoa(i))
}

func secondSolution() {
	increasedMeasurements := 0
	a := make([]int, 3)
	b := make([]int, 3)
	numbers := getNumbers()

	for i := 0; i < len(numbers); i++ {
		if i+4 <= len(numbers) {
			a = numbers[i : i+3]
			b = numbers[i+1 : i+4]
			aSum := 0
			bSum := 0
			for _, v := range a {
				aSum += v
			}
			for _, v := range b {
				bSum += v
			}
			if bSum > aSum {
				increasedMeasurements++
			}
		}
	}

	fmt.Println("Second Solution: " + strconv.Itoa(increasedMeasurements))
}
