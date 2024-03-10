package main

import (
	"medclin/setup"

	_ "github.com/lib/pq"
)

func main() {
	setup.StartWS()
	setup.StartDB()
}
