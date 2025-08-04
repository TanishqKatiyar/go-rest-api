func HandleRequest(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
}

type User struct {
  ID string `json:"id"`
  Name string `json:"name"`
}

type User struct {
  ID string `json:"id"`
  Name string `json:"name"`
}

err := db.Ping()
if err != nil {
  log.Fatal(err)
}

