package main

type ConfigFile struct {
	PeerPublicVMAddresses     []string `json:"peer_public_vm_addresses"`
	NeededWalletConfirmations int      `json:"needed_confirmations"`
}
