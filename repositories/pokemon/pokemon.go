package repositories

import (
	"context"
	"fmt"
	"pokeapi/database"
	"pokeapi/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PokemonRepository struct{}

type NewMongoRepository interface {
	Create(pokemon model.PokemonMongo) *model.Error
	Read() (*model.Pokemons, *model.Error)
	Update(pokemon model.Pokemon, pokemonId int) *model.Error
	Delete(pokemonId int) *model.Error
}

func New() *PokemonRepository {
	return &PokemonRepository{}
}

func (r *PokemonRepository) Create(pokemon model.PokemonMongo) *model.Error {

	var collectionName string = "pokemon"
	var collection = database.GetCollection(collectionName)
	var ctx = context.Background()

	id := primitive.NewObjectID()
	pokemon.ID = id

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
	var pokemons model.Pokemons
	var collectionName string = "pokemon"
	var collection = database.GetCollection(collectionName)
	var ctx = context.Background()

	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, &model.Error{
			Message: err.Error(),
		}
	}

	for cur.Next(ctx) {
		var pokemon model.PokemonMongo
		err := cur.Decode(&pokemon)

		if err != nil {
			return nil, &model.Error{
				Message: err.Error(),
			}
		}

		pokemons = append(pokemons, pokemon)
	}

	return &pokemons, nil
}

func (r *PokemonRepository) Update(pokemon model.Pokemon, pokemonId int) *model.Error {
	return nil
}

func (r *PokemonRepository) Delete(pokemonId int) *model.Error {
	return nil
}
