package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type ResponsePayload struct {
	UniqueIdentifier *kmipbin.KmipTextString `kmip:"420094"`
	Attribute        []*Attribute            `kmip:"420008"`
}
