package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)

type WebController struct {
	pin   int
	pinOn bool
}

func newWebController(pinNumber int) *WebController {
	w := new(WebController)
	w.pin = pinNumber
	w.pinOn = false

	return w
}

func (w *WebController) Toggle() (curState bool) {
	if w.pinOn {
		embd.DigitalWrite(w.pin, embd.High)
	} else {
		embd.DigitalWrite(w.pin, embd.Low)
	}

	w.pinOn = !w.pinOn

	return w.pinOn
}

func (wc *WebController) Handle(w http.ResponseWriter, r *http.Request) {
	pinOn := wc.Toggle()

	output := "Device is switched %s"
	if pinOn {
		fmt.Fprintf(w, output, "on")
	} else {
		fmt.Fprintf(w, output, "off")
	}
}

func startHttp() {
	// start web controller on pin 10
	wc := newWebController(10)

	http.HandleFunc("/", wc.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	embd.InitGPIO()
	defer embd.CloseGPIO()
	embd.SetDirection(10, embd.Out)

	startHttp()
}
