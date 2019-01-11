package kmiperror

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"github.com/rajnikant12345/kmip_g/enums/resultreason"
)

type KmipError struct {
	ErrorCode     kmipbin.KmipEnum
	ResultStatus  kmipbin.KmipEnum
	ResultMessage kmipbin.KmipTextString
	Operation     kmipbin.KmipEnum
}


var MessageCannotBeParsed = KmipError{resultreason.OperationFailed, resultreason.InvalidMessage, "Message cannot be parsed", -1}
var InvalidMessageStructure = KmipError{resultreason.OperationFailed, resultreason.InvalidMessage, "Invalid Message structure" , -1}