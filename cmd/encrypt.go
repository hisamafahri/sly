package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// initCmd will initalize Sly app in the repo
var encryptCmd = &cobra.Command{
	Use:     "encrypt",
	Aliases: []string{"e"},
	Short:   "Encrypt files",
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
		fmt.Printf("--- m:\n%v\n\n", m)

	},
}
