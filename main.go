package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
)

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatalf("err to open file: %v", err)
	}

	fileMemorize, err := os.OpenFile("memorize.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("err to open file: %v", err)
	}

	clear := confirm("Do you wanna reset words to memorize?")
	if clear == "Yes" {
		os.Truncate("memorize.txt", 0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	clearScreen()
	for scanner.Scan() {

		counter += 1
		fmt.Println(scanner.Text())
		status := confirm("Do you know this word?")
		if status == "No" {
			fileMemorize.WriteString(scanner.Text() + "\n")
			fmt.Println("i don't know this word")

		} else if status == "Yes" {
			fmt.Println("i know this word")
		}
		clearScreen()

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner err: %v", err)
	}
	defer fileMemorize.Close()

}
func clearScreen() {
	//For Linux & Darwin
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func confirm(label string) string {
	prompt := promptui.Select{
		Label: label,
		Items: []string{"Yes", "No", "Cancel"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("prompt failed: %v\n", err)
	}
	if result == "Cancel" {
		os.Exit(0)
	}
	return result
}
