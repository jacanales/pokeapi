package pokeapi

import (
    "fmt"
    pokemon "github.com/jacanales/pokeapi/internal/domain"
    "github.com/json-iterator/go"
    "io/ioutil"
    "net/http"
)

const (
    EntryPoint = "http://pokeapi.co/api/v2"
    PokemonsEndpoint = "/pokemon"
)

type pokemonRepository struct {
    url string
}

func NewPokemonRepository() pokemon.PokemonRepository {
    return &pokemonRepository{EntryPoint}
}

func (p *pokemonRepository) GetPokemons() (pokemonList []pokemon.Url, err error) {
    var url string
    var pokemonListResult pokemon.PokemonListResult
    url = fmt.Sprintf("%v%v?offset=0&limit=2", EntryPoint, PokemonsEndpoint)

    err = p.parseJsonResponse(url, &pokemonListResult)

    pokemonList = pokemonListResult.Results

    return
}

func (p *pokemonRepository) GetPokemonInfo(pokemonUrl pokemon.Url) (pokemonInfo pokemon.Pokemon, err error) {
    var url string
    url = fmt.Sprintf("%s", pokemonUrl.Url)

    err = p.parseJsonResponse(url, &pokemonInfo)

    return
}

func (p *pokemonRepository) parseJsonResponse(uri string, t interface{}) (err error) {
    response, err := http.Get(uri)
    defer closeResponseBody(response)
    if err != nil {
        return err
    }

    fmt.Println(response.StatusCode)

    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return err
    }

    var json = jsoniter.ConfigCompatibleWithStandardLibrary

    err = json.Unmarshal(contents, &t)
    if err != nil {
        return err
    }

    return
}

func closeResponseBody(r *http.Response) {
    err := r.Body.Close()
    if nil != err {
        fmt.Print(err.Error())
    }
}
