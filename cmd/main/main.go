package main

import (
	db "github.com/gintokos/serverdb/internal"
)

func main() {
	db.MustStartDBserver("config.json")
}
