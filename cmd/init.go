package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/hisamafahri/sly/helper"
	"github.com/spf13/cobra"
)

// initCmd will initalize Sly app in the repo
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize sly",
	Run: func(cmd *cobra.Command, args []string) {

		// create the sly.config.yml
		err := ioutil.WriteFile("sly.config.yaml", []byte(helper.DefaultConfig), 0755)

		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
	},
}
