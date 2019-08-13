package cli

import (
    "github.com/jacanales/pokeapi/internal/domain"
    "github.com/jacanales/pokeapi/internal/storage/csv"
    gocsv "github.com/jacanales/pokeapi/internal/storage/csv"
    "github.com/jacanales/pokeapi/internal/storage/pokeapi"
    s2c "github.com/jacanales/pokeapi/internal/storage/struct2csv"
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

    rootCmd.AddCommand(initPokemonListCmd(read, write))
}

func initPokemonListCmd(read domain.PokemonRepository, write domain.WriteRepository) *cobra.Command {
    pokemonListCmd := &cobra.Command{
        Use: "pokemon",
        Run: getPokemonListFn(read, write),
    }

    pokemonListCmd.Flags().StringP("writer", "w", "csv", "Select writer")

    return pokemonListCmd
}

func getPokemonListFn(read domain.PokemonRepository, write domain.WriteRepository) CobraFn {
    return func(cmd *cobra.Command, args []string) {

        list, _ := read.GetPokemons(pokemonsEndpoint)

        opt, _ := cmd.Flags().GetString("writer")

        switch opt {
        default:
        case "csv":
            write = csv.NewWriteListRepository()
        case "s2c":
            write = s2c.NewWriteListRepository()
        case "gocsv":
            write = gocsv.NewWriteListRepository()
        }

        err := write.StorePokemonList(list)
        if nil != err {
            log.Fatal(err)
        }
    }
}
