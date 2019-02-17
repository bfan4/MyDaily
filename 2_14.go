package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnectr struct {
	name string
}

func (pc PhoneConnectr) Name() string {
	return pc.name
}

func (pc *PhoneConnectr) Rename(name string) {
	pc.name = name
	return
}

func (pc PhoneConnectr) Connect() {
	fmt.Println("Connect:", pc.name)
}

func Disconnect(usb USB) {
	if pc, ok := usb.(PhoneConnectr); ok {
		fmt.Println("Disconnected:", pc.name)
		return
	}
	fmt.Println("unknown disconnected")

}

func main() {
	var a USB
	a = PhoneConnectr{"IPhone X"}
	a.Connect()
	Disconnect(a)

}
