package main

import (
	"cdrawer/cmd"
	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

func main() {
	options := serial.OpenOptions{
		PortName:   "COM1",
		BaudRate:   9600,
		DataBits:   8,
		StopBits:   1,
		ParityMode: serial.PARITY_NONE,
	}

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
	cmd.Execute()
}
