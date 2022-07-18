package main

import (
	"cdrawer/cmd"
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("D:/FreelancingProjects/Fiverr/New folder/cdrawer/my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	P := cfg.Section("").Key("PortName").String()
	B, _ := cfg.Section("").Key("BaudRate").Uint()
	D, _ := cfg.Section("").Key("DataBits").Uint()
	S, _ := cfg.Section("").Key("StopBits").Uint()
	fmt.Println("PortName:", P, "BaudRate:", B, "DataBits:", D, "StopBits:", S)
	cmd.Execute()
}
