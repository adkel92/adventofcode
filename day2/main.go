package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Type Scores
const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

// Outcome Scores
const (
	LOSS = 0
	DRAW = 3
	WIN  = 6
)

// Score Count
type score struct {
	userScore      int
	predictedScore int
}

func main() {

	// setup game variables
	var game = map[string]int{
		"A": ROCK,
		"B": PAPER,
		"C": SCISSORS,
		"X": ROCK,
		"Y": PAPER,
		"Z": SCISSORS,
	}

	// import "strategy" guide
	data, err := os.Open("./data.txt")

	// Error handling for file read
	if err != nil {
		log.Fatal(err)
	}

	// Scan data per line
	scanner := bufio.NewScanner(data)

	// add each line in text file to array
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	lines = append(lines)

	// Set users game score
	userScore := 0

	// Loop each array and calculates score
	for _, round := range lines {
		items := strings.Fields(round)

		opponent, user := game[items[0]], game[items[1]]

		// DRAW: X - X = 0
		if user-opponent == 0 {
			userScore += DRAW + user
		}

		// WIN: PAPER(2) - ROCK(1) = 1
		// WIN: SCISSORS(3) - PAPER(2) = 1
		// WIN: ROCK(1) - SCISSORS(3) = -2
		if user-opponent == 1 {
			userScore += WIN + user
		}

		if user-opponent == -2 {
			userScore += WIN + user
		}

		// LOSS: ROCK(1) - PAPER(2) = -1
		// LOSS: PAPER(2) - SCISSORS(3) = -1
		// LOSS: SCISSORS(3) - ROCK(1) = 2

		if user-opponent == -1 {
			userScore += LOSS + user
		}

		if user-opponent == 2 {
			userScore += LOSS + user
		}
	}

	// Set predictedScore
	predictedScore := 0

	// Loop through each line same as before
	for _, round := range lines {
		items := strings.Fields(round)

		opponent, user := game[items[0]], items[1]

		// X means you have to lose
		// -1 of opponenet will always result in a lose
		if user == "X" {
			score := opponent - 1
			if score == 0 {
				score = DRAW
			}
			predictedScore += LOSS + score
		}

		// Y means you need to draw
		if user == "Y" {
			predictedScore += DRAW + opponent
		}

		// Z means you need to WIN
		// +1 on opponent will result in win unless it equals 4
		if user == "Z" {
			score := opponent + 1
			if score == 4 {
				score = 1
			}
			predictedScore += WIN + score
		}
	}

	fmt.Println("Answer 1 =", userScore)
	fmt.Println("Answer 2 =", predictedScore)
}
