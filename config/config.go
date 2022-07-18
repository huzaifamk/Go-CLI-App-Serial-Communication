package config

import (
	"github.com/jacobsa/go-serial/serial"
	"gopkg.in/ini.v1"
)

var cfg, _ = ini.Load("D:/FreelancingProjects/Fiverr/New folder/cdrawer/my.ini")

var P = cfg.Section("").Key("PortName").String()
var B, _ = cfg.Section("").Key("BaudRate").Uint()
var D, _ = cfg.Section("").Key("DataBits").Uint()
var S, _ = cfg.Section("").Key("StopBits").Uint()
var Pm, _ = cfg.Section("").Key("ParityMode").Uint()

func CheckPm(x uint) serial.ParityMode {
	switch Pm {
	case 0:
		return serial.PARITY_NONE
	case 1:
		return serial.PARITY_ODD
	case 2:
		return serial.PARITY_EVEN
	}
	return serial.PARITY_NONE
}

// Here you can edit your configurations, you can also make new configurations and apply it to any of the command, just by changing the name of var <Options> in your cmd files.
var Options = serial.OpenOptions{
	PortName:   P,
	BaudRate:   B,
	DataBits:   D,
	StopBits:   S,
	ParityMode: CheckPm(Pm),
}

//Here is a new sample configuration, you can uncomment and use this if you want to and edit it :D

// var Options2 = serial.OpenOptions{
// 	PortName:   "TEST",
// 	BaudRate:   1234,
// 	DataBits:   8,
// 	StopBits:   1,
// 	ParityMode: serial.PARITY_NONE,
// }
