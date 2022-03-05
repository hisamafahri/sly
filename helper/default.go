package helper

var DefaultConfig = `---
---
# Make sure you have already configure a GPG key
# that published in https://keys.openpgp.org/. If you haven't,
# please read the docs first.

# Key server for your gpg keys.
# Re-publish ALL of your keys if you change your key server (not recommended).
server: keys.openpgp.org

# user's hierarchy groups to manage level of access for the protected files
# the first group in the group list will be the DEFAULT GROUP
groups:
	dev:
		- .env.dev
		- folder/secret.txt
	staging:
		- .env.stage
		- folder/another-secret.txt

	# if you want a group to have all of the files that another group have just assign it with a 'group.' prefix. 
	# example: '- group.anotherGroupName'.
	prod:
		- group.staging
		- .env.prod

# list of users and its incorporated group
users:
 	user1@email.com: ['dev']
 	user2@email.com: ['dev', prod]

`
