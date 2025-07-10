package memory

type ServerItem struct {
	CurrencyCode     string  `json:"currencyCode"`
	TierMinimumUnits int     `json:"tierMinimumUnits"`
	RetailPrice      float64 `json:"retailPrice"`
	//UnitPrice            float64
	ArmRegionName string `json:"armRegionName"`
	//Location             string
	EffectiveStartDate string `json:"effectiveStartDate"`
	//MeterId              string
	//MeterName            string
	//ProductId            string
	//SkuId                string
	ProductName string `json:"productName"`
	SkuName     string `json:"skuName"`
	//ServiceName          string
	//ServiceId            string
	//ServiceFamily        string
	UnitOfMeasure string `json:"unitOfMeasure"`
	Type          string `json:"type"`
	//IsPrimaryMeterRegion bool
	//ArmSkuName           string
}
