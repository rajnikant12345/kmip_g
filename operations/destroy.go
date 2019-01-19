package operations

import (
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/objects"
	"context"
	"github.com/rajnikant12345/kmip_g/enums/resultreason"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"fmt"
	"github.com/rajnikant12345/kmip_g/kmipservice"
)


type OpDestroy struct {
}


func (op *OpDestroy) DoOp(r *objects.KmipStruct, batchNum int , ks *kmipservice.KmipService, idPlaceHolder *kmipbin.KmipTextString) *objects.BatchItem {

	fmt.Println("=====================hitting destroy====================")

	batchReq := r.GetRequestMessage().BatchItem[batchNum]

	if batchReq == nil || batchReq.RequestPayload == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	id := ReadUniqueIdOfPayLoad(batchReq.RequestPayload , idPlaceHolder)
	if id == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	batchReq.RequestPayload.UniqueIdentifier = append( batchReq.RequestPayload.UniqueIdentifier , id)

	resBatch := objects.BatchItem{}
	resBatch.ResponsePayload = &objects.ResponsePayload{}


	uid , err := ks.DestroyCallBack(context.Background() , string(*batchReq.RequestPayload.UniqueIdentifier[0]) )

	if uid == "" {
		return prepareCreateEroorResponse(err)
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
