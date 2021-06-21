package kmip

import "github.com/Zstro/kmip-go/kmip14"

// 4.11
//
type GetRequestPayload struct {
	UniqueIdentifier string
}

type GetResponsePayload struct {
	ObjectType       kmip14.ObjectType
	UniqueIdentifier string
	SymmetricKey     SymmetricKey
	PrivateKey       PrivateKey
	PublicKey        PublicKey
}
