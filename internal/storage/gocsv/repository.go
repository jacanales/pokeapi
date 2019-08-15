package gocsv

import (
    "encoding/csv"
    "github.com/gocarina/gocsv"
    pokemon "github.com/jacanales/pokeapi/internal/domain"
    "os"
)

const (
    File = "/tmp/pokemon_list.csv"
)

type PokemonRow struct {
    Name    string `csv:"name"`
    Url     string `csv:"url"`
}

type writeListRepository struct {

}

func NewWriteListRepository () pokemon.WriteRepository {
    return &writeListRepository{}
}

func (w writeListRepository) StorePokemonList(l []pokemon.Url) error {
    file, err := os.Create(File)
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

    rows := []*PokemonRow{}

    for _, value := range l {
        rows = append(rows, &PokemonRow{value.Name, value.Url})
    }

    err = gocsv.MarshalFile(&rows, file)
    if nil != err {
        return err
    }

    return nil
}
