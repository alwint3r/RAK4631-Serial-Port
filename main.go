package main

import (
	"fmt"
	"log"
	"strings"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

func findPort(vid, pid string) (*enumerator.PortDetails, error) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return nil, err
	}

	for _, port := range ports {
		if port.VID == vid && port.PID == pid {
			return port, nil
		}
	}

	return nil, nil
}

func main() {
	vid := "239A"
	pid := "8029"

	portDetails, err := findPort(vid, pid)
	if err != nil {
		log.Fatal(err)
	}

	if portDetails == nil {
		log.Fatal("No port found")
	}

	log.Println("Found port:", portDetails.Name)

	port, err := serial.Open(portDetails.Name, &serial.Mode{
		BaudRate: 115200,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Opened port:", portDetails.Name)

	stringBuffer := ""

	for {
		buf := make([]byte, 1024)
		_, err := port.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		for _, b := range buf {
			if b == '\n' {
				stringBuffer = strings.ReplaceAll(stringBuffer, "\r", "")
				fmt.Println(stringBuffer)
				stringBuffer = ""
			} else {
				stringBuffer += string(b)
			}
		}
	}
}
