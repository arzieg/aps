package azrserver

// AzrServerRepository is a interface that defines the rules around what a Azr Ressource Respository hat to be able to perform
type AzrServerRepository interface {
	GetAzrRetailPrices(serviceFamily string)
	// GetAzrRetailPrices(equipmenttype, date)
	// GetServer(uuid.UUID)
	// GetServerSpecs(uuid.UUID)
}
