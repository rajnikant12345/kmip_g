package kmiperror

import "github.com/rajnikant12345/kmip_g/kmipbin"

type KmipError struct {
	ErrorCode     kmipbin.KmipEnum
	ResultStatus  kmipbin.KmipEnum
	ResultMessage kmipbin.KmipTextString
}
