package operations

import (
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/objects"
	"fmt"
	"context"
	"github.com/rajnikant12345/kmip_g/enums/resultreason"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"github.com/rajnikant12345/kmip_g/kmipservice"
	"github.com/rajnikant12345/kmip_g/enums/keyenums"
	"github.com/rajnikant12345/kmip_g/enums/operation"
)


type OpCreate struct {
}

func prepareCreateEroorResponse(kmipError kmiperror.KmipError) *objects.BatchItem {
	resBatch := objects.BatchItem{}
	resBatch.ResultStatus = &kmipError.ResultStatus
	resBatch.ResultMessage = &kmipError.ResultMessage
	resBatch.ResultReason = &kmipError.ResultReason
	if kmipError.Operation != 0 {
		resBatch.Operation = &kmipError.Operation
	}
	return &resBatch
}

func (op *OpCreate) DoOp(r *objects.KmipStruct, batchNum int, ks *kmipservice.KmipService, idPlaceHolder *kmipbin.KmipTextString) *objects.BatchItem {


	batchReq := r.GetRequestMessage().BatchItem[batchNum]

	if batchReq == nil || batchReq.RequestPayload == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	if batchReq.RequestPayload.ObjectType == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	if *batchReq.RequestPayload.ObjectType !=  keyenums.SymmetricKey {
		err := kmiperror.FeatureNotSupported
		err.Operation = operation.Create
		err.ResultMessage = "Invalid object type"
		return prepareCreateEroorResponse(kmiperror.FeatureNotSupported)
	}


	var attrList []*objects.Attribute

	if  batchReq.RequestPayload.TemplateAttribute != nil {
		attrList = batchReq.RequestPayload.TemplateAttribute.Attribute
	}

	if batchReq.RequestPayload.Attribute != nil {
		attrList = append(attrList , batchReq.RequestPayload.Attribute...)
	}


	fmt.Println("=================Create operation=========================")


	resBatch := objects.BatchItem{}
	resBatch.ResponsePayload = &objects.ResponsePayload{}

	uid , errs := ks.CreateCallBack(context.Background(), attrList)

	if uid == "" {
		return prepareCreateEroorResponse(errs)
	}

	uidk := kmipbin.KmipTextString(uid)
	resultStatus := kmipbin.KmipEnum(resultreason.Success)
	resBatch.ResultStatus = &resultStatus
	resBatch.Operation = batchReq.Operation
	resBatch.UniqueBatchItemID = batchReq.UniqueBatchItemID
	*idPlaceHolder = uidk
	resBatch.ResponsePayload.UniqueIdentifier = append(resBatch.ResponsePayload.UniqueIdentifier, &uidk)
	return &resBatch
}
