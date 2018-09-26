package rpc

type GetTransactionResponse struct {
	Amount        float64
	Fee           float64
	Confirmations int
	Blockhash     string
	Blockindex    string
	Blocktime     int
	Txid          string
	Time          int
	TimeReceived  int
	Details       []Detail
	Hex           string
}

type Detail struct {
	Account  string
	Address  string
	Category string
	Amount   float64
	Vout     float64
	Fee      float64 `json:"fee,omitempty"`
}

func SendTransaction(rawTxn string) string {
	txnHash, err := QueryBytes("sendrawtransaction", []string{
		rawTxn,
	})

	if err != nil {
		return ""
	}

	return string(txnHash)
}
