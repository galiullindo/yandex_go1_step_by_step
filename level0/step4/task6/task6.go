package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func getFloat(r *bufio.Reader) (float64, error) {
	line, err := r.ReadString('\n')

	if err != nil && err != io.EOF {
		return 0, err
	}

	number, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func check(err error) {
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	x, err := getFloat(reader)
	check(err)

	y, err := getFloat(reader)
	check(err)

	z, err := getFloat(reader)
	check(err)

	m, err := getFloat(reader)
	check(err)

	n, err := getFloat(reader)
	check(err)

	value := (5 * x * math.Sin(2*y)) / (z + math.Pow(n, math.Log(m)))
	fmt.Println(value)
}
