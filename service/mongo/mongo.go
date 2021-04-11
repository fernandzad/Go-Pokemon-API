package mongo

import (
	"pokeapi/model"
	repositories "pokeapi/repositories/pokemon"
)

type MongoService struct {
	mongRepository repositories.PokemonRepository
}

type NewMongoService interface {
	Create(pokemon model.PokemonMongo) *model.Error
	Read() (*model.Pokemons, *model.Error)
	Update(pokemon model.PokemonMongo, pokemonId string) *model.Error
	Delete(pokemonId string) *model.Error
}

func (ms *MongoService) Create(pokemon model.PokemonMongo) *model.Error {

	err := ms.Create(pokemon)

	if err != nil {
		return &model.Error{
			Message: err.Message,
		}
	}

	return nil
}

func (ms *MongoService) Read() (*model.Pokemons, *model.Error) {

	pokemons, err := ms.Read()

	if err != nil {
		return nil, &model.Error{
			Message: err.Message,
		}
	}

	return pokemons, nil
}

func (ms *MongoService) Update(pokemon model.PokemonMongo, pokemonId string) *model.Error {
	err := ms.Update(pokemon, pokemonId)

	if err != nil {
		return &model.Error{
			Message: err.Message,
		}
	}

	return nil
}

func (ms *MongoService) Delete(pokemonId string) *model.Error {

	err := ms.Delete(pokemonId)

	if err != nil {
		return &model.Error{
			Message: err.Message,
		}
	}

	return nil
}
