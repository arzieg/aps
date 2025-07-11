package memory

import (
	"aps"
	"aps/aggregate"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const azureRetailPricesAPI = "https://prices.azure.com/api/retail/prices"

type MemoryRepository struct {
	servers map[string]aggregate.AzrServerPrice
}

// new is a factory function to generate a new repository of customers
func New() *MemoryRepository {
	return &MemoryRepository{
		servers: make(map[string]aggregate.AzrServerPrice),
	}
}

func (mr *MemoryRepository) GetAzrRetailPrices(serviceFamily string) (servers []*aps.Server, err error) {

	filter := fmt.Sprintf("serviceFamily eq '%s' and armRegionName eq 'westeurope' and serviceName eq 'Virtual Machines'", serviceFamily)
	params := url.Values{}
	params.Add("$filter", filter)

	url := fmt.Sprintf("%s?%s", azureRetailPricesAPI, params.Encode())
	fmt.Printf("\nURL=%s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	var data ServerItemResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Failed to parse response: %v", err)
	}

	for _, item := range data.Items {
		fmt.Printf("SKU: %s | Price: %.2f %s | Region: %s  | SKUID: %s\n",
			item.SkuName, item.RetailPrice, item.CurrencyCode, item.Location, item.SkuId)
	}

	for _, item := range data.Items {
		srv := &aps.Server{
			SkuName:      item.SkuName,
			RetailPrice:  item.RetailPrice,
			CurrencyCode: item.CurrencyCode,
			Location:     item.Location,
			SkuId:        item.SkuId,
		}
		servers = append(servers, srv)
	}

	return servers, nil
}
