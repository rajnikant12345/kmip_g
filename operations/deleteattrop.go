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



type OpDelAttr struct {
}


func (op *OpDelAttr) DoOp(r *objects.KmipStruct, batchNum int , ks *kmipservice.KmipService, idPlaceHolder *kmipbin.KmipTextString) *objects.BatchItem {

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


	var attrlist  []string

	if len(batchReq.RequestPayload.AttributeName) != 0 {
		for i:=0;i<len(batchReq.RequestPayload.AttributeName);i++ {
			if batchReq.RequestPayload.AttributeName[i] != nil {
				attrlist = append(attrlist , string(*batchReq.RequestPayload.AttributeName[i]))
			}
		}
	}

	for i:=0;i<len(attrListTemplate);i++ {
		if attrListTemplate[i].AttributeName != nil {
			attrlist = append(attrlist , string(*attrListTemplate[i].AttributeName))
		}
	}

	resBatch := objects.BatchItem{}
	resBatch.ResponsePayload = &objects.ResponsePayload{}

	uid , attr, errs := ks.DeleteAttributeCallBack(context.Background() , string(*id), attrlist )

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
