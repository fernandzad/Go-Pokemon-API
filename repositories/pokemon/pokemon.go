package repositories

import (
	"context"
	"fmt"
	"pokeapi/database"
	"pokeapi/model"
)

type PokemonRepository struct{}

type NewMongoRepository interface {
	Create(pokemon model.PokemonMongo) *model.Error
	Read() (*model.Pokemons, *model.Error)
	Update(pokemon model.Pokemon, pokemonId int) *model.Error
	Delete(pokemonId int) *model.Error
}

var collectionName string = "pokemon"
var collection = database.GetCollection(collectionName)
var ctx = context.Background()

func New() *PokemonRepository {
	return &PokemonRepository{}
}

func (r *PokemonRepository) Create(pokemon model.PokemonMongo) *model.Error {

	_, err := collection.InsertOne(ctx, pokemon)

	if err != nil {
		return &model.Error{
			Message: err.Error(),
		}
	}

	fmt.Println("Pokemon successfully created")
	return nil
}

func (r *PokemonRepository) Read() (*model.Pokemons, *model.Error) {
	return nil, nil
}

func (r *PokemonRepository) Update(pokemon model.Pokemon, pokemonId int) *model.Error {
	return nil
}

func (r *PokemonRepository) Delete(pokemonId int) *model.Error {
	return nil
}
