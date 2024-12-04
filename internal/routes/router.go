package routes

import (
	"log"
	"net/http"

	"github.com/TorchofFire/uRelay-guild/config"
	"github.com/TorchofFire/uRelay-guild/internal/connections"
	"github.com/gorilla/websocket"
)

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	if websocket.IsWebSocketUpgrade(request) {
		connections.Handler(writer, request)
		return
	}
	fs := http.FileServer(http.Dir("./public"))
	http.StripPrefix("/", fs).ServeHTTP(writer, request)
}

func Init() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/channels", channels)
	http.HandleFunc("/users", users)
	http.HandleFunc("/profile/{id}", profile)

	if config.SecureProtocol {
		if config.CertPath == "" {
			log.Fatal("SecureProtocol is enabled, but CertPath is not set.")
		}
		log.Printf("Starting HTTPS server on https://%s", config.ServerID)
		if err := http.ListenAndServeTLS(config.ServerID, config.CertPath+"/cert.pem", config.CertPath+"/key.pem", nil); err != nil {
			log.Fatal("Failed to start HTTPS server:", err)
		}
	} else {
		log.Printf("Starting HTTP server on http://%s", config.ServerID)
		if err := http.ListenAndServe(config.ServerID, nil); err != nil {
			log.Fatal("Failed to start HTTP server:", err)
		}
	}
}
