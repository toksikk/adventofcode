package four

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	draws, boards := readInput()

	fmt.Printf("1st winning board: %d\n", calculateWinningBoard(findFirstWinningBoard(boards, createPlayboards(boards), draws)))
	fmt.Printf("2nd winning board: %d\n", calculateWinningBoard(findLastWinningBoard(boards, createPlayboards(boards), draws)))
}

func calculateWinningBoard(board [][]int, playboard [][]bool, draw int) int {
	r := 0
	for i, row := range playboard {
		for j, val := range row {
			if !val {
				r += board[i][j]
			}
		}
	}
	return r * draw
}

func isIntInSlice(i int, slice []int) bool {
	for _, v := range slice {
		if i == v {
			return true
		}
	}
	return false
}

func findLastWinningBoard(boards [][][]int, playboard [][][]bool, draws []int) ([][]int, [][]bool, int) {
	winners := make([]int, 0)
	for _, draw := range draws {
		for i, board := range boards {
			for j, row := range board {
				for k, num := range row {
					if num == draw {
						playboard[i][j][k] = true
						if !isIntInSlice(i, winners) && checkIfWinner(playboard[i]) {
							winners = append(winners, i)
							if len(winners) == len(boards) {
								return board, playboard[i], draw
							}
						}
					}
				}
			}
		}
	}
	return nil, nil, 0
}

func findFirstWinningBoard(boards [][][]int, playboard [][][]bool, draws []int) ([][]int, [][]bool, int) {
	for _, draw := range draws {
		for i, board := range boards {
			for j, row := range board {
				for k, num := range row {
					if num == draw {
						playboard[i][j][k] = true
						if checkIfWinner(playboard[i]) {
							return board, playboard[i], draw
						}
					}
				}
			}
		}
	}
	return nil, nil, 0
}

func checkIfWinner(board [][]bool) bool {
	// check horizontals
	var checks int
	for _, row := range board { // rows
		checks = 0
		for _, v := range row {
			if v {
				checks++
			}
		}
		if checks == len(row) {
			return true
		}
	}

	for i := 0; i < len(board[0]); i++ {
		checks = 0
		for j := 0; j < len(board); j++ {
			if board[j][i] == true {
				checks++
			}
		}
		if checks == len(board) {
			return true
		}
	}
	return false
}
func createPlayboards(boards [][][]int) [][][]bool {
	playboard := make([][][]bool, len(boards))

	for i := range playboard {
		playboard[i] = make([][]bool, len(boards[0]))
		for j := range playboard[i] {
			playboard[i][j] = make([]bool, len(boards[0][0]))
		}
	}

	return playboard
}

func readInput() ([]int, [][][]int) {
	file, err := os.Open("./pkg/four/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	boards := make([][][]int, 0)

	scanner := bufio.NewScanner(file)

	// first line are the draws
	scanner.Scan()
	drawsString := scanner.Text()
	draws := make([]int, 0)
	for _, v := range strings.Split(drawsString, ",") {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		draws = append(draws, n)
	}
	scanner.Scan()

	board := *new([][]int)

	for scanner.Scan() {
		if scanner.Text() == "" {
			boards = append(boards, board)
			board = *new([][]int)
			continue
		}
		currentLine := scanner.Text()

		currentLine = strings.ReplaceAll(currentLine, "  ", " ")
		if len(currentLine) > 0 && currentLine[0] == ' ' {
			currentLine = strings.Replace(currentLine, " ", "", 1)
		}

		splitCurrentLine := strings.Split(currentLine, " ")
		row := *new([]int)
		for i := 0; i < len(splitCurrentLine); i++ {
			x, err := strconv.Atoi(splitCurrentLine[i])
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, x)
		}
		board = append(board, row)
	}

	boards = append(boards, board)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return draws, boards
}
