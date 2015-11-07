package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/parnurzeal/gorequest"
)

const particleAPIBase string = "https://api.spark.io/v1/"

type Core struct {
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
	DeviceID    string `json:"device_id"`
}

func ParticleDevices() []*accessory.Accessory {
	file, err := ioutil.ReadFile("./config/particle.json")
	if err != nil {
		log.Fatalln("Error opening Particle config file", err.Error())
	}

	var cores []Core
	if err = json.Unmarshal(file, &cores); err != nil {
		log.Fatalln("Error parsing Particle config file", err.Error())
	}

	lightBulbs := []interface{}{}

	for _, b := range cores {
		var device Core = b
		log.Printf("Creating Particle accessory \"%v\"...\n", device.Name)

		info := model.Info{
			Name:         device.Name,
			Manufacturer: "Particle",
		}

		light := accessory.NewLightBulb(info)
		light.OnStateChanged(func(on bool) {
			if on {
				callParticleFunction(device, "animate", "1,100,255")
			} else {
				callParticleFunction(device, "setColor", "0,0,0")
			}
		})
		light.OnBrightnessChanged(func(val int) {
			cl := colorful.Hsv(light.GetHue(), light.GetSaturation()/100, float64(light.GetBrightness())/100)
			callParticleFunction(device, "setColor", fmt.Sprintf("%v,%v,%v", int(cl.R*255), int(cl.G*255), int(cl.B*255)))
		})
		light.OnHueChanged(func(val float64) {
			cl := colorful.Hsv(light.GetHue(), light.GetSaturation()/100, float64(light.GetBrightness())/100)
			callParticleFunction(device, "setColor", fmt.Sprintf("%v,%v,%v", int(cl.R*255), int(cl.G*255), int(cl.B*255)))
		})
		light.OnSaturationChanged(func(val float64) {
			cl := colorful.Hsv(light.GetHue(), light.GetSaturation()/100, float64(light.GetBrightness())/100)
			callParticleFunction(device, "setColor", fmt.Sprintf("%v,%v,%v", int(cl.R*255), int(cl.G*255), int(cl.B*255)))
		})

		lightBulbs = append(lightBulbs, light.Accessory)
	}

	accessories := make([]*accessory.Accessory, len(lightBulbs))
	for i, bulb := range lightBulbs {
		accessories[i] = bulb.(*accessory.Accessory)
	}

	return accessories
}

func callParticleFunction(core Core, name string, args string) {
	log.Println(args)
	request := gorequest.New()
	_, body, errs := request.Post(particleAPIBase+"devices/"+core.DeviceID+"/"+name).
		Set("Authorization", "Bearer "+core.AccessToken).
		Send("args=" + args).
		End()
	if errs != nil {
		log.Println(errs)
	} else {
		log.Println(body)
	}
}
