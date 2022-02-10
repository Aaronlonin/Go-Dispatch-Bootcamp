package model

// Pokemons type is an alias for a slice of Pokemon.
type Pokemons []Pokemon

// Pokemon struct represents a single Pokemon.
type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
