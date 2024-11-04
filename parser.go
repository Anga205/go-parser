package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func CheckBraces(code string) bool {
	stack := []rune{}
	for _, char := range code {
		if char == '{' {
			stack = append(stack, char)
		} else if char == '}' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func CheckIfStatements(code string) bool {
	lines := strings.Split(code, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "if") {
			parts := strings.Fields(trimmedLine)
			if len(parts) < 2 || parts[1] == "{" {
				return false
			}
		}
	}
	return true
}

func CheckWhileStatements(code string) bool {
	lines := strings.Split(code, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "while") {
			parts := strings.Fields(trimmedLine)
			if len(parts) < 2 || parts[1] == "{" {
				return false
			}
		}
	}
	return true
}

func CheckForStatements(code string) bool {
	lines := strings.Split(code, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "for") {
			parts := strings.Split(trimmedLine[3:], ";")
			if len(parts) != 3 {
				return false
			}
		}
	}
	return true
}

func main() {
	codeBytes, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	code := string(codeBytes)
	var (
		braceMatch, ifStatementsCorrect, whileStatementsCorrect, forStatementsCorrect bool
		doneCh                                                                        = make(chan bool, 4)
	)

	start := time.Now()

	go func() {
		braceMatch = CheckBraces(code)
		if braceMatch {
			log.Println("Braces matched successfully.")
		} else {
			log.Println("Braces do not match, exiting code.")
		}
		doneCh <- braceMatch
	}()

	go func() {
		ifStatementsCorrect = CheckIfStatements(code)
		if ifStatementsCorrect {
			log.Println("'If' statements are correct.")
		} else {
			log.Println("'If' statements are incorrect, exiting code.")
		}
		doneCh <- ifStatementsCorrect
	}()

	go func() {
		whileStatementsCorrect = CheckWhileStatements(code)
		if whileStatementsCorrect {
			log.Println("'While' statements are correct.")
		} else {
			log.Println("'While' statements are incorrect, exiting code.")
		}
		doneCh <- whileStatementsCorrect
	}()

	go func() {
		forStatementsCorrect = CheckForStatements(code)
		if forStatementsCorrect {
			log.Println("'For' statements are correct.")
		} else {
			log.Println("'For' statements are incorrect, exiting code.")
		}
		doneCh <- forStatementsCorrect
	}()

	// Wait for all goroutines to finish
	for i := 0; i < 4; i++ {
		if !<-doneCh {
			elapsed := time.Since(start)
			log.Println("Finished checks in", elapsed)
			fmt.Println("--------------------------------")
			fmt.Println("Code is invalid.")
			return
		}
	}

	elapsed := time.Since(start)
	log.Println("Finished checks in", elapsed)
	fmt.Println("--------------------------------")
	fmt.Println("Code is valid.")
}
