package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type BatchItem struct {
	Operation         *kmipbin.KmipEnum       `kmip:"42005C"`
	UniqueBatchItemID *kmipbin.KmipByteString `kmip:"420093"`
	ResultStatus	  *kmipbin.KmipEnum
	ResultReason	  *kmipbin.KmipEnum
	ResultMessage	  *kmipbin.KmipTextString
	AsynchronousCorrelationValue	*kmipbin.KmipByteString
	ResponsePayload		*ResponsePayload
	RequestPayload    *RequestPayload         `kmip:"420079"`
	MessageExtension  *MessageExtension       `kmip:"420051"`
}
