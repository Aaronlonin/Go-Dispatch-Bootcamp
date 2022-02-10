package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Aaronlonin/capstone/model"
)

// usecase is the interface that wraps the usecase's methods
// Catch
type usecase interface {
	Catch() (model.Pokemons, error)
}

// pokemonController implements PokemonUsecase interface.
type pokemonController struct {
	usecase usecase
}

// New returns a new PokemonController instance.
func New(uc usecase) *pokemonController {
	return &pokemonController{
		usecase: uc,
	}
}

// Catch calls the usecase to return caught pokemon.
func (pc *pokemonController) Catch(w http.ResponseWriter, r *http.Request) {
	log.Println("In controller - Catch")

	// get all employees from the usecase
	pokemon, err := pc.usecase.Catch()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error catching pokemon")

		log.Fatalf("getting catching pokemon: %v", err)
	}

	// special handling if employees is empty
	if len(pokemon) == 0 {
		log.Println("no pokemon found")
		w.WriteHeader(http.StatusNotFound)

		fmt.Fprintln(w, "no pokemon found")
		return
	}

	jsonData, err := json.Marshal(pokemon)
	if err != nil {
		log.Println("error marshalling pokemon")
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintf(w, "error marshalling pokemon: %v\n", err)
	}

	// this is fine
	log.Printf("pokemon found: %+v\n", pokemon)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}
