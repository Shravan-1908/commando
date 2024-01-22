package main

import (
	"github.com/shravanasati/commando"
)

func main() {
	registry := commando.NewCommandRegistry()
	registry.SetExecutableName("reactor")
	registry.Register("create")
	registry.Parse(nil)
}
