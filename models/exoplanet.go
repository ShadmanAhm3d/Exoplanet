package models

import "errors" // Import errors package

// Define the Exoplanet struct
type Exoplanet struct {
  ID          string   `json:"name"`
  Name        string   `json:"name"`
  Description string   `json:"description"`
  Distance    float64  `json:"distance"`
  Radius      float64  `json:"radius"`
  Type        string   `json:"type"`
  Mass        *float64 `json:"mass,omitempty"` // Correct json tag formatting
}

// Define the ExoplanetStore struct
type ExoplanetStore struct {
  Exoplanet map[string]Exoplanet
}

// NewExoPlanetStore function should return *ExoplanetStore
func NewExoPlanetStore() *ExoplanetStore {
  return &ExoplanetStore{
    Exoplanet: make(map[string]Exoplanet),
  }
}

// Define the error variables
var (
  ErrNotFound = errors.New("exoplanet not found")
  ErrInvalid  = errors.New("invalid exoplanet")
)

