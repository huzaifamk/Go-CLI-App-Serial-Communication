package config

import "github.com/jacobsa/go-serial/serial"

 var Options = serial.OpenOptions{
	PortName:   "COM123",
	BaudRate:   9600,
	DataBits:   8,
	StopBits:   1,
	ParityMode: serial.PARITY_NONE,
}