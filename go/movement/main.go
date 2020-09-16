package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := readInput(os.Stdin, os.Stdout)
	numbers := genNumbers(n)
	printNumbers(numbers, os.Stdout)
}

func readInput(input io.Reader, output io.Writer) int {
	reader := bufio.NewReader(input)
	for {
		fmt.Fprint(output, "Введите целое число от 1 до 100: ")
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Fprintln(output, "Введено недопустимое значение")
		}
		input := strings.TrimRight(string(line), "\r\n")
		n64, err := strconv.ParseInt(input, 10, 64)
		if err != nil || n64 < 1 || n64 > 100 {
			fmt.Fprintln(output, "Введено недопустимое значение")
		} else {
			return int(n64)
		}
	}
}

func genNumbers(n int) [][]int {
	if n < 1 {
		return nil
	}
	numbers := make([][]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = make([]int, n)
	}
	vx, vy := 1, 0
	x, y := 0, 0
	for i := 1; i <= n*n; i++ {
		if !isEmpty(numbers, x+vx, y+vy) {
			// change direction
			if vx == 1 && vy == 0 {
				vx, vy = 0, 1 // down
			} else if vx == 0 && vy == 1 {
				vx, vy = -1, 0 // left
			} else if vx == -1 && vy == 0 {
				vx, vy = 0, -1 // up
			} else {
				vx, vy = 1, 0 // right
			}
		}
		numbers[y][x] = i
		x += vx
		y += vy
	}

	return numbers
}

func isEmpty(numbers [][]int, x, y int) bool {
	if y >= len(numbers) || y < 0 || x >= len(numbers[y]) || x < 0 {
		return false
	}

	if numbers[y][x] != 0 {
		return false
	}

	return true
}

func printNumbers(numbers [][]int, output io.Writer) {
	for y := 0; y < len(numbers); y++ {
		for x := 0; x < len(numbers[y]); x++ {
			fmt.Fprintf(output, "%d ", numbers[y][x])
		}
		fmt.Fprintln(output)
	}
}
