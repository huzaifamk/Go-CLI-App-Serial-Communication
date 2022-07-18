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
	Use:   "open-01020304",
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

		// Here you can change the configrations, just change the c := config.<Options> in the code below, You can use Options2 if you want to or Options3 as well but you also have to make it in config file, just copy paste the code of Options and change the name of <Options> to <Options3>, thats it :)
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
