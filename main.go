package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/nicklanng/papertrailer/pkg/config"
	"github.com/nicklanng/papertrailer/pkg/events"
)

var knownIssues []string

func loadKnownIssues() (result []string) {
	// load known issues
	file, err := os.Open("knownissues.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if txt != "" {
			result = append(result, txt)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func isKnownIssue(message string) bool {
	for _, ki := range knownIssues {
		if strings.Contains(message, ki) {
			return true
		}
	}

	return false
}

func main() {
	groupID := os.Args[1]

	config.GetConfig()
	config.CheckState()

	knownIssues = loadKnownIssues()

	response := &events.EventResponse{}

	for {
		fmt.Print(".")
		response = events.Fetch(groupID, response.MaximumID)

		if len(response.Events) > 0 {
			fmt.Println()
		}

		for _, evt := range response.Events {
			if !isKnownIssue(evt.Message) {
				color.Set(color.Bold)
			}

			color.Set(color.FgCyan)
			fmt.Print(evt.DisplayRecievedAt + " ")

			color.Set(color.FgYellow)
			fmt.Print(evt.Hostname + " ")

			if isKnownIssue(evt.Message) {
				color.Set(color.FgGreen)
				fmt.Print("KNOWN ")
				fmt.Print("\n")
			} else {
				color.Set(color.FgRed)
				fmt.Print("NEW ")
				color.Set(color.FgWhite)
				fmt.Print(evt.Message + "\n")
			}

			color.Unset()
		}
	}
}
