package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var mockDatabase = []User{
	{"1", "admin", "admin@example.com", time.Now()},
	{"2", "tanishq", "tanishq@example.com", time.Now()},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mockDatabase)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "API is operational"})
}

func main() {
	http.HandleFunc("/api/v1/users", GetUsers)
	http.HandleFunc("/api/v1/health", HealthCheck)

	log.Println("Go REST API successfully started on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
