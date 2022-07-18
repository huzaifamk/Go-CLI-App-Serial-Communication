package cmd

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var open01Cmd = &cobra.Command{
	Use:                   "open-70316363",
	Short:                 "cdrawer open 70316363 - Open Cash Drawer at 1B 70 31 63 63 00 00 00",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		data2, err := hex.DecodeString("70316363")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Conversion to Bytes successful: ")
		fmt.Printf("% x \n", data2)
		fmt.Println("Configuring Port: ", port)
		fmt.Println("The drawer is opening...")

		options := serial.OpenOptions{
			PortName:   "COM1",
			BaudRate:   9600,
			DataBits:   8,
			StopBits:   1,
			ParityMode: serial.PARITY_NONE,
		}

		fmt.Println("The configurations for your drawer are:")
		fmt.Println("Port Name: ", options.PortName)
		fmt.Println("Baud Rate: ", options.BaudRate)
		fmt.Println("Data Bits: ", options.DataBits)
		fmt.Println("Stop Bits: ", options.StopBits)
		fmt.Println("Parity Mode: ", options.ParityMode)

		port, err := serial.Open(options)
		if err != nil {
			log.Fatalf("serial.Open: %v", err)
		}

		defer port.Close()

		b := []byte{0x1b, 0x70, 0x31, 0x63, 0x63, 0x00, 0x00, 0x00}
		n, err := port.Write(b)
		if err != nil {
			log.Fatalf("port.Write: %v", err)
		}

		fmt.Println("Wrote", n, "bytes.")

	},
}

// // init is called after packages' init functions have been called.
// func init() {
// 	rootCmd.AddCommand(open01Cmd)
// 	open01Cmd.PersistentFlags().StringVarP(&port, "port", "p", "COM1", "serial port")
// }
