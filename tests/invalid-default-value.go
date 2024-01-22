package main

import (
	"github.com/shravanasati/commando"
)

func main() {
	registry := commando.NewCommandRegistry()
	registry.SetExecutableName("reactor")
	registry.Register("create").AddFlag("dir,d", "directory", commando.String, 21)
	registry.Parse(nil)
}
