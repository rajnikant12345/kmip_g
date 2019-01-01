package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type RequestPayload struct {
	UniqueIdentifier *kmipbin.KmipTextString `kmip:"420094"`
	Attribute        []*Attribute            `kmip:"420008"`
	TemplateAttribute TemplateAttribute		`kmip:"420091"`
}
