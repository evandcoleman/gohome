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
	DeviceID:  10,
	Dimmable:  true,
}, Device{
	Name:      "Corner Lamp",
	HouseCode: "C",
	DeviceID:  8,
	Dimmable:  true,
}, Device{
	Name:      "Couch Lights",
	HouseCode: "C",
	DeviceID:  6,
	Dimmable:  true,
}, Device{
	Name:      "Desk Lights",
	HouseCode: "C",
	DeviceID:  2,
	Dimmable:  true,
}, Device{
	Name:      "Dining Room",
	HouseCode: "C",
	DeviceID:  4,
	Dimmable:  true,
}, Device{
	Name:      "Fireplace Lights",
	HouseCode: "C",
	DeviceID:  1,
	Dimmable:  false,
}, Device{
	Name:      "Lava Lamps",
	HouseCode: "E",
	DeviceID:  1,
	Dimmable:  false,
}, Device{
	Name:      "TV Lights",
	HouseCode: "C",
	DeviceID:  7,
	Dimmable:  true,
}}

func InstallX10Devices() {
	lightBulbs := []interface{}{}

	for _, b := range bulbs {
		log.Printf("Installing X10 device \"%v\"...\n", b.Name)

		info := model.Info{
			Name:         b.Name,
			Manufacturer: "X10",
		}

		light := accessory.NewLightBulb(info)
		light.OnStateChanged(func(on bool) {
			var action string
			if on {
				action = "on"
			} else {
				action = "off"
			}
			var method string
			if b.Dimmable {
				method = "pl"
			} else {
				method = "rf"
			}

			cmd := fmt.Sprintf("%s %s%i %s", method, b.HouseCode, b.DeviceID, action)
			writeCommand(cmd)
		})

		lightBulbs = append(lightBulbs, light.Accessory)
	}

	accessories := make([]*accessory.Accessory, len(lightBulbs))
	for i, bulb := range lightBulbs {
		accessories[i] = bulb.(*accessory.Accessory)
	}

	t, err := hap.NewIPTransport("10000000", accessories[0], accessories[1:]...)
	if err != nil {
		log.Fatal(err)
	}

	t.Start()
}

func writeCommand(cmd string) {
	// TODO: Queue commands and execute serially with a delay
	conn, err := net.Dial("tcp", "localhost:1099")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	_, err = writer.WriteString(cmd)
	if err != nil {
		log.Println(err)
	}
}
