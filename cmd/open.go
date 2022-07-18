package cmd

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
	"github.com/spf13/cobra"
)

var port string

var openCmd = &cobra.Command{
	Use:   "open 1234",
	Short: "Open Cash Drawer on specific serial port",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := hex.DecodeString("01020304")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Conversion to Bytes successful: ")
		fmt.Printf("% x \n", data)
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

		b := []byte{0x0c, 0x01, 0x02, 0x03, 0x04, 0x00, 0x00, 0x00}
		n, err := port.Write(b)
		if err != nil {
			log.Fatalf("port.Write: %v", err)
		}

		fmt.Println("Wrote", n, "bytes.")

	},
}

func init() {
	rootCmd.AddCommand(openCmd)
	openCmd.PersistentFlags().StringVarP(&port, "port", "p", "COM1", "serial port")
}
