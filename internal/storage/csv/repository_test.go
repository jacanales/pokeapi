package csv

import (
    "fmt"
    pokemon "github.com/jacanales/pokeapi/internal/domain"
    "strconv"
    "testing"
)

func BenchmarkWriteListRepository_StorePokemonList(b *testing.B) {
    repo := NewWriteListRepository()
    b.ResetTimer()
    for n:= 0; n < b.N; n++ {
        _ = repo.StorePokemonList([]pokemon.Url{
            { Name: fmt.Sprintf("Test %s", strconv.Itoa(n)), Url: fmt.Sprintf("http://localhost/test/%s", strconv.Itoa(n)) },
        })
    }
}
