package operations

import (
	"github.com/rajnikant12345/kmip_g/objects"
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

func ReadUniqueIdOfPayLoad(payLoad *objects.RequestPayload , idPlaceHolder *kmipbin.KmipTextString  ) *kmipbin.KmipTextString {
	if len(payLoad.UniqueIdentifier) == 0 && payLoad.UniqueIdentifier[0] == nil && *idPlaceHolder == "" {
		return nil
	}
	if payLoad.UniqueIdentifier[0] != nil {
		return payLoad.UniqueIdentifier[0]
	}

	if *idPlaceHolder != "" {
		return idPlaceHolder
	}
	return nil

}

