package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type ResponseHeader struct {
	ProtocolVersion *ProtocolVersion `kmip:"420069"`
	TimeStamp       *kmipbin.KmipDate
	Nonce           *Nonce
	AttestationType *kmipbin.KmipEnum
	BatchCount      *kmipbin.KmipInt `kmip:"42000D"`
}
