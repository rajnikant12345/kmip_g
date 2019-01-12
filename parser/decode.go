package parser

import "io"
import (
	"errors"
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/objects"
)


type KMIPDecoder struct {
	reader io.ReadWriter
}

func NewDecoder(reader io.ReadWriter) (*KMIPDecoder , error ){
	if reader == nil {
		return nil , errors.New("Invalid writer in encoder.")
	}
	k := KMIPDecoder{reader}
	return &k, nil
}

func (e *KMIPDecoder) Decode( kmipRequest *objects.KmipStruct  ) *kmiperror.KmipError {
	kmipData, err := TTLVReader(e.reader)

	if err != nil {
		return &kmiperror.MessageCannotBeParsed
	}
	return UnmaeshalAllRequest(kmipRequest, kmipData)
}


func (e *KMIPDecoder) DecodeResponse( kmipResponse *objects.KmipStructResponse  ) *kmiperror.KmipError {
	kmipData, err := TTLVReader(e.reader)

	if err != nil {
		return &kmiperror.MessageCannotBeParsed
	}
	return UnmaeshalAllResponse(kmipResponse, kmipData)
}
