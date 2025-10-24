package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	version "github.com/hursty1/sshelp/version"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use: "upgrade",
	Short:   "Installs the latest version of the CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		if !IsOutdated(version.Get()) {
            fmt.Println("sshelp is up to date.")
            return
        }
        
        
        // if runtime.GOOS == "windows" {
        //     fmt.Println("run: go install github.com/hursty1/sshelp@latest")
        //     fmt.Println("to upgrade to the latest version.")
        //     return
        // }
        
        fmt.Println("Upgrading sshelp to the latest version...")
        command := exec.Command("go", "install", "github.com/hursty1/sshelp@latest")
        command.Stdout = os.Stdout
        command.Stderr = os.Stderr

        if err := command.Run(); err != nil {
            fmt.Printf("Upgrade failed: %v\n", err)
            return
        }

        fmt.Println("Upgrade completed. Verifying version...")

        command = exec.Command("sshelp", "version")
        b, err := command.Output()
        if err != nil {
            fmt.Println("Error checking new version:", err)
            return
        }

        re := regexp.MustCompile(`v\d+\.\d+\.\d+`)
        v := re.FindString(string(b))
        fmt.Printf("Successfully upgraded to %s!\n", v)
        os.Exit(0)
	},
}


func LatestVersion() (string, error) {
    resp, err := http.Get("https://api.github.com/repos/hursty1/sshelp/tags")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var tags []struct {
        Name string `json:"name"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
        return "", err
    }

    if len(tags) == 0 {
        return "", fmt.Errorf("no tags found")
    }

    // first tag is latest (GitHub sorts by date)
    return tags[0].Name, nil
}

func IsOutdated(current string) bool {
    latest, err := LatestVersion()
    if err != nil {
        fmt.Println("update check failed:", err)
        return false
    }

    current = strings.TrimSpace(current)
    latest = strings.TrimSpace(latest)

    if current == "" {
        fmt.Println("unknown current version, skipping update check")
        return false
    }

    if current == "dev" {
        fmt.Println("development build detected — upgrading to latest release...")
        return true
    }

    if latest != current {
        fmt.Printf("A new version is available: %s → %s\n", current, latest)
        return true
    }

    return false
}