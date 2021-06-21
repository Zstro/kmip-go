package kmip

// 4.4
//
type ReKeyRequestPayload struct {
	UniqueIdentifier  string
	Offset            []byte
	TemplateAttribute TemplateAttribute
}

type ReKeyResponsePayload struct {
	UniqueIdentifier  string
	TemplateAttribute *TemplateAttribute
}
