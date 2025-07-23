package main
import "fmt"
func main() { fmt.Println("Server starting...") }

func HandleRequest(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
}

