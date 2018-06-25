package rpc

// VerifyMessageResponse model.
type VerifyMessageResponse struct {
	Result bool `json:"result"`
	Error  struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	ID string `json:"id"`
}

// VerifyMessage gets the daemon node info.
func VerifyMessage(
	addressHash string,
	signature string,
	message string,
) (*VerifyMessageResponse, error) {
	p := []string{addressHash, signature, message}
	var r *VerifyMessageResponse
	err := query("verifymessage", p, &r)

	return r, err
}
