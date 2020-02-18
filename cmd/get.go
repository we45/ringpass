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
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

type Secret struct {
	Service string `json:"service"`
	Key     string `json:"key"`
	Value   string `json:"value"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		copy, _ := cmd.Flags().GetBool("copy")
		jsonVal, _ := cmd.Flags().GetBool("json")
		service, _ := cmd.Flags().GetString("service")
		key, _ := cmd.Flags().GetString("key")

		getFromKeyring(service, key, copy, jsonVal)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().BoolP("copy", "c", false, "copy retrieved secret to system clipboard")
	getCmd.Flags().BoolP("json", "j", false, "Render secret as json on stdout")
	getCmd.Flags().StringP("service", "s", "", "Service that needs to be queried from keyring")
	getCmd.Flags().StringP("key", "k", "", "Key that needs to be queried for value to be retrieved")
	getCmd.MarkFlagRequired("service")
	getCmd.MarkFlagRequired("key")
}

func getFromKeyring(service string, key string, copy bool, jsonVal bool) {
	secret, err := keyring.Get(service, key)
	if err != nil {
		log.Fatal("Unable to retrieve value based on service and key.")
	}
	if copy == true {
		clipboard.WriteAll(secret)
		log.Println("Successfully written secret to system clipboard")
		os.Exit(0)
	}

	if jsonVal == true {
		newSecret := Secret{
			Service: service,
			Key:     key,
			Value:   secret,
		}
		secretMap, _ := json.Marshal(newSecret)
		fmt.Println(string(secretMap))
	}
}
