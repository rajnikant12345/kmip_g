package operations

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/objects"
)

type OpCreate struct {
}

type CreateObjectAttributes struct {
	CryptographicAlgorithm kmipbin.KmipEnum
	KeyLength              kmipbin.KmipInt
	CryptographicUsageMask kmipbin.KmipInt
	CustomAttributes       map[string]string
}


func fillAttributes(co *CreateObjectAttributes , attributes []*objects.Attribute) *kmiperror.KmipError {
	return nil
}

func prepareCreateEroorResponse(kmipError kmiperror.KmipError) *objects.BatchItem {
	resBatch := objects.BatchItem{}
	resBatch.ResultStatus = &kmiperror.InvalidMessageStructure.ResultStatus
	resBatch.ResultMessage = &kmiperror.InvalidMessageStructure.ResultMessage
	resBatch.ResultReason = &kmiperror.InvalidMessageStructure.ResultReason
	return &resBatch
}

func (op *OpCreate) DoOp(r *objects.KmipStruct, batchNum int) *objects.BatchItem {
	batchReq := r.GetRequestMessage().BatchItem[batchNum]

	if batchReq == nil || batchReq.RequestPayload == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	if batchReq.RequestPayload.ObjectType == nil {
		return prepareCreateEroorResponse(kmiperror.InvalidMessageStructure)
	}

	attributes := &CreateObjectAttributes{}

	if batchReq.RequestPayload.TemplateAttribute != nil {
		fillAttributes(attributes , batchReq.RequestPayload.TemplateAttribute.Attribute)
	}

	if len(batchReq.RequestPayload.Attribute) != 0 {
		fillAttributes(attributes , batchReq.RequestPayload.Attribute)
	}

	resBatch := objects.BatchItem{}
	return &resBatch
}
