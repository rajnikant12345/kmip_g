package operations

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/objects"
	"fmt"
)

type OpCreate struct {
}

func prepareCreateEroorResponse(kmipError kmiperror.KmipError) *objects.BatchItem {
	resBatch := objects.BatchItem{}
	resBatch.ResultStatus = &kmiperror.InvalidMessageStructure.ResultStatus
	resBatch.ResultMessage = &kmiperror.InvalidMessageStructure.ResultMessage
	resBatch.ResultReason = &kmiperror.InvalidMessageStructure.ResultReason
	return &resBatch
}

func (op *OpCreate) DoOp(r *objects.KmipStruct, batchNum int) *objects.BatchItem {

	var AttributeMap = make(map[kmipbin.KmipTextString]interface{})

	batchReq := r.GetRequestMessage().BatchItem[batchNum]

	if batchReq == nil || batchReq.RequestPayload == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	if batchReq.RequestPayload.ObjectType == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	if batchReq.RequestPayload.TemplateAttribute != nil {
		if len(batchReq.RequestPayload.TemplateAttribute.Attribute) != 0 {
			for _,v := range batchReq.RequestPayload.TemplateAttribute.Attribute {
				_,ok := AttributeMap[*v.AttributeName]
				if ok {
					return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
				}
				AttributeMap[*v.AttributeName] = v.AttributeValue
			}
		}
	}
	if len(batchReq.RequestPayload.Attribute) != 0 {
		for _,v := range batchReq.RequestPayload.Attribute {
			_,ok := AttributeMap[*v.AttributeName]
			if ok {
				return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
			}
			AttributeMap[*v.AttributeName] = v.AttributeValue
		}
	}
	var contactInformation *kmipbin.KmipTextString
	var usageMask *kmipbin.KmipInt
	if val,ok :=  AttributeMap["Contact Information"];ok {
		contactInformation = val.(*kmipbin.KmipTextString)
	}

	if val,ok :=  AttributeMap["Cryptographic Usage Mask"];ok {
		usageMask = val.(*kmipbin.KmipInt)
	}





	resBatch := objects.BatchItem{}
	return &resBatch
}
