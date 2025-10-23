package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use: "sshelp",
	Short: "Help remember all the ssh connections you have",
	Long: "Helps manage all of the ssh connections you might have and reminds you of the password for that connection",
}



func Run() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}