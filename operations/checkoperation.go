package operations

import (
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/objects"
	"context"
	"github.com/rajnikant12345/kmip_g/enums/resultreason"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"fmt"
	"github.com/rajnikant12345/kmip_g/kmipservice"
	"github.com/rajnikant12345/kmip_g/enums/operation"
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


type OpCheck struct {
}


func (op *OpCheck) DoOp(r *objects.KmipStruct, batchNum int , ks *kmipservice.KmipService, idPlaceHolder *kmipbin.KmipTextString ) *objects.BatchItem {

	fmt.Println("=====================hitting check op====================")

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


	uid , attrmap , errs := ks.CheckCallBack(context.Background() , string(*id) )

	if errs.ResultReason != 0 {
		return prepareCreateEroorResponse(errs)
	}

	_ , ok := attrmap["Cryptographic Usage Mask"]

	if !ok {
		err := kmiperror.IllegalOperationStructure
		err.Operation = operation.Check
		err.ResultMessage = "Check cannot be performed: object has no 'CryptographicUsageMask'"
		return prepareCreateEroorResponse(err)
	}


	if batchReq.RequestPayload.CryptographicUsageMask != nil {
		mask := int(*batchReq.RequestPayload.CryptographicUsageMask)
		usageMaks := attrmap["Cryptographic Usage Mask"].(int)
		fmt.Println(mask , usageMaks)
		if (mask & usageMaks) == 0 {
			err := kmiperror.PermissionDenied
			err.Operation = operation.Check
			err.ResultMessage = "'Cryptographic Usage Mask' prohibited"
			return prepareCreateEroorResponse(err)
		}

	}

	uidk := kmipbin.KmipTextString(uid)
	resultStatus := kmipbin.KmipEnum(resultreason.Success)
	resBatch.ResultStatus = &resultStatus
	resBatch.Operation = batchReq.Operation
	resBatch.UniqueBatchItemID = batchReq.UniqueBatchItemID
	*idPlaceHolder = ""
	resBatch.ResponsePayload.UniqueIdentifier = append(resBatch.ResponsePayload.UniqueIdentifier, &uidk)
	return &resBatch
}
