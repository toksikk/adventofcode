package two

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var depth int
var horizontalPosition int
var aim int

func Run() {
	commands := readInput()
	fmt.Printf("1st final position: %d\n", solve(commands, false))
	fmt.Printf("2nd final position: %d\n", solve(commands, true))
}

func solve(commands []string, alt bool) int {
	depth = 0
	horizontalPosition = 0
	aim = 0

	for _, v := range commands {
		command := strings.Split(v, " ")
		switch command[0] {
		case "forward":
			i, _ := strconv.Atoi(command[1])
			forward(i, alt)
		case "up":
			i, _ := strconv.Atoi(command[1])
			up(i, alt)
		case "down":
			i, _ := strconv.Atoi(command[1])
			down(i, alt)
		}
	}
	return horizontalPosition * depth
}

func up(d int, alt bool) {
	if !alt {
		depth -= d
	} else {
		aim -= d
	}
}

func down(d int, alt bool) {
	if !alt {
		depth += d
	} else {
		aim += d
	}
}

func forward(d int, alt bool) {
	horizontalPosition += d
	if alt {
		depth += aim * d
	}
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
