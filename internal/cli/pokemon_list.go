package cli

import (
    "fmt"
    "github.com/jacanales/pokeapi/internal/domain"
    "github.com/jacanales/pokeapi/internal/storage/csv"
    gocsv "github.com/jacanales/pokeapi/internal/storage/csv"
    "github.com/jacanales/pokeapi/internal/storage/pokeapi"
    s2c "github.com/jacanales/pokeapi/internal/storage/struct2csv"
    "github.com/spf13/cobra"
    "log"
)

func init() {
    var read domain.PokemonRepository
    var write domain.WriteRepository
    var info domain.PokemonInfoRepository

    read = pokeapi.NewPokemonRepository()

    rootCmd.AddCommand(initPokemonListCmd(read, write))
    rootCmd.AddCommand(initPokemonsInfoFn(read, info))
}

func initPokemonListCmd(read domain.PokemonRepository, write domain.WriteRepository) *cobra.Command {
    pokemonListCmd := &cobra.Command{
        Use: "list",
        Short: "Get list of of pokemons",
        Run: getPokemonListFn(read, write),
    }

    pokemonListCmd.Flags().StringP("writer", "w", "csv", "Select writer")

    return pokemonListCmd
}

func initPokemonsInfoFn(read domain.PokemonRepository, write domain.PokemonInfoRepository) *cobra.Command {
    pokemonsInfoCmd := &cobra.Command{
        Use: "list-info",
        Short: "Get information of all the listed pokemons",
        Run: getPokemonsInfoFn(read, write),
    }

    return pokemonsInfoCmd
}

func getPokemonListFn(read domain.PokemonRepository, write domain.WriteRepository) CobraFn {
    return func(cmd *cobra.Command, args []string) {

        list, _ := read.GetPokemons()

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

func getPokemonsInfoFn(read domain.PokemonRepository, write domain.PokemonInfoRepository) CobraFn {
    return func(cmd *cobra.Command, args []string) {
        write = csv.NewWritePokemonRepository()

        list, _ := read.GetPokemons()

        p := make(chan domain.Pokemon)
        done := make(chan bool)

        for _, value := range list {
            go func(value domain.Url, p chan domain.Pokemon, done chan bool) {
                info, err := read.GetPokemonInfo(value)
                if nil != err {
                    fmt.Println(err.Error())

                    return
                }

                p <- info
            }(value, p, done)
        }

        i := 0
        for i <= len(list) {
            select {
            case <-p:
                fmt.Print("Write")
                _ = write.StorePokemonInfo(<-p)
                done <- true

            case <-done:
                i++
            }

        }
    }
}
