package main

import (
	"log"

	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
	"github.com/go-martini/martini"
)

func main() {
	// Install HAP devices
	x10Accessories := X10Devices()
	particleAccessories := ParticleDevices()

	accessories := append(x10Accessories, particleAccessories...)

	bridge := accessory.NewLightBulb(model.Info{
		Name:         "Bridge",
		Manufacturer: "Evan",
	})
	t, err := hap.NewIPTransport("10000000", bridge.Accessory, accessories...)
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
