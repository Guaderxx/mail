package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	mail "github.com/xhit/go-simple-mail/v2"
)

var (
	tmpUser = `{
    "host": "smtp.gmail.com",
    "port": 587,
    "username": "example@gmail.com",
    "password": "",
    "connect-timeout": 10,
    "send-timeout": 10
}
`
	U User
)

type User struct {
	Host           string `json:"host"`
	Port           int    `json:"port"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	EncryptionType mail.Encryption
	ConnectTimeout int64 `json:"connect-timeout"`
	SendTimeout    int64 `json:"send-timeout"`
}

func InitUserFile() {
	homeDir := os.Getenv("HOME")
	// default user.json
	// init user file
	UserFile = filepath.Join(homeDir, ".config", "mail", "user.json")
	if _, err := os.Stat(UserFile); os.IsNotExist(err) {
		f, err1 := os.Create(UserFile)
		if err1 != nil {
			log.Fatal(err1)
		}
		defer f.Close()
		fmt.Fprintf(f, "%s", tmpUser)
	}
}
func LoadUserFile() {
	tmp, err := ioutil.ReadFile(UserFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(tmp, &U)
	if err != nil {
		log.Fatal(err)
	}
	U.EncryptionType = mail.EncryptionSTARTTLS
}
