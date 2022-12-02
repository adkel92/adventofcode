package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Read calorie data
	data, err := os.Open("./data.txt")

	// Error handling for file read
	if err != nil {
		log.Fatal(err)
	}

	// Scan data per line
	scanner := bufio.NewScanner(data)

	// Add all data including whitespace to lines variable
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	lines = append(lines, "")

	// Loop through each line
	current := 0
	total := [3]int{0, 0, 0}
	for _, calorie := range lines {
		// Convert to number
		num, _ := strconv.Atoi(calorie)

		// add current calorie num to current
		current += num

		// If calorie is blank then calculate if this current is higher than total
		if calorie == "" {
			for i := range total {
				if current >= total[i] {
					total[i], current = current, total[i]
				}
			}
			current = 0
		}
	}

	fmt.Println("Answer 1", total[0])
	fmt.Println("Answer 2", total[0]+total[1]+total[2])

}
