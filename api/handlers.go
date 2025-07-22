func HandleRequest(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
}

type User struct {
  ID string `json:"id"`
  Name string `json:"name"`
}

