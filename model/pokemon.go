package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Pokemon Type
type Pokemon struct {
	ID        int       `json:ID,omitempty`
	Name      string    `json:Name,omitempty`
	URL       string    `json:Url,omitempty`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:updated_at,omitempty`
}

type PokemonMongo struct {
	ID        primitive.ObjectID `bson:"_id" json:ID,omitempty`
	Name      string             `json:Name,omitempty`
	URL       string             `json:Url,omitempty`
	CreatedAt time.Time          `bson:"created_at" json:created_at`
	UpdatedAt time.Time          `bson:"updated_at" json:updated_at,omitempty`
}

// Pokemons pokemons list
type Pokemons []PokemonMongo

type SinglePokeExternal struct {
	Name string `json:name,omitempty`
	URL  string `json:url,omitempty`
}

type PokemonExternal struct {
	Count    int                  `json:count`
	Next     string               `json:next`
	Previous string               `json:previous`
	Results  []SinglePokeExternal `json:results`
}

type TypeNumberFilter struct {
	Even string
	Odd  string
}
