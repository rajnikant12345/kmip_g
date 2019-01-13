package operations

import (
	"fmt"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/objects"
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
			for _, v := range batchReq.RequestPayload.TemplateAttribute.Attribute {
				_, ok := AttributeMap[*v.AttributeName]
				if ok && *v.AttributeName != "Name" {
					return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
				}
				AttributeMap[*v.AttributeName] = v.AttributeValue
			}
		}
	}
	if len(batchReq.RequestPayload.Attribute) != 0 {
		for _, v := range batchReq.RequestPayload.Attribute {
			fmt.Println(*v.AttributeName)
			_, ok := AttributeMap[*v.AttributeName]
			if ok && *v.AttributeName != "Name"  {
				return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
			}
			AttributeMap[*v.AttributeName] = v.AttributeValue
		}
	}

	contactInformation := GetAttribute("Contact Information", AttributeMap)
	usageMask := GetAttribute("Cryptographic Usage Mask", AttributeMap)
	activationdate := GetAttribute("Activation Date", AttributeMap)
	processStartDate := GetAttribute("Process Start Date", AttributeMap)
	protectStopDate := GetAttribute("Protect Stop Date", AttributeMap)
	cryptographicAlgorithm := GetAttribute("Cryptographic Algorithm", AttributeMap)
	cryptographicLength := GetAttribute("Cryptographic Length", AttributeMap)

	if contactInformation != nil {
		fmt.Println("Contact Information: ",*contactInformation.(*kmipbin.KmipTextString))
	}

	if usageMask != nil {
		fmt.Println("Usage Mask: ",*usageMask.(*kmipbin.KmipInt))
	}

	if activationdate != nil {
		fmt.Println("Activation Date: ",*activationdate.(*kmipbin.KmipDate))
	}

	if processStartDate != nil {
		fmt.Println("Activation Date: ",*processStartDate.(*kmipbin.KmipDate))
	}

	if protectStopDate != nil {
		fmt.Println("Activation Date: ",*protectStopDate.(*kmipbin.KmipDate))
	}

	if cryptographicAlgorithm != nil {
		fmt.Println("Cryptographic Algorithm: ",*cryptographicAlgorithm.(*kmipbin.KmipEnum))
	}

	if cryptographicLength != nil {
		fmt.Println("Cryptographic Length: ",*cryptographicLength.(*kmipbin.KmipInt))
	}

	fmt.Println("==========================================")


	resBatch := objects.BatchItem{}
	return &resBatch
}
