package parser

import (
	"encoding/xml"
	"fmt"
	"io"
)

func Parser(rw io.ReadWriter) {
	kmipBin, err := TTLVReader(rw)
	if err != nil {
		return
	}
	r := KmipStruct{}
	UnmaeshalAllT(&r, kmipBin)

	//r.Unmarshal(kmipBin[8:])

	x, _ := xml.MarshalIndent(r, "", "  ")

	fmt.Println(string(x))

	v := DataAttributesKmip{}

	if r.GetBatcItem(0).GetRequestPayload().TemplateAttribute != nil {
		for i:=0;i<len(r.GetBatcItem(0).GetRequestPayload().TemplateAttribute.Attribute);i++ {
			v.Unpack(r.GetBatcItem(0).GetRequestPayload().TemplateAttribute.Attribute[i].Data)
		}

	}
	for i:=0;i<len(r.GetBatcItem(0).GetRequestPayload().Attribute);i++ {
		v.Unpack(r.GetBatcItem(0).GetRequestPayload().Attribute[i].Data)
	}

	if r.GetBatcItem(0).GetRequestPayload().Template != nil {
		for i:=0;i<len(r.GetBatcItem(0).GetRequestPayload().Template.Attribute);i++ {
			v.Unpack(r.GetBatcItem(0).GetRequestPayload().Template.Attribute[i].Data)
		}

	}




	//r1 := RequestHeader{}

	//UnmaeshalAll(r1, nil)

}
