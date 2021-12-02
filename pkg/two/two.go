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

		x, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		switch command[0] {
		case "forward":
			horizontalPosition += x
			if alt {
				depth += aim * x
			}
		case "up":
			if !alt {
				depth -= x
			} else {
				aim -= x
			}
		case "down":
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
