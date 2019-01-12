package operations

import (
	"github.com/rajnikant12345/kmip_g/objects"
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

const ProtoMajor = 1
const ProtoMinor = 2

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

func processRequest(k *objects.KmipStruct) ( *objects.KmipStructResponse, *kmiperror.KmipError ) {
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

	for i,batch := range batchList {
		if batch == nil {
			return nil , &kmiperror.InvalidMessageStructure
		}
		operation := batch.Operation
		if operation == nil {
			batch := objects.BatchItem{}
			batch.ResultMessage = &kmiperror.InvalidMessageStructure.ResultMessage
			batch.ResultReason = &kmiperror.InvalidMessageStructure.ResultReason
			batch.ResultStatus = &kmiperror.InvalidMessageStructure.ResultStatus
			res.ResponseMessage.BatchItem = append(res.ResponseMessage.BatchItem , &batch)

		}else {
			op,ok := OpMap[*batch.Operation]
			if !ok {
				err := kmiperror.InvalidMessageStructure
				err.Operation = *batch.Operation
				err.ResultMessage = "Server does not support operation"
				batchres := objects.BatchItem{}
				batchres.ResultMessage = &err.ResultMessage
				batchres.ResultReason = &err.ResultReason
				batchres.ResultStatus = &err.ResultStatus
				batchres.Operation = batch.Operation
				res.ResponseMessage.BatchItem = append(res.ResponseMessage.BatchItem , &batchres)
			} else {
				batchres := op.DoOp(k,i)
				res.ResponseMessage.BatchItem = append(res.ResponseMessage.BatchItem , batchres)
			}
		}
	}
	return &res , nil
}

func MakeKmipResponse(err  *kmiperror.KmipError ) *objects.KmipStructResponse {

	var major, minor kmipbin.KmipInt =  ProtoMajor , ProtoMinor
	res := objects.KmipStructResponse{}
	res.ResponseMessage = &objects.ResponseMessage{}
	res.ResponseMessage.ResponseHeader = &objects.ResponseHeader{}
	res.ResponseMessage.ResponseHeader.ProtocolVersion = &objects.ProtocolVersion{
		&major,
		&minor,
	}

	batch := objects.BatchItem{}

	batch.ResultMessage = &err.ResultMessage
	batch.ResultReason = &err.ResultReason
	batch.ResultStatus = &err.ResultStatus
	res.ResponseMessage.BatchItem = append( res.ResponseMessage.BatchItem , &batch )

	return &res
}

func DoKmip(req *objects.KmipStruct ) *objects.KmipStructResponse {
	res , err := processRequest(req)

	if err != nil  && res == nil {
		return MakeKmipResponse(err)
	}
	return res
}
