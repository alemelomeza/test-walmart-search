package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var db Storage

func main() {
	// Favicon handler
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Files server
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public"))
	fs := http.FileServer(filesDir)
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Index handler
	http.HandleFunc("/", IndexHandler)

	// Search handker
	http.HandleFunc("/search", SearchHandler)

	// Server start up
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

func getEnvStorage() StorageType {
	switch os.Getenv("ENVIRONMENT") {
	case "test":
		return Memory
	default:
		return Mongo
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
