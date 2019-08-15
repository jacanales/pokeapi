package pokeapi

import "testing"

func BenchmarkPokemonRepository_GetPokemons(b *testing.B) {
    repo := NewPokemonRepository()
    b.ResetTimer()
    for n := 0; n < b.N; n++ {
        repo.GetPokemons()
    }
}
