package csv

import (
    "encoding/csv"
    "fmt"
    pokemon "github.com/jacanales/pokeapi/internal/domain"
    "os"
)

const (
    CsvFile = "/tmp/pokemon_list.csv"
)


type writeListRepository struct {

}

func NewWriteListRepository () pokemon.WriteRepository {
    return &writeListRepository{}
}

func NewWritePokemonRepository () pokemon.PokemonInfoRepository {
    return &writeListRepository{}
}

func (w writeListRepository) StorePokemonList(l []pokemon.Url) error {
    file, err := os.Create(CsvFile)
    if nil != err {
        return err
    }

    defer func() {
        e := file.Close()

        if nil != e {
            err = e
        }
    }()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    _ = writer.Write([]string{"Name", "URL"})

    for _, value := range l {
        err := writer.Write(value.ToArray())
        if nil != err {
            return err
        }
    }

    return nil
}


func (w writeListRepository) StorePokemonInfo(p pokemon.Pokemon) error {
    var file *os.File
    var err error

    if _, err := os.Stat(CsvFile); err == nil {
        file, err = os.OpenFile(CsvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if nil != err {
            fmt.Println(err.Error())
            return err
        }
    } else if os.IsNotExist(err) {
        file, err = os.Create(CsvFile)
        if nil != err {
            fmt.Println(err.Error())
            return err
        }
    }

    defer func() {
        e := file.Close()

        if nil != e {
            err = e
        }
    }()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    err = writer.Write(p.ToArray())
    if nil != err {
        return err
    }

    return nil
}
