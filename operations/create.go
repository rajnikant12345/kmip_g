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

func (op *OpCreate) DoOp(r *objects.KmipStruct, batchNum int, ks *kmipservice.KmipService) *objects.BatchItem {

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
					return prepareCreateEroorResponse(kmiperror.CreteObjectErrorMultipleInstance)
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
				return prepareCreateEroorResponse(kmiperror.CreteObjectErrorMultipleInstance)
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

	uid , err := ks.CreateCallBack(context.Background(), AttributeMap , NameList)

	if err.ResultMessage != "" {
		return prepareCreateEroorResponse(err)
	}

	uidk := kmipbin.KmipTextString(uid)
	resultStatus := kmipbin.KmipEnum(resultreason.Success)
	resBatch.ResultStatus = &resultStatus
	resBatch.Operation = batchReq.Operation
	resBatch.ResponsePayload.UniqueIdentifier = append(resBatch.ResponsePayload.UniqueIdentifier, &uidk)
	return &resBatch
}
