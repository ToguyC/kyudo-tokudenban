package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"

	"tokutenban/config"
	"tokutenban/models"
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
	seedFlag    bool
)

func enableCORS(r *mux.Router) http.Handler {
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins(strings.Split(os.Getenv("CORS_ORIGINS"), ","))

	return handlers.CORS(credentials, methods, origins)(r)
}

func migrateDatabase() {
	db, err := config.DatabaseConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Migrating database...")
	config.MigrateDatabase(db)
}

func seedDatabase() {
	db, err := config.DatabaseConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Seeding database...")
	config.SeedDatabase(db)
}

func main() {
	flag.BoolVar(&migrateFlag, "migrate", false, "Migrate the database")
	flag.BoolVar(&seedFlag, "seed", false, "Seed the database")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if migrateFlag {
		migrateDatabase()
	}
	if seedFlag {
		seedDatabase()
	}
	if migrateFlag || seedFlag {
		return
	}

	r := mux.NewRouter()

	// REST API
	restRouter := r.PathPrefix(apiVersion).Subrouter()
	restRouter.HandleFunc("/hello", helloHandler).Methods(http.MethodGet)

	// WebSocket server
	r.HandleFunc("/ws", wsHandler)

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", enableCORS(r)); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// helloHandler handles the REST API /hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, _ := config.DatabaseConnection()

	var clubs []models.Club
	db.Find(&clubs)

	// for _, c := range clubs {
	// 	db.Delete(&c)
	// }

	json.NewEncoder(w).Encode(clubs)
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
