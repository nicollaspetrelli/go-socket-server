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
		log.Println("On connection established ", so.Id())

		so.On("joinRoom", func(company string) {
			so.Join(company)
			log.Println("New user connected in room ", company)
		})

		// CORS Fix
		// server.SetAllowRequest(func(req *http.Request) error {
		// 	req.Header.Add("Access-Control-Allow-Origin", "*")
		// 	req.Header.Add("Access-Control-Allow-Credentials", "true")
		// 	req.Header.Add("Access-Control-Allow-Headers", "Content-Type")
		// 	return nil
		// })
	})

	fmt.Println("Server running on localhost: 5001")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.Handle("/socket.io/", server)

	http.HandleFunc("/sendLocation", func(w http.ResponseWriter, r *http.Request) {
		company := r.URL.Query().Get("company")
		if company == "" {
			http.Error(w, "Missing company", http.StatusBadRequest)
			return
		}

		randLat := rand.Float64()*180 - 90   // random number between -90 and 90
		randLong := rand.Float64()*360 - 180 // random number between -180 and 180
		eventName := "location"

		server.BroadcastTo(company, eventName, fmt.Sprintf("Latitude: %f, Longitude: %f", randLat, randLong))
	})

	log.Fatal(http.ListenAndServe(":5001", nil))
}
