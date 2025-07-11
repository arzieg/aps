package aggregate

import (
	"aps"
	"errors"
)

var (
	ErrInvalidSKUID = errors.New("the skuid is unknown")
)

// Interface
type ServerDataProvider interface {
	GetAzrRetailPrices(serviceFamily string) ([]*aps.Server, error)
}

type AzrServerPrice struct {
	servers []*aps.Server
}

type AzrServerSpec struct {
	serverSpecs []*aps.ServerSpec
}

type AzrServer struct {
	AzrServerPrice
	AzrServerSpec
}

func NewAzrServerPrice(provider ServerDataProvider, serviceFamily string) (*AzrServerPrice, error) {
	servers, err := provider.GetAzrRetailPrices(serviceFamily)
	if err != nil {
		return nil, err
	}

	return &AzrServerPrice{
		servers: servers,
	}, nil
}

func (a *AzrServerPrice) GetServers() []*aps.Server {
	return a.servers
}

/*
func (s *AzrServerPrice) GetID() (skuid string) {
	return s.server.SkuId
}

func (s *AzrServerPrice) SetID(skuid string) {
	if s.server.SkuId == "" {
		s.server = &aps.Server{}
	}
	s.server.SkuId = skuid
}

func (s *AzrServerPrice) GetAll() *AzrServerPrice {
	server := AzrServerPrice
	return server
}
*/
