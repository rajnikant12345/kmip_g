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


type OpDelAttr struct {
}


func (op *OpDelAttr) DoOp(r *objects.KmipStruct, batchNum int , ks *kmipservice.KmipService) *objects.BatchItem {

	fmt.Println("=====================hitting del attr op====================")

	batchReq := r.GetRequestMessage().BatchItem[batchNum]

	if batchReq == nil || batchReq.RequestPayload == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	if len(batchReq.RequestPayload.UniqueIdentifier) == 0 {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	if batchReq.RequestPayload.UniqueIdentifier[0] == nil {
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

	var attrlist  []string

	if len(batchReq.RequestPayload.AttributeName) != 0 {
		for i:=0;i<len(batchReq.RequestPayload.AttributeName);i++ {
			if batchReq.RequestPayload.AttributeName[i] != nil {
				attrlist = append(attrlist , string(*batchReq.RequestPayload.AttributeName[i]))
			}
		}
	}

	for k, _ := range AttributeMap {
		attrlist = append(attrlist , k)
	}

	resBatch := objects.BatchItem{}
	resBatch.ResponsePayload = &objects.ResponsePayload{}

	uid ,attr , errs := ks.DeleteAttributeCallBack(context.Background() , string(*batchReq.RequestPayload.UniqueIdentifier[0]), attrlist )

	if uid != "" {
		return prepareCreateEroorResponse(errs)
	}

	if len(attr) != 0 {
		resBatch.ResponsePayload.Attribute = attr
	}

	uidk := kmipbin.KmipTextString(uid)
	resultStatus := kmipbin.KmipEnum(resultreason.Success)
	resBatch.ResultStatus = &resultStatus
	resBatch.Operation = batchReq.Operation
	resBatch.UniqueBatchItemID = batchReq.UniqueBatchItemID
	resBatch.ResponsePayload.UniqueIdentifier = append(resBatch.ResponsePayload.UniqueIdentifier, &uidk)
	return &resBatch
}
