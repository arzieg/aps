package aggregate

import (
	"aps"
	"errors"
)

var (
	ErrInvalidSKUID = errors.New("the skuid is unknown")
)

type AzrServerPrice struct {
	server *aps.Server
}

type AzrServerSpec struct {
	serverSpec *aps.ServerSpec
}

type AzrServer struct {
	AzrServerPrice
	AzrServerSpec
}

func NewAzrServerPrice(skuid string) (AzrServerPrice, error) {
	if skuid == "" {
		return AzrServerPrice{}, ErrInvalidSKUID
	}

	server := &aps.Server{
		SkuId: skuid,
	}

	return AzrServerPrice{
		server: server,
	}, nil

}

func (s *AzrServerPrice) GetID() (skuid string) {
	return s.server.SkuId
}

func (s *AzrServerPrice) SetID(skuid string) {
	if s.server.SkuId == "" {
		s.server = &aps.Server{}
	}
	s.server.SkuId = skuid
}
