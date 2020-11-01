package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

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

	// Healthcheck
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("ok")) //nolint
	})

	// Server start up
	log.Fatal(http.ListenAndServe(":"+getEnvVariable("PORT"), nil))
}

func getEnvVariable(key string) string {
	if env := os.Getenv("ENVIRONMENT"); env == "test" || env == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
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
