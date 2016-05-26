package main

import (
	"log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/go-martini/martini"
)

func main() {
	// Install HAP devices
	x10Accessories := X10Devices()
	particleAccessories := ParticleDevices()

	accessories := append(x10Accessories, particleAccessories...)

	bridge := accessory.NewLightbulb(accessory.Info{
		Name:         "Bridge",
		Manufacturer: "Evan",
	})
	config := hc.Config{Pin: "10000000"}
	t, err := hc.NewIPTransport(config, bridge.Accessory, accessories...)
	if err != nil {
		log.Fatal(err)
	}

	t.Start()

	// Setup REST API
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.RunOnAddr(":5591")
}
