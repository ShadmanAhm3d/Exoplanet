package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shadmanAhm3d/ExoplanetProject/handlers"
)


func main(){
  fmt.Println("Started Application on :8080")

  r := mux.NewRouter();

  r.HandleFunc("/exoplanet", handlers.AddExoplanet).Methods("POST");
  
  log.Fatal(http.ListenAndServe(":8080", r));

  







  
}
