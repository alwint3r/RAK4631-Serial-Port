package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/acarl005/stripansi"
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

func openPort(name string) (serial.Port, error) {
	return serial.Open(name, &serial.Mode{
		BaudRate: 115200,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	})
}

func main() {
	vid := "239A"
	pid := "8029"

	log.Println("Looking for device with VID:", vid, "PID:", pid)

	var portDetails *enumerator.PortDetails
	var err error

	for {
		portDetails, err = findPort(vid, pid)
		if err != nil {
			log.Fatal(err)
		}

		if portDetails != nil {
			break
		}

		time.Sleep(time.Second * 1)
	}

	log.Println("Found port:", portDetails.Name)

	port, err := openPort(portDetails.Name)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Opened port:", portDetails.Name)

	stringBuffer := ""

	for {
		buf := make([]byte, 1024)
		_, err := port.Read(buf)
		if err != nil {
			errMsg := err.Error()
			errMsg = strings.ToLower(errMsg)
			if errMsg == "device not configured" || strings.Contains(errMsg, "port has been closed") || strings.Contains(errMsg, "the device does not recognize the command") {
				newPort, err := openPort(portDetails.Name)
				if err != nil {
					continue
				}
				port = newPort
			} else {
				log.Fatal(err)
			}
		}

		for _, b := range buf {
			if b == 0 {
				continue
			}

			if b == '\n' {
				stringBuffer = strings.ReplaceAll(stringBuffer, "\r", "")
				stringBuffer = strings.ReplaceAll(stringBuffer, "\n", "")
				stringBuffer = stripansi.Strip(stringBuffer)
				fmt.Printf("%s\r\n", stringBuffer)
				stringBuffer = ""
			} else {
				stringBuffer += string(b)
			}
		}
	}
}

// ESC[0;31m E
