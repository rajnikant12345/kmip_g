package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type ResponseHeader struct {
	ProtocolVersion *ProtocolVersion
	TimeStamp       *kmipbin.KmipDate
	Nonce           *Nonce
	AttestationType *kmipbin.KmipEnum
	BatchCount      *kmipbin.KmipInt
}
