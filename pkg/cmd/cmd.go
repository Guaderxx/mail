package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Guaderxx/mail/pkg/mail"
	"github.com/Guaderxx/mail/pkg/model"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mail",
	Short: "mail is a simple tool to send email",
	Long: `Just a demo for snap
    Feel free`,
	Run: func(cmd *cobra.Command, args []string) {
		mail.Send()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&model.ConfigFile, "file", "f", "", "config file (default is $HOME/.config/mail/mail.json")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if model.ConfigFile == "" {
		initConfigDir()
		model.InitMailFile()
	}
	if model.UserFile == "" {
		model.InitUserFile()
	}
	model.LoadMailFile()
	model.LoadUserFile()
}

func initConfigDir() {
	homeDir := os.Getenv("HOME")
	// init ~/.config
	defaultConfDir := filepath.Join(homeDir, ".config")
	if _, err := os.Stat(defaultConfDir); os.IsNotExist(err) {
		err1 := os.Mkdir(defaultConfDir, 0777)
		if err1 != nil {
			log.Fatal(err1)
		}

	}
	// init ~/.config/mail
	defaultMailDir := filepath.Join(defaultConfDir, "mail")
	if _, err := os.Stat(defaultMailDir); os.IsNotExist(err) {
		err1 := os.Mkdir(defaultMailDir, 0777)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
}
