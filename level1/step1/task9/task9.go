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

func main() {
	reader := bufio.NewReader(os.Stdin)

	number, err := getFloat(reader)
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}

	var squareRoot float64 = -1
	if number >= 0 {
		squareRoot = math.Sqrt(number)
		fmt.Printf("%0.3f\n", squareRoot)
		return
	}

	fmt.Printf("%0.0f\n", squareRoot)
}
