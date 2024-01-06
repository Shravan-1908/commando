package main

import (
	"fmt"

	"github.com/Shravan-1908/commando"
)

func main() {

	// set CLI executable, version and description
	commando.
		SetExecutableName("reactor").
		SetVersion("v1.0.0").
		SetDescription("Reactor is a command-line tool to generate React projects.\nIt helps you create components, write test cases, start a development server and much more.").
		SetEventListener(func(eventName string) {
			//fmt.Println("event-name: ", eventName)
		})

	// configure the root-command
	// $ reactor <category>  --verbose|-V  --version|-v  --help|-h
	commando.
		Register(nil).
		AddArgument("category", "category of the information to look for", ""). // required
		AddFlag("verbose,V", "display log information ", commando.Bool, nil).   // optional
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// register `create` sub-command
	// $ reactor create <name> [version] [files] --dir|-d <dir>  --type|-t [type]  --timeout [timeout]  --verbose|-v  --help|-h  --no-clean
	commando.
		Register("create").
		SetDescription("This command creates a component of a given type and outputs component files in the project directory.").
		SetShortDescription("creates a component").
		AddArgument("name", "name of the component to create", "").                                  // required
		AddArgument("version", "version of the component", "1.0.0").                                 // optional
		AddArgument("files...", "files to remove once component is created", "").                    // variadic, optional
		AddFlag("dir, d", "output directory for the component files", commando.String, nil).         // required
		AddFlag("type, t", "type of the component to create", commando.String, "simple_type").       // optional
		AddFlag("timeout", "operation timeout in seconds", commando.Int, 60).                        // optional
		AddFlag("verbose,v", "display logs while creating the component files", commando.Bool, nil). // optional
		AddFlag("no-clean", "avoid cleanup of the component directory", commando.Bool, nil).         // optional, inverted flag
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// register `serve` sub-command
	// $ reactor serve --verbose|-v  --help|-h
	commando.
		Register("serve").
		SetDescription("This command starts the Webpack dev-server on an available port.").
		SetShortDescription("starts a development server").
		AddFlag("verbose,v", "display logs while serving the project", commando.Bool, nil). // optional
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// register `build` sub-command
	// $ reactor build  --dir|-d <dir>  --verbose|-v  --help|-h
	commando.
		Register("build").
		SetDescription("This command builds the project with Webpack and outputs the build files in the given directory.").
		SetShortDescription("creates build artifacts").
		AddFlag("dir,d", "output directory of the build files", commando.String, nil).      // required
		AddFlag("verbose,v", "display logs while serving the project", commando.Bool, nil). // optional
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// parse command-line arguments from the STDIN
	commando.Parse(nil)
}
