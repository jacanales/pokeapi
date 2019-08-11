package pokeapi

import (
    "encoding/json"
    "fmt"
    pokemon "github.com/jacanales/pokeapi/internal/domain"
    "io/ioutil"
    "net/http"
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

func (p *pokemonRepository) GetPokemons(endpoint string) (pokemonList []pokemon.Url, err error) {
    response, err := http.Get(fmt.Sprintf("%v%v", EntryPoint, endpoint))
    if err != nil {
        return nil, err
    }

    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    var pokemonListResult pokemon.PokemonListResult

    err = json.Unmarshal(contents, &pokemonListResult)
    if err != nil {
        fmt.Println(err.Error())
        return nil, err
    }

    pokemonList = pokemonListResult.Results

    return
}
