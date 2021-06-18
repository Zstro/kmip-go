package kmip

// 4.30
//
type DecryptRequestPayload struct {
	UniqueIdentifier string
	CryptoParams     CryptographicParameters
	Data             []byte
	IVCounterNonce   []byte
	CorrelationValue []byte
	InitIndicator    bool
	FinalIndicator   bool
	AdditionalData   []byte
	AuthTag          []byte
}

type DecryptResponsePayload struct {
	UniqueIdentifier string
	Data             []byte
	CorrelationValue []byte
}
