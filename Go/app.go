package main
 
import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)
 
func main() {
  r := mux.NewRouter()
  r.HandleFunc("/api/sets/{name}/cards", func(w http.ResponseWriter, r *http.Request) {    
    vars := mux.Vars(r)
    name := vars["name"]
    
    dao := NewDataAccess()
    result, _ := dao.CardsForSet(name)
    dao.Close()
    
    fmt.Fprint(w, result)
  })
      
  fmt.Println("Listening on port 9292")
  http.Handle("/", r)
  http.ListenAndServe(":9292", nil)
}