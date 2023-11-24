package main

import (
	"os"

	"git.sr.ht/~jamesponddotco/wpmod/internal/app"
)

func main() {
	os.Exit(app.Run(os.Args))
}
