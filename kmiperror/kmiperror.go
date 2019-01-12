package kmiperror

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"github.com/rajnikant12345/kmip_g/enums/resultreason"
)

type KmipError struct {
	ResultReason  kmipbin.KmipEnum
	ResultStatus  kmipbin.KmipEnum
	ResultMessage kmipbin.KmipTextString
	Operation     kmipbin.KmipEnum
}


var MessageCannotBeParsed = KmipError{resultreason.InvalidMessage, resultreason.OperationFailed, "Message cannot be parsed", 0}
var InvalidMessageStructure = KmipError{resultreason.InvalidMessage, resultreason.OperationFailed, "Invalid Message structure" , 0}