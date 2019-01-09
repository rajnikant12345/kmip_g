package parser

import (
	"io"
	"github.com/rajnikant12345/kmip_g/objects"
)

type KmipStruct struct {
	RequestMessage *objects.RequestMessage
}

func Parser(rw io.ReadWriter) *KmipStruct {
	kmipBin, err := TTLVReader(rw)
	if err != nil {
		return nil
	}

	r := KmipStruct{}
	UnmaeshalAllRequest(&r, kmipBin)
	//r.Unmarshal(kmipBin[8:])
	//xml.MarshalIndent(r, "", "  ")
	//fmt.Println(string(x))
	return &r
}

func ResPonseParser(rw io.ReadWriter) *objects.KmipStructResponse {
	kmipBin, err := TTLVReader(rw)
	if err != nil {
		return nil
	}

	r := objects.KmipStructResponse{}
	UnmaeshalAllResponse(&r, kmipBin)

	//r.Unmarshal(kmipBin[8:])

	//x, _ := xml.MarshalIndent(r, "", "  ")

	//fmt.Println(string(x))

	return &r
}


