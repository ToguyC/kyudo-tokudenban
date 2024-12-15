package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"

	"tokutenban/config"
)

// API version prefix
const apiVersion = "/api/v1"

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	migrateFlag bool
)

func main() {
	flag.BoolVar(&migrateFlag, "migrate", false, "Migrate the database")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if migrateFlag {
		db, err := config.DatabaseConnection()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		log.Println("Migrating database...")
		config.MigrateDatabase(db)
		return
	}

	r := mux.NewRouter()

	// REST API
	restRouter := r.PathPrefix(apiVersion).Subrouter()
	restRouter.HandleFunc("/hello", helloHandler).Methods(http.MethodGet)

	// WebSocket server
	r.HandleFunc("/ws", wsHandler)

	log.Println("Server is running on :8080")

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:5173", "http://127.0.0.1:5173"})

	if err := http.ListenAndServe(":8080", handlers.CORS(credentials, methods, origins)(r)); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// helloHandler handles the REST API /hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// wsHandler handles WebSocket connections
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		log.Printf("Received: %s", message)
		log.Printf("Message type: %d", messageType)

		// Echo the message back to the client
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
	}
}
