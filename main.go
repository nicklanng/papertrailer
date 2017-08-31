package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"

	"github.com/nicklanng/papertrailer/pkg/config"
	"github.com/nicklanng/papertrailer/pkg/events"
)

var knownIssues []string

func isKnownIssue(message string) bool {
	for _, ki := range knownIssues {
		if strings.Contains(message, ki) {
			return true
		}
	}

	return false
}

func main() {

	config.GetConfig()
	knownIssues = config.GetKnownIssues()
	config.CheckState()

	if len(os.Args) < 2 {
		fmt.Println("Usage: papertrailer <group_id>")
		return
	}
	groupID := os.Args[1]

	response := &events.EventResponse{}

	spinnerIndex := 0
	spinner := []string{"|", "/", "-", "\\"}

	var knownIssues int
	lastSeenDate := ""
	lastSeenID := ""

	for {
		knownIssues = 0
		fmt.Print("\r" + spinner[spinnerIndex])
		spinnerIndex = (spinnerIndex + 1) % len(spinner)

		response = events.Fetch(groupID, lastSeenID)
		if len(response.Events) == 0 {
			continue
		}
		if response.MaximumID != "" {
			lastSeenID = response.MaximumID
		}

		fmt.Print("\r")

		for _, evt := range response.Events {
			lastSeenDate = evt.DisplayRecievedAt
			if isKnownIssue(evt.Message) {
				knownIssues++
				continue
			}

			color.Set(color.FgCyan)
			fmt.Print(evt.DisplayRecievedAt + " ")

			color.Set(color.FgYellow)
			fmt.Print(evt.Hostname + " ")

			color.Set(color.FgWhite)
			fmt.Print(evt.Message + "\n")

			color.Unset()
		}

		if knownIssues == 0 {
			continue
		}

		color.Set(color.FgCyan)
		fmt.Print(lastSeenDate + " ")
		color.Set(color.FgGreen)
		fmt.Println("Found " + strconv.Itoa(knownIssues) + " known issues")
		color.Unset()
	}
}
