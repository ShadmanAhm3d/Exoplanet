package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/shadmanAhm3d/ExoplanetProject/models"
)


var store = models.NewExoPlanetStore()

func AddExoplanet (w http.ResponseWriter , r *http.Request){

  var exoplanets models.Exoplanet;
  //ye request ki baat horhi hai 
  if err:=json.NewDecoder(r.Body).Decode(&exoplanets);err!=nil{
    json.NewEncoder(w).Encode(map[string]string{"error":"Some Error has occured"})
  }

  
  exoplanets.ID  = uuid.NewString()
  store.Exoplanet[exoplanets.ID]=exoplanets;


  w.Header().Set("Content-type","application/json")
  json.NewEncoder(w).Encode(exoplanets);

}
