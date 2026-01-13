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

func main() {
	var minHeight float64
	var isAssigned bool = false

	reader := bufio.NewReader(os.Stdin)

	for range 3 {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("ошибка чтения: %s\n", err)
		}

		height, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if err != nil {
			fmt.Printf("ошибка преобразования: %s\n", err)
		}

		if !isAssigned {
			minHeight = height
			isAssigned = true
		} else {
			minHeight = math.Min(minHeight, height)
		}
	}

	fmt.Println(minHeight)
}
