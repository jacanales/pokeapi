package pokeapi

import (
    "fmt"
    pokemon "github.com/jacanales/pokeapi/internal/domain"
)

const (
    EntryPoint = "http://pokeapi.co/api/v2"
)

type pokemonRepository struct {
    url string
}

func NewPokemonRepository() pokemon.PokemonRepository {
    return &pokemonRepository{EntryPoint}
}

func (p *pokemonRepository) GetPokemons() ([]pokemon.Pokemon, error) {
    fmt.Println("Test")

    return nil, nil
}
