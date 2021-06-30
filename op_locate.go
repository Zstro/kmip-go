package kmip

import "github.com/Zstro/kmip-go/kmip14"

// 4.9
//
type LocateRequestPayload struct {
	MaximumItem       int32
	OffsetItem        int32
	StorageStatusMask int32
	ObjectGroupMember kmip14.ObjectGroupMember
	Attribute         []Attribute
}

type LocateResponsePayload struct {
	LocatedItem      int32
	UniqueIdentifier []string
}

// for swxa
type LocateKeyRequestPayload struct {
	Attribute         []Attribute
}
