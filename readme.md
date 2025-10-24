SSH Utility

# Installing

go install github.com/hursty1/ssh_tool/cmd/sshelp@latest

# Upgrades

git commit -m "add xyz"
git tag v0.2.0 && git push origin v0.2.0
sshelp upgrade

# Github Actions

Set-Content version.txt "v0.1.4" -NoNewline -Encoding utf8 ## Encoding fix

Enable write permissions for workflows

Go to:
GitHub → Settings → Actions → General → Workflow permissions

Select:


bottom-right corner → click the encoding label (like “UTF-8 with BOM”)

choose “Save with Encoding → UTF-8” (no BOM)

make sure the only contents are:


# Usage

Helps manage all of the ssh connections you might have and reminds you of the password for that connection

Usage:
  sshelp [command]

Available Commands:
  add         Add New device to the config file
  completion  Generate the autocompletion script for the specified shell
  delete      Delete a device config
  help        Help about any command
  list        List all of the configured hosts
  select      Select a device to start a ssh connection

Flags:
  -h, --help   help for sshelp

Use "sshelp [command] --help" for more information about a command.




# Notes

This is a tool to assist me with managing my different remote clients (Raspberry Pi's)

It will use a yaml config file to assign a host name, user, password, ipaddress, friendly name

I found that trying to control the connection within go and auto enter the password was overkill so it will preview what the password is prior 
to starting the connection

**
# TODO

Add password hashing (enter a master password to unlock)
Add ability to update device record (sshelp --update pi3) as example would ask for verify / change values