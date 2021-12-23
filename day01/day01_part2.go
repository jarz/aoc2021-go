package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	incrementCounter := 0
	first, _ := getNumber(scanner)
	second, _ := getNumber(scanner)
	third, _ := getNumber(scanner)
	sum := first + second + third
	printSum(first, second, third, sum)

	for {
		first = second
		second = third
		third, _ = getNumber(scanner)

		if third == -1 {
			break
		}

		newSum := first + second + third
		printSum(first, second, third, newSum)

		if newSum > sum {
			incrementCounter++
		}
		sum = newSum
	}

	fmt.Println(incrementCounter)
}

func getNumber(scanner *bufio.Scanner) (int, error) {
	if ok := scanner.Scan(); !ok {
		return -1, fmt.Errorf("not ok")
	}

	line := scanner.Text()
	out, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	return out, nil
}

func printSum(first, second, third, sum int) {
	fmt.Printf("%d + %d + %d = %d\n", first, second, third, sum)
}
