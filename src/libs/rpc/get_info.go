package rpc

// GetInfoResponse model.
type GetInfoResponse struct {
	Result struct {
		Version         string  `json:"version"`
		ProtocolVersion int64   `json:"protocolversion"`
		WalletVersion   float64 `json:"walletversion"`
		Balance         float64 `json:"balance"`
		NewMint         float64 `json:"newmint"`
		Stake           float64 `json:"stake"`
		Blocks          int64   `json:"blocks"`
		TimeOffset      int64   `json:"timeoffset"`
		MoneySupply     float64 `json:"moneysupply"`
		Connections     int64   `json:"connections"`
		Proxy           string  `json:"proxy"`
		IP              string  `json:"ip"`
		Difficulty      struct {
			ProofOfWork  float64 `json:"proof-of-work"`
			ProofOfStake float64 `json:"proof-of-stake"`
		} `json:"difficulty"`
		Testnet       bool    `json:"testnet"`
		KeyPoolOldest int64   `json:"keypoololdest"`
		KeyPoolSize   int64   `json:"keypoolsize"`
		PayTxFee      float64 `json:"paytxfee"`
		MinInput      float64 `json:"mininput"`
		UnlockedUntil int64   `json:"unlocked_until"`
		Errors        string  `json:"errors"`
	} `json:"result,omitempty"`
	Error struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
	ID string `json:"id,omitempty"`
}

// GetInfo gets the daemon node info.
func GetInfo() (*GetInfoResponse, error) {
	var info *GetInfoResponse
	err := query("getinfo", nil, &info)

	return info, err
}
