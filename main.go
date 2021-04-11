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

func main() {

	// insertPokemon()
	readPokemons()

	csvService := csvservice.New()
	httpService := httpservice.New()
	usecase := usecase.New(csvService, httpService)
	controller := controller.New(usecase)

	router := router.New(controller)
	r := router.InitRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
