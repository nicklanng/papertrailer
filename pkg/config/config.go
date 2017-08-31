package config

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
)

type Configuration struct {
	Token           string `json:"token"`
	KnownIssuesPath string `json:"knownissuespath"`
}

var Config Configuration

func GetConfig() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(usr.HomeDir + "/.config/papertrailer/config.json")
	if err != nil {
		log.Println(err)
		CreateConfig()
		CreateKnownIssues()
		log.Fatalln("Created new config file, please edit contents!")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Println("Failed to decode configuration file")
		log.Fatalf("Error: %s", err)
	}
}

func CreateConfig() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	var EmptyStruct Configuration

	err = os.MkdirAll(usr.HomeDir+"/.config/papertrailer/", os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(usr.HomeDir + "/.config/papertrailer/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	raw, err := json.Marshal(EmptyStruct)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = file.Write(raw)
	if err != nil {
		log.Fatalln(err)
	}

	file.Close()
}

func CheckState() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if Config.Token == "" {
		log.Fatalln("No token Specified, please edit " + usr.HomeDir + "/.config/papertrailer/config.json")
	}

}
