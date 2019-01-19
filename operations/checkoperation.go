package operations

import (
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/objects"
	"context"
	"github.com/rajnikant12345/kmip_g/enums/resultreason"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"fmt"
	"github.com/rajnikant12345/kmip_g/kmipservice"
	"github.com/rajnikant12345/kmip_g/enums/operation"
)

type OpCheck struct {
}


func (op *OpCheck) DoOp(r *objects.KmipStruct, batchNum int , ks *kmipservice.KmipService, idPlaceHolder *kmipbin.KmipTextString ) *objects.BatchItem {

	fmt.Println("=====================hitting check op====================")

	batchReq := r.GetRequestMessage().BatchItem[batchNum]

	if batchReq == nil || batchReq.RequestPayload == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	id := ReadUniqueIdOfPayLoad(batchReq.RequestPayload , idPlaceHolder)
	if id == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	resBatch := objects.BatchItem{}
	resBatch.ResponsePayload = &objects.ResponsePayload{}


	uid , attrmap , errs := ks.CheckCallBack(context.Background() , string(*id) )

	if uid == "" {
		return prepareCreateEroorResponse(errs)
	}

	_ , ok := attrmap["Cryptographic Usage Mask"]

	if !ok {
		err := kmiperror.IllegalOperationStructure
		err.Operation = operation.Check
		err.ResultMessage = "Check cannot be performed: object has no 'CryptographicUsageMask'"
		return prepareCreateEroorResponse(err)
	}


	if batchReq.RequestPayload.CryptographicUsageMask != nil {
		mask := int(*batchReq.RequestPayload.CryptographicUsageMask)
		usageMaks := attrmap["Cryptographic Usage Mask"].(int)
		if (mask & usageMaks) == 0 {
			err := kmiperror.PermissionDenied
			err.Operation = operation.Check
			err.ResultMessage = "'Cryptographic Usage Mask' prohibited"
			br := prepareCreateEroorResponse(err)
			br.ResponsePayload = &objects.ResponsePayload{}
			br.ResponsePayload.CryptographicUsageMask = batchReq.RequestPayload.CryptographicUsageMask
			return br
		}

	}

	uidk := kmipbin.KmipTextString(uid)
	resultStatus := kmipbin.KmipEnum(resultreason.Success)
	resBatch.ResultStatus = &resultStatus
	resBatch.Operation = batchReq.Operation
	resBatch.UniqueBatchItemID = batchReq.UniqueBatchItemID
	*idPlaceHolder = ""
	resBatch.ResponsePayload.UniqueIdentifier = append(resBatch.ResponsePayload.UniqueIdentifier, &uidk)
	return &resBatch
}
