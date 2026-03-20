module github.com/RiMaBo/go_pokedex

go 1.26.1

replace internal/pokeapi => ./internal/pokeapi
require internal/pokeapi v0.0.0

replace internal/pokecache => ./internal/pokecache
require internal/pokecache v0.0.0
