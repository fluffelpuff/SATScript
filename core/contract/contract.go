package contract

type Contract struct {
}

// Gibt an wieivle Peer Verbindungen der Contract hat
func (obj *Contract) GetTotalPeerCount() (uint, error) {
	return 0, nil
}

// Gibt alle Contract Informationen aus
func (obj *Contract) GetContractInformation() (*ContractInformations, error) {
	conti := ContractInformations{ContractID: "CID", State: "active", Amount: 100000, TotalPeers: 0}
	return &conti, nil
}

type ContractInformations struct {
	ContractID string
	State      string
	Amount     uint64
	TotalPeers uint64
}
