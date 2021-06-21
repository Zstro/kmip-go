package kmip

// 4.29
//
type EncryptRequestPayload struct {
	UniqueIdentifier string
	CryptoParams     CryptographicParameters
	Data             []byte
	IVCounterNonce   []byte
	CorrelationValue []byte
	InitIndicator    bool
	FinalIndicator   bool
	AdditionalData   []byte
}

type EncryptResponsePayload struct {
	UniqueIdentifier string
	Data             []byte
	IVCounterNonce   []byte
	CorrelationValue []byte
	AuthTag          []byte
}
