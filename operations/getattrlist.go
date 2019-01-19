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



type OpGetAttrList struct {
}


func (op *OpGetAttrList) DoOp(r *objects.KmipStruct, batchNum int , ks *kmipservice.KmipService, idPlaceHolder *kmipbin.KmipTextString) *objects.BatchItem {

	fmt.Println("=====================hitting del attr op====================")

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

	uid , attrs, errs := ks.GetAttributeListCallBack(context.Background() , string(*id) )

	if uid == "" {
		return prepareCreateEroorResponse(errs)
	}


	for i:=0;i<len(attrs);i++ {
		at := kmipbin.KmipTextString(attrs[i])
		resBatch.ResponsePayload.AttributeName = append( resBatch.ResponsePayload.AttributeName ,&at )
	}


	uidk := kmipbin.KmipTextString(uid)
	resultStatus := kmipbin.KmipEnum(resultreason.Success)
	resBatch.ResultStatus = &resultStatus
	resBatch.Operation = batchReq.Operation
	resBatch.UniqueBatchItemID = batchReq.UniqueBatchItemID
	resBatch.ResponsePayload.UniqueIdentifier = append(resBatch.ResponsePayload.UniqueIdentifier, &uidk)

	*idPlaceHolder = ""
	vvv,_ := xml.Marshal(resBatch)

	fmt.Println(string(vvv))

	return &resBatch
}
