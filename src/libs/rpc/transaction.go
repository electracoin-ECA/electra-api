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

func GetTransaction(txnId string) (*GetTransactionResponse, error) {
	var txn *GetTransactionResponse
	err := query("gettransaction", []string{
		txnId,
	}, &txn)

	if err != nil {
		return nil, err
	}

	return txn, nil

}

func SendTransaction(rawTxn string) string {
	txnHash, err := queryBytes("sendrawtransaction", []string{
		rawTxn,
	})

	if err != nil {
		return ""
	}

	return string(txnHash)
}
