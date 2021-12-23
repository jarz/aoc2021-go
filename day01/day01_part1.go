package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	last := 0
	incrementCounter := 0

	for {
		line, err := reader.ReadString('\n')

		if err != nil || len(line) == 0 {
			break
		}

		line = strings.TrimSuffix(line, "\n")

		depth, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if depth > last {
			incrementCounter++
		}

		last = depth
	}

	fmt.Println(incrementCounter)
}
