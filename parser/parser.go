package parser

import (
	"io"
	"github.com/rajnikant12345/kmip_g/objects"
)


func Parser(rw io.ReadWriter) *objects.KmipStruct {
	kmipBin, err := TTLVReader(rw)
	if err != nil {
		return nil
	}
	r := objects.KmipStruct{}
	UnmaeshalAllRequest(&r, kmipBin)
	return &r
}

func ResPonseParser(rw io.ReadWriter) *objects.KmipStructResponse {
	kmipBin, err := TTLVReader(rw)
	if err != nil {
		return nil
	}
	r := objects.KmipStructResponse{}
	UnmaeshalAllResponse(&r, kmipBin)
	return &r
}


