package operations

import (
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/objects"
	"fmt"
	"context"
	"github.com/rajnikant12345/kmip_g/enums/resultreason"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"github.com/rajnikant12345/kmip_g/kmipservice"
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

	var AttributeMap = make(map[string]interface{})

	var NameList []*objects.Name

	batchReq := r.GetRequestMessage().BatchItem[batchNum]

	if batchReq == nil || batchReq.RequestPayload == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	if batchReq.RequestPayload.ObjectType == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	AttributeMap, NameList, err := ReadTemplateAttributes( batchReq.RequestPayload.TemplateAttribute)
	if err != nil {
		return prepareCreateEroorResponse(*err)
	}


	AttributeMap1, NameList1, err := ReadAttributes( batchReq.RequestPayload.Attribute)
	if err != nil {
		return prepareCreateEroorResponse(*err)
	}

	NameList = append(NameList,NameList1...)

	for k, v := range AttributeMap1 {
		AttributeMap[k] = v
	}


	fmt.Println("=================Create operation=========================")


	resBatch := objects.BatchItem{}
	resBatch.ResponsePayload = &objects.ResponsePayload{}

	uid , errs := ks.CreateCallBack(context.Background(), AttributeMap , NameList)

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
