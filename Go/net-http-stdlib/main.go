package main

// Import all the required packages
import (
	"fmt"
	"log"
	"net/http"
)

type server struct{}

// Our server's default Body
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status_code":"200","desc":"server is up!"}`))
}

// Our main function where we'll point our http handler funcs
func main() {
	s := &server{}
	http.Handle("/", s)
  
  // log.Fatal will end the program with a panic if http.ListenAndServe can not start the server.
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)) // here 8080 is our port eg: localhost:8080
}
