package service

import (
	"log"

	errz "github.com/Aaronlonin/capstone/errors"
	"github.com/Aaronlonin/capstone/model"
	"github.com/sirupsen/logrus"
)

// PokemonMap is an alias for a map of pokemon.
type PokemonMap map[int]model.Pokemon

// pokemonOrder is an auxiliary function that helps to sort pokemonData by their ID
// as we are using a map we can't ensure the order of the keys is preserved.
var pokemonOrder []int = []int{1, 2}

// db is a sample data to be used in the service as a placeholder for the real data.
var db PokemonMap = map[int]model.Pokemon{
	1: {
		ID:   1,
		Name: "Pikachu",
		Type: "Electric",
	},
	2: {
		ID:   2,
		Name: "Charmander",
		Type: "Fire/Dragon",
	},
}

// PokemonServicestruct implements PokemonService interface.
type PokemonService struct {
	logger *logrus.Logger
	data   PokemonMap
}

// New returns a new PokemonService instance.
func New(em PokemonMap) *PokemonService {
	if em == nil {
		em = db
	}

	return &PokemonService{
		data: em,
	}
}

// Catch returns all caught pokemon.
func (es *PokemonService) Catch() (model.Pokemons, error) {
	log.Println("In service - GetAllEmployees")

	if err := es.dataValidation(); err != nil {
		return nil, err
	}

	// convert data from map to an slice of Pokemon
	pokemons := make(model.Pokemons, 0, len(es.data))

	// preserve the order
	for _, id := range pokemonOrder {
		pokemons = append(pokemons, es.data[id])
	}

	return pokemons, nil
}

// dataValidation is an auxiliary function that checks if the data has beem initialized or if it is empty
// returns a matching ServiceError if any of these conditions are met.
func (es *PokemonService) dataValidation() error {
	log.Println("In service - dataValidation")

	// special handling if data is nil
	if es.data == nil {
		return errz.ErrDataNotInitialized
	}

	// special handling if data is empty
	if len(es.data) == 0 {
		return errz.ErrEmptyData
	}

	return nil
}
