package cli

import (
    "fmt"
    "github.com/jacanales/pokeapi/internal/domain"
    "github.com/jacanales/pokeapi/internal/storage/pokeapi"
    "github.com/spf13/cobra"
)

const (
    pokemonsEndpoint = "/pokemon"
)

func init() {
    var repo domain.PokemonRepository
    repo = pokeapi.NewPokemonRepository()

    rootCmd.AddCommand(initPokemonListCmd(repo))
}

func initPokemonListCmd(repository domain.PokemonRepository) *cobra.Command {
    pokemonListCmd := &cobra.Command{
        Use: "pokemon",
        Run: getPokemonListFn(repository),
    }

    return pokemonListCmd
}

func getPokemonListFn(repository domain.PokemonRepository) CobraFn {
    return func(cmd *cobra.Command, args []string) {

        fmt.Println("Pokemon list: ")
        list, _ := repository.GetPokemons(pokemonsEndpoint)

        fmt.Println(list)
    }
}
