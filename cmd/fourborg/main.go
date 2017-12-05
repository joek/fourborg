package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/joek/fourborg/gobot/fourborg"
	"github.com/joek/robotwebhandlers/ws"

	"github.com/joek/picoborgrev"
	"github.com/joek/robotwebhandlers/webcam"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	var addr = flag.String("addr", ":8080", "http service address")
	var webcamHost = flag.String("webcamHost", "localhost", "Host of webcam image.")
	var webcamPort = flag.Uint("webcamPort", 8080, "Port of webcam image.")
	var assetPath = flag.String("assetPath", "./assets", "Folder with html assets")

	flag.Parse()

	com := make(chan *ws.BotCommand)
	h := ws.NewHub(com)
	go h.Run()

	r := raspi.NewAdaptor()
	motor := picoborgrev.NewDriver(r, "motor", 0x44)
	borg := fourborg.NewFourBorgDriver(r, "rev", motor)

	work := func() {

		go func() {
			for c := range com {
				// TODO: Input validation
				if c.Motor != nil {
					borg.SetMotorLeft(c.Motor.Left)
					borg.SetMotorRight(c.Motor.Right)
				} else if c.Event == "Disconnect" {
					borg.Halt()
				}
			}
		}()
	}

	robot := gobot.NewRobot("borgbot",
		[]gobot.Connection{r},
		[]gobot.Device{borg},
		work,
	)

	go robot.Start()
	defer robot.Stop()

	webcamURL := fmt.Sprintf("%s:%d", *webcamHost, *webcamPort)
	wh := webcam.NewHandler(
		webcamURL,
	)

	log.Println("Robot started")

	http.HandleFunc("/webcam", func(w http.ResponseWriter, r *http.Request) { wh.Handle(w, r) })
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { h.ServeWs(w, r) })
	http.Handle("/", http.FileServer(http.Dir(*assetPath)))

	log.Println("Start webserver")

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// - Sensor output (broadcast)
}
