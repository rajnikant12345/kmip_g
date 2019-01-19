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



type OpModAttr struct {
}


func (op *OpModAttr) DoOp(r *objects.KmipStruct, batchNum int , ks *kmipservice.KmipService, idPlaceHolder *kmipbin.KmipTextString) *objects.BatchItem {

	fmt.Println("=====================hitting del attr op====================")

	batchReq := r.GetRequestMessage().BatchItem[batchNum]

	if batchReq == nil || batchReq.RequestPayload == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	id := ReadUniqueIdOfPayLoad(batchReq.RequestPayload , idPlaceHolder)
	if id == nil {
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

	uid , attr, errs := ks.ModifyAttributesCallBack(context.Background() , string(*id), attrListTemplate )

	if uid == "" {
		return prepareCreateEroorResponse(errs)
	}


	uidk := kmipbin.KmipTextString(uid)
	resultStatus := kmipbin.KmipEnum(resultreason.Success)
	resBatch.ResultStatus = &resultStatus
	resBatch.Operation = batchReq.Operation
	resBatch.UniqueBatchItemID = batchReq.UniqueBatchItemID
	resBatch.ResponsePayload.Attribute = attr
	resBatch.ResponsePayload.UniqueIdentifier = append(resBatch.ResponsePayload.UniqueIdentifier, &uidk)

	*idPlaceHolder = ""
	vvv,_ := xml.Marshal(resBatch)

	fmt.Println(string(vvv))

	return &resBatch
}
