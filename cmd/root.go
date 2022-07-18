package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "cdrawer",
	Version:           "1.0.0",
	Short:             "Cash Drawer CLI",
	Long:              `Cash Drawer CLI is a command line interface for the Cash Drawer.`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cdrawer.yaml)")
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	// options := serial.OpenOptions{
	// 	PortName:   "COM1",
	// 	BaudRate:   9600,
	// 	DataBits:   8,
	// 	StopBits:   1,
	// 	ParityMode: serial.PARITY_NONE,
	// }

	// port, err := serial.Open(options)
	// if err != nil {
	// 	log.Fatalf("serial.Open: %v", err)
	// }

	// defer port.Close()

	// b := []byte{0x0c, 0x01, 0x02, 0x03, 0x04, 0x00, 0x00, 0x00}
	// n, err := port.Write(b)
	// if err != nil {
	// 	log.Fatalf("port.Write: %v", err)
	// }

	// fmt.Println("Wrote", n, "bytes.")
}
