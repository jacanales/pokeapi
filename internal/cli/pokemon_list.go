package cli

import (
    "fmt"
    "github.com/jacanales/pokeapi/internal/domain"
    "github.com/jacanales/pokeapi/internal/storage/pokeapi"
    "github.com/jacanales/pokeapi/internal/storage/csv"
    "github.com/spf13/cobra"
    "log"
)

const (
    pokemonsEndpoint = "/pokemon"
)

func init() {
    var read domain.PokemonRepository
    var write domain.WriteRepository

    read = pokeapi.NewPokemonRepository()
    write = csv.NewWriteListRepository()

    rootCmd.AddCommand(initPokemonListCmd(read, write))
}

func initPokemonListCmd(read domain.PokemonRepository, write domain.WriteRepository) *cobra.Command {
    pokemonListCmd := &cobra.Command{
        Use: "pokemon",
        Run: getPokemonListFn(read, write),
    }

    return pokemonListCmd
}

func getPokemonListFn(read domain.PokemonRepository, write domain.WriteRepository) CobraFn {
    return func(cmd *cobra.Command, args []string) {

        fmt.Println("Pokemon list: ")
        list, _ := read.GetPokemons(pokemonsEndpoint)

        err := write.StorePokemonList(list)
        if nil != err {
            log.Fatal(err)
        }
    }
}
