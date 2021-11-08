package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatalf("err to open file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner err: %v", err)
	}

}

func confirm() bool {
	prompt := promptui.Select{
		Label: "Select[Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("prompt failed: %v\n", err)
	}
	return result == "Yes"
}
