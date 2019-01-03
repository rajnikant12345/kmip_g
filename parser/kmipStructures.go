package parser

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

type RequestMessage struct {
	RequestHeader *RequestHeader `kmip:"420077"`
	BatchItem     []*BatchItem   `kmip:"42000F"`
}

type KmipStruct struct {
	RequestMessage  *RequestMessage  `kmip:"420078"`
	ResponseMessage *ResponseMessage `kmip:"42007B"`
}

type KmipStructResponse struct {
	ResponseMessage *ResponseMessage `kmip:"42007B"`
}

type ResponseMessage struct {
	ResponseHeader *ResponseHeader `kmip:"42007A"`
	BatchItem      []*BatchItem    `kmip:"42000F"`
}

func (k *KmipStruct) GetProtocol() *ProtocolVersion {
	if k.RequestMessage == nil {
		return nil
	}
	if k.RequestMessage.RequestHeader == nil {
		return nil
	}
	return k.RequestMessage.RequestHeader.ProtocolVersion
}

func (k *KmipStruct) GetMaximumResponseSize() *kmipbin.KmipInt {
	if k.RequestMessage == nil {
		return nil
	}
	if k.RequestMessage.RequestHeader == nil {
		return nil
	}
	return k.RequestMessage.RequestHeader.MaximumResponseSize
}

func (k *KmipStruct) GetAsynchronousIndicator() *kmipbin.KmipBoolean {
	if k.RequestMessage == nil {
		return nil
	}
	if k.RequestMessage.RequestHeader == nil {
		return nil
	}
	return k.RequestMessage.RequestHeader.AsynchronousIndicator
}

func (k *KmipStruct) GetAttestationCapableIndicator() *kmipbin.KmipBoolean {
	if k.RequestMessage == nil {
		return nil
	}
	if k.RequestMessage.RequestHeader == nil {
		return nil
	}
	return k.RequestMessage.RequestHeader.AttestationCapableIndicator
}

func (k *KmipStruct) GetAttestationType() *kmipbin.KmipEnum {
	if k.RequestMessage == nil {
		return nil
	}
	if k.RequestMessage.RequestHeader == nil {
		return nil
	}
	return k.RequestMessage.RequestHeader.AttestationType
}

func (k *KmipStruct) GetAuthentication() *Authentication {
	if k.RequestMessage == nil {
		return nil
	}
	if k.RequestMessage.RequestHeader == nil {
		return nil
	}
	return k.RequestMessage.RequestHeader.Authentication
}

func (k *KmipStruct) GetBatchErrorContinuationOption() *kmipbin.KmipEnum {
	if k.RequestMessage == nil {
		return nil
	}
	if k.RequestMessage.RequestHeader == nil {
		return nil
	}
	return k.RequestMessage.RequestHeader.BatchErrorContinuationOption
}

func (k *KmipStruct) GetBatchOrderOption() *kmipbin.KmipBoolean {
	if k.RequestMessage == nil {
		return nil
	}
	if k.RequestMessage.RequestHeader == nil {
		return nil
	}
	return k.RequestMessage.RequestHeader.BatchOrderOption
}

func (k *KmipStruct) GetBatchCount() *kmipbin.KmipInt {
	if k.RequestMessage == nil {
		return nil
	}
	if k.RequestMessage.RequestHeader == nil {
		return nil
	}
	return k.RequestMessage.RequestHeader.BatchCount
}

func (k *KmipStruct) GetTimeStamp() *kmipbin.KmipDate {
	if k.RequestMessage == nil {
		return nil
	}
	if k.RequestMessage.RequestHeader == nil {
		return nil
	}
	return k.RequestMessage.RequestHeader.TimeStamp
}

func (k *KmipStruct) GetBatcItem(index int) *BatchItem {
	if k.RequestMessage == nil {
		return nil
	}
	if index > len(k.RequestMessage.BatchItem)-1 || index < 0 {
		return nil
	}
	return k.RequestMessage.BatchItem[index]
}
