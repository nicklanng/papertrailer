package config

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
)

func Save(token, knownIssuesPath string) {
  Config.Token = token
  Config.KnownIssuesPath = knownIssuesPath

  usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

  err = os.MkdirAll(usr.HomeDir+"/.config/papertrailer/", os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(usr.HomeDir + "/.config/papertrailer/config.json")
	if err != nil {
		log.Fatalln(err)
	}

  raw, err := json.Marshal(Config)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = file.Write(raw)
	if err != nil {
		log.Fatalln(err)
	}

	file.Close()
}
