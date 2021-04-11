package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"pokeapi/controller"
	"pokeapi/model"
	repositories "pokeapi/repositories/pokemon"
	"pokeapi/router"
	csvservice "pokeapi/service/csv"
	httpservice "pokeapi/service/http"
	"pokeapi/usecase"
)

func insertPokemon() {
	pokemonRepository := repositories.New()

	pokemon := model.PokemonMongo{
		Name:      "Charmander",
		URL:       "https://pokeapi.co/api/v2/pokemon/4/",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := pokemonRepository.Create(pokemon)

	if err != nil {
		fmt.Print(err.Message)
	}
}

func readPokemons() {
	pokemonRepository := repositories.New()

	pokemons, err := pokemonRepository.Read()

	if err != nil {
		fmt.Print(err.Message)
	}

	fmt.Println(*pokemons)
}

func updatePokemon() {
	pokemonRepository := repositories.New()
	id := "60725516a48879e027760072"

	pokemon := model.PokemonMongo{
		Name: "Charizard",
		URL:  "https://pokeapi.co/api/v2/pokemon/6/",
	}

	err := pokemonRepository.Update(pokemon, id)

	if err != nil {
		fmt.Print(err.Message)
	}
}

func deletePokemon() {
	pokemonRepository := repositories.New()
	id := "6072578fca89d45ecac4fad8"

	err := pokemonRepository.Delete(id)

	if err != nil {
		fmt.Print(err.Message)
	}
}

func main() {

	// insertPokemon()
	// updatePokemon()
	// readPokemons()
	// deletePokemon()
	// readPokemons()

	csvService := csvservice.New()
	httpService := httpservice.New()
	usecase := usecase.New(csvService, httpService)
	controller := controller.New(usecase)

	router := router.New(controller)
	r := router.InitRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
