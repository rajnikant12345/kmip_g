package parser

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

type RequestHeader struct {
	ProtocolVersion              *ProtocolVersion
	MaximumResponseSize          *kmipbin.KmipInt
	AsynchronousIndicator        *kmipbin.KmipBoolean
	AttestationCapableIndicator  *kmipbin.KmipBoolean
	AttestationType              *kmipbin.KmipEnum
	Authentication               *Authentication
	BatchErrorContinuationOption *kmipbin.KmipEnum
	BatchOrderOption             *kmipbin.KmipBoolean
	TimeStamp                    *kmipbin.KmipDate
	BatchCount                   *kmipbin.KmipInt
}
