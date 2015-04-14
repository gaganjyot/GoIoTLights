package main

import(  "fmt"
         "github.com/gorilla/mux"
         "net/http"
         "github.com/kidoman/embd"
       _ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
         "log"
      )

var light = 1
func Toggle(w http.ResponseWriter, r *http.Request) {
    if light == 1 {
        light = 0
        embd.DigitalWrite(10, embd.High)
        fmt.Fprintf(w, "OFF")
    } else {
        light = 1
        embd.DigitalWrite(10, embd.Low)
        fmt.Fprintf(w, "ON")
    }
}

func main() {
    embd.InitGPIO()
    defer embd.CloseGPIO()
    embd.SetDirection(10, embd.Out)
    r := mux.NewRouter()
    r.HandleFunc("/", Toggle)
    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":8", r))
}
