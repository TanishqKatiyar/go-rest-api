package main
import "fmt"
func main() { fmt.Println("Server starting...") }

func HandleRequest(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
}

err := db.Ping()
if err != nil {
  log.Fatal(err)
}

type User struct {
  ID string `json:"id"`
  Name string `json:"name"`
}

