/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
	"golang.org/x/crypto/ssh/terminal"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		service, _ := cmd.Flags().GetString("service")
		value, _ := cmd.Flags().GetString("value")

		setKeyRingValue(key, value, service)

	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringP("key", "k", "", "Key for the secret to be set")
	setCmd.Flags().StringP("service", "s", "", "service for which you are setting the key")
	setCmd.Flags().StringP("value", "v", "", "Value for the key that you are storing in the keyring")
	setCmd.MarkFlagRequired("key")
	setCmd.MarkFlagRequired("service")
}

func getPasswordFromStdin() string {
	fmt.Print("Enter Value: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal("Unable to read value from stdin")
	}
	return strings.TrimSpace(string(bytePassword))
}

func setKeyRingValue(key string, value string, service string) {
	if value == "" {
		value = getPasswordFromStdin()
	}
	err := keyring.Set(service, key, value)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully stored Key, Value and Service in Keyring")
}
