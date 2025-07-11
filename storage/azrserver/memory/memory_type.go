package memory

type ServerItem struct {
	CurrencyCode     string  `json:"currencyCode"`
	TierMinimumUnits float64 `json:"tierMinimumUnits"`
	RetailPrice      float64 `json:"retailPrice"`
	//UnitPrice            float64
	ArmRegionName      string `json:"armRegionName"`
	Location           string `json:"location"`
	EffectiveStartDate string `json:"effectiveStartDate"`
	//MeterId              string
	//MeterName            string
	ProductId   string `json:"productId"`
	SkuId       string `json:"skuId"`
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

type ServerItemResponse struct {
	Items []ServerItem `json:"Items"`
	Next  string       `json:"NextPageLink"`
}
