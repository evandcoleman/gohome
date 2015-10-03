package main

import "github.com/go-martini/martini"

func main() {
	// Install HAP devices
	InstallX10Devices()

	// Setup REST API
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.RunOnAddr(":5591")
}
