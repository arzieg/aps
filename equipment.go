package aps

import (
	"github.com/google/uuid"
)

// https://learn.microsoft.com/de-de/rest/api/cost-management/retail-prices/azure-retail-prices
// https://prices.azure.com/api/retail/prices?currencyCode='EUR'&$filter=serviceFamily eq 'Compute'

type Server struct {
	ID                   uuid.UUID
	CurrencyCode         [3]string
	TierMinimumUnits     float64
	RetailPrice          float64
	UnitPrice            float64
	ArmRegionName        string
	Location             string
	EffectiveStartDate   string
	MeterId              string
	MeterName            string
	ProductId            string
	SkuId                string
	ProductName          string
	SkuName              string
	ServiceName          string
	ServiceId            string
	ServiceFamily        string
	UnitOfMeasure        string
	Type                 string
	IsPrimaryMeterRegion bool
	ArmSkuName           string
}

type ServerSpec struct {
	ID    uuid.UUID
	specs []string
}
