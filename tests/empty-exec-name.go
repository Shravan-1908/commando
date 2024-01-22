package main

import (
	"github.com/shravanasati/commando"
)

func main() {
	registry := commando.NewCommandRegistry()
	registry.SetExecutableName("")
	registry.Parse(nil)
}
