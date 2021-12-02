package two

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var depth int = 0
var horizontalPosition int = 0
var aim int = 0

func Run() {
	commands := readInput()
	for _, v := range commands {
		command := strings.Split(v, " ")
		switch command[0] {
		case "forward":
			i, _ := strconv.Atoi(command[1])
			forward(i, false)
		case "up":
			i, _ := strconv.Atoi(command[1])
			up(i, false)
		case "down":
			i, _ := strconv.Atoi(command[1])
			down(i, false)
		}
	}
	fmt.Printf("1st final position: %d\n", horizontalPosition*depth)

	depth = 0
	horizontalPosition = 0

	for _, v := range commands {
		command := strings.Split(v, " ")
		switch command[0] {
		case "forward":
			i, _ := strconv.Atoi(command[1])
			forward(i, true)
		case "up":
			i, _ := strconv.Atoi(command[1])
			up(i, true)
		case "down":
			i, _ := strconv.Atoi(command[1])
			down(i, true)
		}
	}
	fmt.Printf("2nd final position: %d\n", horizontalPosition*depth)
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
