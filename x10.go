package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
)

type Device struct {
	Name      string
	HouseCode string
	DeviceID  int
	Dimmable  bool
}

var bulbs = []Device{Device{
	Name:      "Bed Lamp",
	HouseCode: "C",
	DeivceID:  10,
	Dimmable:  true,
}, Device{
	Name:      "Corner Lamp",
	HouseCode: "C",
	DeivceID:  8,
	Dimmable:  true,
}, Device{
	Name:      "Couch Lights",
	HouseCode: "C",
	DeivceID:  6,
	Dimmable:  true,
}, Device{
	Name:      "Desk Lights",
	HouseCode: "C",
	DeivceID:  2,
	Dimmable:  true,
}, Device{
	Name:      "Dining Room",
	HouseCode: "C",
	DeivceID:  4,
	Dimmable:  true,
}, Device{
	Name:      "Fireplace Lights",
	HouseCode: "C",
	DeivceID:  1,
	Dimmable:  false,
}, Device{
	Name:      "Lava Lamps",
	HouseCode: "E",
	DeivceID:  1,
	Dimmable:  false,
}, Device{
	Name:      "TV Lights",
	HouseCode: "C",
	DeivceID:  7,
	Dimmable:  true,
}}

func InstallX10Devices() {
	for _, b := range bulbs {
		go func(b Device) {
			log.Printf("Installing X10 device \"%v\"...\n", b.Name)

			info = model.Info{
				Name:         b.Name,
				Manufacturer: "X10",
			}

			light := accessory.NewLightBulb(info)
			light.OnStateChanged(func(on bool) {
				var string action
				if on {
					action = "on"
				} else {
					action = "off"
				}
				var string method
				if b.Dimmable {
					method = "pl"
				} else {
					method = "rf"
				}

				cmd := fmt.Sprintf("%s %s%i %s", method, b.HouseCode, b.DeviceID, action)
				writeCommand(cmd)
			})

			t, err := hap.NewIPTransport("10000000", light.Accessory)
			if err != nil {
				log.Fatal(err)
			}

			t.Start()
		}(b)
	}
}

func writeCommand(cmd string) {
	// TODO: Queue commands and execute serially with a delay
	conn, err := net.Dial("tcp", "localhost:1099")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	_, err := write.WriteString(cmd)
	if err != nil {
		log.Println(err)
	}
}
