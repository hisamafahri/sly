package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// initCmd will initalize Sly app in the repo
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Initialize sly",
	Run: func(cmd *cobra.Command, args []string) {

		// create the sly.config.yml
		err := ioutil.WriteFile("sly.config.yml", []byte(`# Make sure you have already configure a GPG key
# that published in https://keys.openpgp.org/. If you haven't,
# please read the docs first.

# user's hierarchy groups to manage level of access for the protected files
# the first group in the group list will be the DEFAULT GROUP
groups:
	# - dev:
		# - .env.dev
		# - folder/secret.txt

	# if you want a group to have all of the files that another group have just assign it with a 'group.' prefix. 
	# example: '- group.anotherGroupName'.
	# - prod:
		# - group.local
		# - .env.prod

# list of users and its incorporated group
users:
	# - user1@email.com: ['dev']
	# - user2@email.com: ['dev', prod]
`), 0755)

		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
	},
}
