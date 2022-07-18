package cmd

import (
	helper "cdrawer/helpers"
	"fmt"

	"github.com/spf13/cobra"
)

var port string

var openCmd = &cobra.Command{
	Use:   "open 1234",
	Short: "Open Cash Drawer on specific serial port",
	Run: func(cmd *cobra.Command, args []string) {
		s := helper.HexToBytes(args[0])
		fmt.Println("Conversion to Bytes successful: ", s)
		fmt.Println("Open", port)

	},
}

func init() {
	rootCmd.AddCommand(openCmd)
	openCmd.PersistentFlags().StringVarP(&port, "port", "p", "COM1", "serial port")
}
