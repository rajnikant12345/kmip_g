package operations

import (
	"github.com/rajnikant12345/kmip_g/objects"
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

func isPrptocolSupported( version objects.ProtocolVersion ) ( bool , *kmiperror.KmipError )  {
	if version.ProtocolVersionMajor == nil || version.ProtocolVersionMinor == nil {
		return false, &kmiperror.MessageCannotBeParsed
	}
	if *version.ProtocolVersionMajor != 1 {
		return false , &kmiperror.InvalidMessageStructure
	}

	if *version.ProtocolVersionMinor == 0 {
		return true , nil
	}

	if *version.ProtocolVersionMinor == 1 {
		return true , nil
	}

	if *version.ProtocolVersionMinor == 2 {
		return true , nil
	}
	*version.ProtocolVersionMinor = 2
	return true , nil
}

func ProcessRequest(k *objects.KmipStruct) ( *objects.KmipStructResponse, *kmiperror.KmipError ) {
	if k == nil {
		return nil , &kmiperror.MessageCannotBeParsed
	}
	requestMessage := k.GetRequestMessage()
	if requestMessage == nil {
		return nil , &kmiperror.MessageCannotBeParsed
	}
	requestHeader := requestMessage.GetRequestHeader()
	if requestHeader == nil {
		return nil , &kmiperror.MessageCannotBeParsed
	}

	if requestHeader.ProtocolVersion == nil {
		return nil , &kmiperror.MessageCannotBeParsed
	}
	_,err  := isPrptocolSupported(*requestHeader.ProtocolVersion)
	if err != nil {
		return nil, err
	}
	batchList := requestMessage.GetBatchList()
	if batchList == nil {
		return nil , &kmiperror.MessageCannotBeParsed
	}

	res := objects.KmipStructResponse{}
	res.ResponseMessage = &objects.ResponseMessage{}
	res.ResponseMessage.ResponseHeader = &objects.ResponseHeader{}
	res.ResponseMessage.ResponseHeader.ProtocolVersion = requestHeader.ProtocolVersion

	batchLength := kmipbin.KmipInt(len(batchList))
	res.ResponseMessage.ResponseHeader.SetBatchCount(batchLength)

	for _,batch := range batchList {
		if batch == nil {
			return nil , &kmiperror.MessageCannotBeParsed
		}
		operation := batch.Operation
		if operation == nil {

		}
	}


	return nil , nil
}
