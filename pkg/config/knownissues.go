package config

import (
	"bufio"
	"log"
	"os"
	"os/user"
	"strings"
)

func GetKnownIssues() (result []string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	knownIssuesPath := usr.HomeDir + "/.config/papertrailer/knownissues"
	if Config.KnownIssuesPath != "" {
		knownIssuesPath = Config.KnownIssuesPath
	}

	file, err := os.Open(knownIssuesPath)
	if err != nil {
		CreateKnownIssues()
		return []string{}
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

func CreateKnownIssues() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Create(usr.HomeDir + "/.config/papertrailer/knownissues")
	if err != nil {
		log.Fatalln(err)
	}
}
