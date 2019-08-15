package domain

import "strconv"

type Url struct {
    Name string `json:"name"`
    Url  string `json:"url"`
}

type Ability struct {
    Ability Url  `json:"ability"`
    Hidden  bool `json:"is_hidden"`
    Slot    int  `json:"slot"`
}

type Forms struct {
    Name string `json:"name"`
    Url  string `json:"url"`
}

type GameIndices struct {
    GameIndex int `json:"game_index"`
    Version   Url `json:"version"`
}

type Item struct {
    Item           Url `json:"item"`
    VersionDetails []struct {
        Rarity  int `json:"rarity"`
        Version Url `json:"version"`
    } `json:"version_details"`
}

type Move struct {
    Move                Url `json:"move"`
    VersionGroupDetails []struct {
        LevelLearnedAt  int `json:"level_learned_at"`
        MoveLearnMethod Url `json:"move_learn_method"`
        VersionGroup    Url `json:"version_group"`
    } `json:"version_group_details"`
    Name  string `json:"name"`
    Order int    `json:"order"`
}

type Sprites struct {
    BackDefault      string `json:"back_default"`
    BackFemale       string `json:"back_female"`
    BackShiny        string `json:"back_shiny"`
    BackShinyFemale  string `json:"back_shiny_female"`
    FrontDefault     string `json:"front_default"`
    FrontFemale      string `json:"front_female"`
    FrontShiny       string `json:"front_shiny"`
    FrontShinyFemale string `json:"front_shiny_female"`
}

type Stat struct {
    BaseStat int `json:"base_stat"`
    Effort   int `json:"effort"`
    Stat     Url `json:"stat"`
}

type PokemonType struct {
    Slot int `json:"slot"`
    Type Url `json:"type"`
}

type Pokemon struct {
    Abilities              []Ability     `json:"abilities"`
    BaseExperience         int           `json:"base_experience"`
    Forms                  []Forms       `json:"forms"`
    GameIndices            []GameIndices `json:"game_indices"`
    Height                 int           `json:"height"`
    HeldItems              []Item        `json:"held_items"`
    Id                     int           `json:"id"`
    IsDefault              bool          `json:"is_default"`
    LocationAreaEncounters string        `json:"location_area_encounters"`
    Moves                  []Move        `json:"moves"`
    Name                   string        `json:"name"`
    Order                  int           `json:"order"`
    Species                Url           `json:"species"`
    Sprites                Sprites       `json:"sprites"`
    Stats                  []Stat        `json:"stats"`
    Types                  []PokemonType `json:"types"`
    Weight                 int           `json:"weight"`
}

type PokemonListResult struct {
    Count    int    `json:"count"`
    Previous string `json:"previous"`
    Next     string `json:"next"`
    Results  []Url  `json:"results"`
}

// PokemonRepository definition of methods to access pokemon's data
type PokemonRepository interface {
    GetPokemons() ([]Url, error)
    GetPokemonInfo(pokemonUrl Url) (pokemonInfo Pokemon, err error)
}

type WriteRepository interface {
    StorePokemonList(l []Url) error
}

type PokemonInfoRepository interface {
    StorePokemonInfo(p Pokemon) error
}

func (u Url) ToArray() (list []string) {
    list = append(list, u.Name)
    list = append(list, u.Url)

    return
}

func (p Pokemon) ToArray() (list []string) {
    list = append(list, strconv.Itoa(p.Id))
    list = append(list, p.Name)
    list = append(list, strconv.Itoa(p.BaseExperience))
    list = append(list, strconv.Itoa(p.Height))

    return
}
