package main

import (
	"aps/aggregate"
	"aps/storage/azrserver/memory"
	"fmt"
	"log"
	"os"
)

func main() {

	/*
		myServer := memory.New()
		serviceFamily := "Compute"
		myAggServer, err := myServer.GetAzrRetailPrices(serviceFamily)
		if err != nil {
			log.Fatalf("Got error %v", err)
		}
	*/

	// Initialize the repository
	repo := &memory.MemoryRepository{}

	// Service families you want to fetch prices for
	serviceFamilies := []string{"Compute", "Storage", "Networking"}

	for _, sf := range serviceFamilies {
		azrPrice, err := aggregate.NewAzrServerPrice(repo, sf)
		if err != nil {
			log.Printf("Error loading prices for %s: %v", sf, err)
			continue
		}

		fmt.Printf("\n--- Prices for service family: %s ---\n", sf)
		for _, srv := range azrPrice.GetServers() { // assumes a helper method to expose servers
			fmt.Printf("SKU: %s | Price: %.2f %s | Region: %s  | SKUID: %s\n",
				srv.SkuName, srv.RetailPrice, srv.CurrencyCode, srv.Location, srv.SkuId)
		}
	}

	os.Exit(0)
}
