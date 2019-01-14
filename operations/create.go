package operations

import (
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/objects"
	"fmt"
	"github.com/rajnikant12345/kmip_g/callbacks"
	"context"
	"github.com/rajnikant12345/kmip_g/enums/resultreason"
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

/*


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

 */


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

	var AttributeMap = make(map[string]interface{})

	var NameList []*objects.Name

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
				name := string(*v.AttributeName)
				_, ok := AttributeMap[name]
				if ok && *v.AttributeName != "Name" {
					return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
				}

				if *v.AttributeName == "Name" {
					NameList = append(NameList , v.AttributeValue.(*objects.Name))
				}else {
					AttributeMap[name] = v.AttributeValue
				}
			}
		}
	}
	if len(batchReq.RequestPayload.Attribute) != 0 {
		for _, v := range batchReq.RequestPayload.Attribute {
			fmt.Println(*v.AttributeName)
			name := string(*v.AttributeName)
			_, ok := AttributeMap[name]
			if ok && *v.AttributeName != "Name"  {
				return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
			}

			if *v.AttributeName == "Name" {
				NameList = append(NameList , v.AttributeValue.(*objects.Name))
			} else {
				AttributeMap[name] = v.AttributeValue
			}
		}
	}

	fmt.Println("==========================================")


	resBatch := objects.BatchItem{}
	resBatch.ResponsePayload = &objects.ResponsePayload{}

	uid , err := callbacks.CreateCallBack(context.Background(), AttributeMap , NameList)

	if err != nil {
		er := kmiperror.InvalidMessageStructure
		er.ResultReason = resultreason.CryptographicFailure
		er.ResultMessage = kmipbin.KmipTextString(err.Error())
		return prepareCreateEroorResponse(er)
	}

	uidk := kmipbin.KmipTextString(uid)
	resultStatus := kmipbin.KmipEnum(resultreason.Success)
	resBatch.ResultStatus = &resultStatus
	resBatch.Operation = batchReq.Operation
	resBatch.ResponsePayload.UniqueIdentifier = append(resBatch.ResponsePayload.UniqueIdentifier, &uidk)
	return &resBatch
}
