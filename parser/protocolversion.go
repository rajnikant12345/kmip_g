package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type ProtocolVersion struct {
	ProtocolVersionMajor *kmipbin.KmipInt
	ProtocolVersionMinor *kmipbin.KmipInt
}
