package main

import (
	"log"

	"github.com/brutella/hc/hap"
	"github.com/go-martini/martini"
)

func main() {
	// Install HAP devices
	x10Accessories := X10Devices()

	accessories := x10Accessories

	t, err := hap.NewIPTransport("10000000", accessories[0], accessories[1:]...)
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
