package repositories

import "pokeapi/model"

type PokemonRepository struct{}

type NewMongoRepository interface {
	Create(pokemon model.Pokemon) *model.Error
	Read() (*model.Pokemons, *model.Error)
	Update(pokemon model.Pokemon, pokemonId int) *model.Error
	Delete(pokemonId int) *model.Error
}

func New() *PokemonRepository {
	return &PokemonRepository{}
}

func (r *PokemonRepository) Create(pokemon model.Pokemon) *model.Error {

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
