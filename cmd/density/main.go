package main

import "github.com/junderhill/density/internal/cmd"

var (
	commit  string
	builtAt string
	version string
)

func main() {
	cmd.Execute()
}
