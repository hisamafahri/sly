package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/hisamafahri/sly/helper"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// initCmd will initalize Sly app in the repo
var importCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"i"},
	Short:   "Import all public keys",
	Run: func(cmd *cobra.Command, args []string) {

		yfile, err := ioutil.ReadFile("sly.config.yaml")

		if err != nil {
			log.Fatal(err)
		}

		m := make(map[interface{}]interface{})

		err = yaml.Unmarshal([]byte(yfile), &m)

		if err != nil {
			log.Fatalf("error: %v", err)
		}

		server := m["server"].(string)
		emails := m["users"]

		for email := range emails.(map[string]interface{}) {
			email = strings.TrimSpace(email)
			server = strings.TrimSpace(server)
			checkKey, err := helper.RunCommand("gpg --list-keys " + email)

			if err != nil {
				fmt.Println("Key " + email + " doesn't exist in this device. Checking on the " + server + " server...")
				_, err := helper.RunCommand("gpg --keyserver " + server + " --search-keys " + email)
				if err != nil {
					log.Fatalf("Error on fetching public key from the server: %v", err)
				}
				fmt.Println("Public key of " + email + " successfully downloaded!")
			}

			fmt.Println(checkKey)
		}

	},
}
