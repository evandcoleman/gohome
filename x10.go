package main

import (
	"fmt"
	"log"
	"net"

	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
)

type Device struct {
	Name      string
	HouseCode string
	DeviceID  int
	Dimmable  bool
}

// TODO: Read this from a config file
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

func X10Devices() []*accessory.Accessory {
	lightBulbs := []interface{}{}

	for _, b := range bulbs {
		var device Device = b
		log.Printf("Creating X10 accessory \"%v\"...\n", device.Name)

		info := model.Info{
			Name:         device.Name,
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
			if device.Dimmable {
				method = "pl"
			} else {
				method = "rf"
			}

			cmd := fmt.Sprintf("%s %s%v %s", method, device.HouseCode, device.DeviceID, action)
			writeCommand(cmd)
		})

		lightBulbs = append(lightBulbs, light.Accessory)
	}

	accessories := make([]*accessory.Accessory, len(lightBulbs))
	for i, bulb := range lightBulbs {
		accessories[i] = bulb.(*accessory.Accessory)
	}

	return accessories
}

func writeCommand(cmd string) {
	// TODO: Queue commands and execute serially with a delay
	log.Printf("Writing command \"%s\"", cmd)
	conn, err := net.Dial("tcp", "localhost:1099")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(cmd + "\n"))
	if err != nil {
		log.Println(err)
	}
}
