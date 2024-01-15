package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal("error establishing new socketio server")
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("On connection established")

		// server.SetAllowRequest(func(req *http.Request) error {
		// 	req.Header.Add("Access-Control-Allow-Origin", "*")
		// 	req.Header.Add("Access-Control-Allow-Credentials", "true")
		// 	req.Header.Add("Access-Control-Allow-Headers", "Content-Type")
		// 	return nil
		// })

		so.Join("realtime") // join the room with the above socketio connection

		so.On("location", func(msg string) {
			log.Println("messge received: " + msg)
		})
	})

	fmt.Println("Server running on localhost: 5001")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.Handle("/socket.io/", server)

	http.HandleFunc("/sendLocation", func(w http.ResponseWriter, r *http.Request) {
		randLat := rand.Float64()*180 - 90   // generates a random number between -90 and 90
		randLong := rand.Float64()*360 - 180 // generates a random number between -180 and 180

		server.BroadcastTo("realtime", "location", fmt.Sprintf("%f,%f", randLat, randLong))
		fmt.Printf("Sent location to room 'realtime'\n")
	})

	log.Fatal(http.ListenAndServe(":5001", nil))
}
