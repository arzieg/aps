package memory

import (
	"aps/aggregate"

	"github.com/google/uuid"
)

type MemoryRepository struct {
	servers map[uuid.UUID]aggregate.AzrServer
}

// new is a factory function to generate a new repository of customers
func New() *MemoryRepository {
	return &MemoryRepository{
		servers: make(map[uuid.UUID]aggregate.AzrServer),
	}
}
