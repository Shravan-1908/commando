package main

import (
	"github.com/Shravan-1908/commando"
)

func main() {
	registry := commando.NewCommandRegistry()
	registry.SetExecutableName("")
	registry.Parse(nil)
}
