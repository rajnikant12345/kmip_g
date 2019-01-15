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


type OpDestroy struct {
}


func (op *OpDestroy) DoOp(r *objects.KmipStruct, batchNum int , ks *kmipservice.KmipService) *objects.BatchItem {

	fmt.Println("=====================hitting destroy====================")

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


	resBatch := objects.BatchItem{}
	resBatch.ResponsePayload = &objects.ResponsePayload{}


	uid , err := ks.DestroyCallBack(context.Background() , string(*batchReq.RequestPayload.UniqueIdentifier[0]) )

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
