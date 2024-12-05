package routes

import (
	"log"
	"net/http"

	"github.com/TorchofFire/uRelay-guild/config"
	"github.com/TorchofFire/uRelay-guild/internal/connections"
	"github.com/gorilla/mux"
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
	router := mux.NewRouter()

	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/guild-info", guildInfo).Methods("GET")
	router.HandleFunc("/channels", channels).Methods("GET")             // TODO: add middleware to require online
	router.HandleFunc("/users", users).Methods("GET")                   // TODO: add middleware to require online
	router.HandleFunc("/profile/{id}", profile).Methods("GET")          // TODO: add middleware to require online
	router.HandleFunc("/text-channel/{id}", textChannel).Methods("GET") // TODO: add middleware to require online
	http.Handle("/", router)

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
