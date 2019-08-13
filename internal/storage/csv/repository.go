package csv

import (
    "encoding/csv"
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
        err := writer.Write(value.ConvertToArray())
        if nil != err {
            return err
        }
    }

    return nil
}
