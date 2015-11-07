# gohome

`gohome` is a go app built on top of [HomeControl](https://github.com/brutella/hc) for controlling various devices in your home via HomeKit from and iOS device.

## Overview

`gohome` was built specifically for my home automation setup which makes use of various X10 modules and some RGB LED strips connected to a [Particle Core](https://particle.io). In my setup, `gohome` is run on a Raspberry Pi 2 which is connected to a [CM15A](http://amzn.com/B0034JYZ8W) X10 interface.

**Disclaimer:** This is not meant to be used by others without modification. This project is setup to work specifically with my setup and my devices. Your setup will probably be different. This is just meant to be an example of how to build a Go app that uses HomeControl.

## Getting Started

1. [Install Go](http://golang.org/doc/install)
2. [Setup your Go workspace](http://golang.org/doc/code.html#Organization)
3. Go get this project and change to that directory

        # Clone project
        $ go get github.com/edc1591/gohome
        $ cd $GOPATH/src/github.com/edc1591/gohome
        
        # Install dependencies
        $ go get

4. Copy the `config-sample` directory to `config` and edit the files according to your setup.

5. Build and install the `gohome` binary

        $ go install
        
6. Run it!

        $ gohome
        
7. Pair your devices with a HomeKit app

## TODO

* Queue X10 commands and execute serially with a delay because X10 can't always handle multiple commands in quick succession.

## License

`gohome` is available under the MIT license. See the LICENSE file for more info.