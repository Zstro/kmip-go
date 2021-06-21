package kmip

import (
	"github.com/Zstro/kmip-go/kmip14"
)

// 4.20
//
type RevokeRequestPayload struct {
	UniqueIdentifier string
	RevocationReason RevocationReason
}

type RevocationReason struct {
	RevocationReasonCode kmip14.RevocationReasonCode
	RevocationMessage    string
}

type RevokeResponsePayload struct {
	UniqueIdentifier string
}
