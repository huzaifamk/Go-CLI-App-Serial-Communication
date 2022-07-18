package config

import "github.com/jacobsa/go-serial/serial"

// Here you can edit your configurations, you can also make new configurations and apply it to any of the command, just by changing the name of var <Options> in your cmd files.
var Options = serial.OpenOptions{
	PortName:   "COM1",
	BaudRate:   9600,
	DataBits:   8,
	StopBits:   1,
	ParityMode: serial.PARITY_NONE,
}

//Here is a new sample configuration, you can uncomment and use this if you want to and edit it :D

// var Options2 = serial.OpenOptions{
// 	PortName:   "TEST",
// 	BaudRate:   1234,
// 	DataBits:   8,
// 	StopBits:   1,
// 	ParityMode: serial.PARITY_NONE,
// }
