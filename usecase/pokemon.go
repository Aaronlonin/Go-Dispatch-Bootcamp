package usecase

import (
	"fmt"
	"log"

	"github.com/Aaronlonin/capstone/model"
)

// PokemonService is the interface that wraps the service's methods
// Catch
type PokemonService interface {
	Catch() (model.Pokemons, error)
}

// PokemonUSecase implements PokemonService interface.
type PokemonUsecase struct {
	service PokemonService
}

// New returns a new PokemonUsecase instance.
func New(s PokemonService) *PokemonUsecase {
	log.Println("In usecase - NewPokemonUsecase")

	return &PokemonUsecase{
		service: s,
	}
}

// Catcg calls the service to catch pokemon.
func (eu *PokemonUsecase) Catch() (model.Pokemons, error) {
	log.Println("In usecase - Catch")

	pokemons, err := eu.service.Catch()
	if err != nil {
		return nil, fmt.Errorf("catching pokemon from usecase: %v", err)
	}

	return pokemons, nil
}
