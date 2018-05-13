package bitgo

// TODO

type Webhook struct {
	Hash     string `json:"hash"`
	Transfer string `json:"transfer"`
	Coin     string `json:"coin"`
	Type     string `json:"type"`
	Wallet   string `json:"wallet"`
}
