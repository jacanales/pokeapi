package csv

import (
    pokemon "github.com/jacanales/pokeapi/internal/domain"
    "github.com/mohae/struct2csv"
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

    writer := struct2csv.NewWriter(file)
    defer writer.Flush()

    _ = writer.WriteStructs(l)

    return nil
}
