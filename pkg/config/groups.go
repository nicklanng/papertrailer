package config

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
)

var Groups GroupList = []Group{}

type Group struct {
	Name         string
	ID           string
	FilterString string
}

type GroupList []Group

func GetGroups() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(usr.HomeDir + "/.config/papertrailer/groups.json")
	defer file.Close()
	if err != nil {
		CreateGroups()
		Groups = make(GroupList, 0)
		return
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Groups)
	if err != nil {
		log.Println("Failed to decode groups file")
		log.Fatalf("Error: %s", err)
	}

	return
}

func CreateGroups() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(usr.HomeDir+"/.config/papertrailer/", os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(usr.HomeDir + "/.config/papertrailer/groups.json")
	if err != nil {
		log.Fatalln(err)
	}

	raw, err := json.Marshal(Groups)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = file.Write(raw)
	if err != nil {
		log.Fatalln(err)
	}

	file.Close()
}
