package memory

type ServerItem struct {
	CurrencyCode       string  `json:"currencyCode"`
	TierMinimumUnits   float64 `json:"tierMinimumUnits"`
	RetailPrice        float64 `json:"retailPrice"`
	UnitPrice          float64 `json:"unitPrice"`
	ArmRegionName      string  `json:"armRegionName"`
	Location           string  `json:"location"`
	EffectiveStartDate string  `json:"effectiveStartDate"`
	MeterId            string  `json:"meterId"`
	MeterName          string  `json:"meterName"`
	ProductId          string  `json:"productId"`
	SkuId              string  `json:"skuId"`
	ProductName        string  `json:"productName"`
	SkuName            string  `json:"skuName"`
	ServiceName        string  `json:"serviceName"`
	ServiceId          string  `json:"serviceId"`
	ServiceFamily      string  `json:"serviceFamily"`
	UnitOfMeasure      string  `json:"unitOfMeasure"`
	Type               string  `json:"type"`
	ArmSkuName         string  `json:"armSkuName"`
}

type ServerItemResponse struct {
	Items []ServerItem `json:"Items"`
	Next  string       `json:"NextPageLink"`
}
