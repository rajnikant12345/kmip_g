package parser

import (
	"io"
	"errors"
	"github.com/rajnikant12345/kmip_g/objects"
)




type KMIPEncoder struct {
	writer io.ReadWriter
}

func NewEncoder(writer io.ReadWriter) (*KMIPEncoder , error ){
	if writer == nil {
		return nil , errors.New("Invalid writer in encoder.")
	}
	k := KMIPEncoder{writer}
	return &k, nil
}

func (e *KMIPEncoder) Encode( kmipResponse *objects.KmipStructResponse ) {
	e.writer.Write(MarshalAllResponse(kmipResponse))
}


func (e *KMIPEncoder) EncodeRequest( kmipRequest *objects.KmipStruct ) {
	e.writer.Write(MarshalAllRequest(kmipRequest))
}
