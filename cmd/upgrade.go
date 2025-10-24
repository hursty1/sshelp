package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use: "upgrade",
	Short:   "Installs the latest version of the CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		info := version.FromContext(cmd.Context())
		if !info.IsOutdated {
			fmt.Println("sshelp CLI is already up to date.")
			return
		}
		// install the latest version
		command := exec.Command("go", "install", "github.com/hursty1/ssh_tool/cmd/sshelp@latest")
		_, err := command.Output()
		cobra.CheckErr(err)

		// Get the new version info
		command = exec.Command("sshelp", "--version")
		b, err := command.Output()
		cobra.CheckErr(err)
		re := regexp.MustCompile(`v\d+\.\d+\.\d+`)
		version := re.FindString(string(b))
		fmt.Printf("Successfully upgraded to %s!\n", version)
		os.Exit(0)
	},
}


func IsOutdated(current string) bool {
    resp, _ := http.Get("https://api.github.com/repos/yourname/sshelp/releases/latest")
    defer resp.Body.Close()
    var data struct {
        TagName string `json:"tag_name"`
    }
    json.NewDecoder(resp.Body).Decode(&data)
    return data.TagName != current
}