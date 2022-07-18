package cmd

import (
	"encoding/hex"
	"fmt"
	"log"

	config "cdrawer/config"

	"github.com/jacobsa/go-serial/serial"
	"github.com/spf13/cobra"
)

var port string

// openCmd represents the open command
var openCmd = &cobra.Command{
<<<<<<< HEAD
	Use:   "open-01020304",
=======
	Use:   "open-1234",
>>>>>>> 6648d9f2dba4b316d4a35d1f4e6a0c8111bad3c9
	Short: "cdrawer open-01020304 - Open Cash Drawer at 0C 01 02 03 04 00 00 00 ",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := hex.DecodeString("01020304")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Conversion to Bytes successful: ")
		fmt.Printf("% x \n", data)
		fmt.Println("Configuring Port: ", port)
		fmt.Println("The drawer is opening...")

		c := config.Options

		fmt.Println("The configurations for your drawer are:")
		fmt.Println("Port Name: ", c.PortName)
		fmt.Println("Baud Rate: ", c.BaudRate)
		fmt.Println("Data Bits: ", c.DataBits)
		fmt.Println("Stop Bits: ", c.StopBits)
		fmt.Println("Parity Mode: ", c.ParityMode)

		port, err := serial.Open(c)
		if err != nil {
			log.Fatalf("serial.Open: %v", err)
		}

		defer port.Close()

		b := []byte{0x0c, 0x01, 0x02, 0x03, 0x04, 0x00, 0x00, 0x00}
		n, err := port.Write(b)
		if err != nil {
			log.Fatalf("port.Write: %v", err)
		}

		fmt.Println("Wrote", n, "bytes.")

	},
}

// init is called after packages' init functions have been called.
func init() {
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(open01Cmd)
	rootCmd.AddCommand(open02Cmd)
	openCmd.PersistentFlags().StringVarP(&port, "port", "p", "COM1", "serial port")
}
