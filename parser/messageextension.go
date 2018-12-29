package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type MessageExtension struct {
	VendorIdentification *kmipbin.KmipTextString
	CriticalityIndicator *kmipbin.KmipBoolean
	VendorExtension      *VendorExtension
}
