package main

import ( 
  "fmt"
  "html"
  "log"
  "net/http"  
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))  
    log.Println("Endpoint: ", html.EscapeString(r.URL.Path), " was hit")
  })
  
  log.Fatal(http.ListenAndServe(":8080", nil))

}
