package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/brutella/hc/accessory"
)

type Device struct {
	Name      string `json:"name"`
	HouseCode string `json:"house_code"`
	DeviceID  int    `json:"device_id"`
	Dimmable  bool   `json:"dimmable"`
}

func X10Devices() []*accessory.Accessory {
	file, err := ioutil.ReadFile("./config/x10.json")
	if err != nil {
		log.Fatalln("Error opening X10 config file", err.Error())
	}

	var bulbs []Device
	if err = json.Unmarshal(file, &bulbs); err != nil {
		log.Fatalln("Error parsing X10 config file", err.Error())
	}

	lightBulbs := []*accessory.Accessory{}

	for _, b := range bulbs {
		var device Device = b
		log.Printf("Creating X10 accessory \"%v\"...\n", device.Name)

		info := accessory.Info{
			Name:         device.Name,
			Manufacturer: "X10",
		}

		light := accessory.NewLightbulb(info)
		light.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
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
		light.Lightbulb.Brightness.OnValueRemoteUpdate(func(val int) {
			if device.Dimmable {
				dimVal := int((float64(val) / 100) * 70)
				cmd := fmt.Sprintf("pl %s%v xdim %v", device.HouseCode, device.DeviceID, dimVal)
				writeCommand(cmd)
			}
		})

		lightBulbs = append(lightBulbs, light.Accessory)
	}

	return lightBulbs
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
