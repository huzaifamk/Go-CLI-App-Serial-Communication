package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var port string

var openCmd = &cobra.Command{
	Use: "open 1-2-3 ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("open", port, args)

	},
}

func init() {
	rootCmd.AddCommand(openCmd)
	openCmd.PersistentFlags().StringVarP(&port, "port", "p", "COM1", "serial port")
}
