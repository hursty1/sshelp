package cmd

import (
	"context"
	"fmt"
	"main/internal/filemanager"
	"os"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use: "sshelp",
	Short: "Help remember all the ssh connections you have",
	Long: "Helps manage all of the ssh connections you might have and reminds you of the password for that connection",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error{
		cfg, err := filemanager.LoadConfig()
		if err != nil {
			return err
		}
		cmd.SetContext(context.WithValue(cmd.Context(), "config", cfg))
		return nil
	},
}



func Run() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(selectCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}