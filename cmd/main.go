package main

import (
	"aps/storage/azrserver/memory"
	"os"
)

func main() {

	myServer := memory.New()

	serviceFamily := "Compute"
	myServer.GetAzrRetailPrices(serviceFamily)

	os.Exit(0)
}
