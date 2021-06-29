package kmip

import "github.com/Zstro/kmip-go/kmip14"

// 4.9
//
type LocateRequestPayload struct {
	MaximumItems      int32
	OffsetItems       int32
	StorageStatusMask int32
	ObjectGroupMember kmip14.ObjectGroupMember
	Attributes        Attributes
}

type LocateResponsePayload struct {
	LocatedItems      int32
	UniqueIdentifiers []string
}
