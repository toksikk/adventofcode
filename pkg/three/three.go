package three

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	values := readInput()
	heatmap := createHeatmap(values)
	fmt.Printf("power: %d\n", calcPower(calcGamma(heatmap, len(values)), calcEpsilon(heatmap, len(values))))
	fmt.Printf("oxygen * co2: %d\n", calcOxygen(values)*calcCO2(values))
}

func calcOxygen(values [][]int) int {
	valueStore := make([][][]int, 0)
	valueStore = append(valueStore, values)
	for i := 0; i < len(values[0]); i++ {
		frequencyMap := getMostCommonValues(valueStore[i])
		newValues := make([][]int, 0)
		for _, v := range valueStore[i] {
			if v[i] == frequencyMap[i] {
				newValues = append(newValues, v)
			}
		}

		valueStore = append(valueStore, newValues)

		if len(newValues) == 1 {
			break
		}
	}

	return convertToDecimal(valueStore[len(valueStore)-1][0])
}

func calcCO2(values [][]int) int {
	valueStore := make([][][]int, 0)
	valueStore = append(valueStore, values)
	for i := 0; i < len(values[0]); i++ {
		frequencyMap := getLeastCommonValues(valueStore[i])
		newValues := make([][]int, 0)
		for _, v := range valueStore[i] {
			if v[i] == frequencyMap[i] {
				newValues = append(newValues, v)
			}
		}

		valueStore = append(valueStore, newValues)

		if len(newValues) == 1 {
			break
		}
	}

	return convertToDecimal(valueStore[len(valueStore)-1][0])
}

func getMostCommonValues(values [][]int) []int {
	ones := make([]int, len(values[0]))
	zeros := make([]int, len(values[0]))
	result := make([]int, len(values[0]))
	for _, v := range values {
		for i, b := range v {
			if b == 1 {
				ones[i]++
			} else if b == 0 {
				zeros[i]++
			}
		}
	}
	for k := range result {
		if ones[k] > zeros[k] {
			result[k] = 1
		} else {
			result[k] = 0
		}
		if ones[k] == zeros[k] {
			result[k] = 1
		}
	}
	return result
}

func getLeastCommonValues(values [][]int) []int {
	ones := make([]int, len(values[0]))
	zeros := make([]int, len(values[0]))
	result := make([]int, len(values[0]))
	for _, v := range values {
		for i, b := range v {
			if b == 1 {
				ones[i]++
			} else if b == 0 {
				zeros[i]++
			}
		}
	}
	for k := range result {
		if ones[k] > zeros[k] {
			result[k] = 0
		} else {
			result[k] = 1
		}
		if ones[k] == zeros[k] {
			result[k] = 0
		}
	}
	return result
}

func calcPower(gamma []int, epsilon []int) int {
	return convertToDecimal(gamma) * convertToDecimal(epsilon)
}

func convertToDecimal(gamma []int) int {
	gammaDecimalValue := 0
	for k, v := range gamma {
		switch k {
		case 0:
			if v == 1 {
				gammaDecimalValue += 2048
			}
		case 1:
			if v == 1 {
				gammaDecimalValue += 1024
			}
		case 2:
			if v == 1 {
				gammaDecimalValue += 512
			}
		case 3:
			if v == 1 {
				gammaDecimalValue += 256
			}
		case 4:
			if v == 1 {
				gammaDecimalValue += 128
			}
		case 5:
			if v == 1 {
				gammaDecimalValue += 64
			}
		case 6:
			if v == 1 {
				gammaDecimalValue += 32
			}
		case 7:
			if v == 1 {
				gammaDecimalValue += 16
			}
		case 8:
			if v == 1 {
				gammaDecimalValue += 8
			}
		case 9:
			if v == 1 {
				gammaDecimalValue += 4
			}
		case 10:
			if v == 1 {
				gammaDecimalValue += 2
			}
		case 11:
			if v == 1 {
				gammaDecimalValue++
			}
		}
	}
	return gammaDecimalValue
}
func convertToDecimal5(gamma []int) int {
	gammaDecimalValue := 0
	for k, v := range gamma {
		switch k {
		case 0:
			if v == 1 {
				gammaDecimalValue += 16
			}
		case 1:
			if v == 1 {
				gammaDecimalValue += 8
			}
		case 2:
			if v == 1 {
				gammaDecimalValue += 4
			}
		case 3:
			if v == 1 {
				gammaDecimalValue += 2
			}
		case 4:
			if v == 1 {
				gammaDecimalValue++
			}
		}
	}
	return gammaDecimalValue
}

func calcGamma(heatmap []int, lines int) []int {
	gamma := make([]int, len(heatmap))
	for k, v := range heatmap {
		if v > lines/2 {
			gamma[k] = 1
		} else {
			gamma[k] = 0
		}
	}
	return gamma
}

func calcEpsilon(heatmap []int, lines int) []int {
	gamma := make([]int, len(heatmap))
	for k, v := range heatmap {
		if v < lines/2 {
			gamma[k] = 1
		} else {
			gamma[k] = 0
		}
	}
	return gamma
}

func createHeatmap(values [][]int) []int {
	heatmap := make([]int, len(values[0]))
	for _, v := range values {
		for k2, v2 := range v {
			if v2 == 1 {
				heatmap[k2]++
			}
		}
	}
	return heatmap
}

func readInput() [][]int {
	file, err := os.Open("./pkg/three/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	values := make([][]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		val := make([]int, 0)
		for _, v := range line {
			x, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			val = append(val, x)
		}
		values = append(values, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return values
}
