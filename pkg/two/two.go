package two

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	commands := readInput()
	fmt.Printf("1st final position: %d\n", solve(commands, false))
	fmt.Printf("2nd final position: %d\n", solve(commands, true))
}

func solve(commands []string, alt bool) int {
	depth := 0
	horizontalPosition := 0
	aim := 0

	for _, v := range commands {
		command := strings.Split(v, " ")
		switch command[0] {
		case "forward":
			x, _ := strconv.Atoi(command[1])
			horizontalPosition += x
			if alt {
				depth += aim * x
			}
		case "up":
			x, _ := strconv.Atoi(command[1])
			if !alt {
				depth -= x
			} else {
				aim -= x
			}
		case "down":
			x, _ := strconv.Atoi(command[1])
			if !alt {
				depth += x
			} else {
				aim += x
			}
		}
	}
	return horizontalPosition * depth
}

func readInput() []string {
	file, err := os.Open("./pkg/two/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	commands := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return commands
}
