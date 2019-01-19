package operations

import (
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/objects"
	"context"
	"github.com/rajnikant12345/kmip_g/enums/resultreason"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"fmt"
	"github.com/rajnikant12345/kmip_g/kmipservice"
	"encoding/xml"
)



type OpLocate struct {
}


func (op *OpLocate) DoOp(r *objects.KmipStruct, batchNum int , ks *kmipservice.KmipService, idPlaceHolder *kmipbin.KmipTextString) *objects.BatchItem {

	fmt.Println("=====================hitting locate op====================")

	batchReq := r.GetRequestMessage().BatchItem[batchNum]

	if batchReq == nil || batchReq.RequestPayload == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}


	var attrListTemplate []*objects.Attribute

	if  batchReq.RequestPayload.TemplateAttribute != nil {
		attrListTemplate = batchReq.RequestPayload.TemplateAttribute.Attribute
	}

	if batchReq.RequestPayload.Attribute != nil {
		attrListTemplate = append(attrListTemplate , batchReq.RequestPayload.Attribute...)
	}


	resBatch := objects.BatchItem{}
	resBatch.ResponsePayload = &objects.ResponsePayload{}

	uid , errs := ks.LocateCallBack(context.Background(), attrListTemplate )

	if len(uid) == 0 {
		return prepareCreateEroorResponse(errs)
	}


	resultStatus := kmipbin.KmipEnum(resultreason.Success)
	resBatch.ResultStatus = &resultStatus
	resBatch.Operation = batchReq.Operation
	resBatch.UniqueBatchItemID = batchReq.UniqueBatchItemID

	for _,i:= range uid {
		p := kmipbin.KmipTextString(i)
		resBatch.ResponsePayload.UniqueIdentifier = append(resBatch.ResponsePayload.UniqueIdentifier, &p)
	}

	*idPlaceHolder = *resBatch.ResponsePayload.UniqueIdentifier[0]
	vvv,_ := xml.Marshal(resBatch)

	fmt.Println(string(vvv))

	return &resBatch
}
