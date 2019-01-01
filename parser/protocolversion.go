package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type ProtocolVersion struct {
	ProtocolVersionMajor *kmipbin.KmipInt `kmip:"42006A"`
	ProtocolVersionMinor *kmipbin.KmipInt `kmip:"42006B"`
}
