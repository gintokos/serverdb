package main

import (
	db "serverdb/internal"
)

func main() {
	db.MustStartDBserver("config.json")
}
