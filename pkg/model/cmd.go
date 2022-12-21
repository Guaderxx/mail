package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	UserFile   = ""
	ConfigFile = ""
	tmpFile    = `{
    "subject": "",
    "to": "",
    "content": ""
}
`
	M Mail
)

type Mail struct {
	Subject string `json:"subject"`
	To      string `json:"to"`
	Content string `json:"content"`
}

func (m Mail) String() string {
	tmp, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return string(tmp)
}
func InitMailFile() {
	homeDir := os.Getenv("HOME")
	// default mail.json
	// init config file
	ConfigFile = filepath.Join(homeDir, ".config", "mail", "mail.json")
	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		f, err1 := os.Create(ConfigFile)
		if err1 != nil {
			log.Fatal(err1)
		}
		defer f.Close()
		fmt.Fprintf(f, "%s", tmpFile)
	}
}

func LoadMailFile() {
	// Load MailFile
	tmp, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(tmp, &M)
	if err != nil {
		log.Fatal(err)
	}
}
