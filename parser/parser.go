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


	//r1 := RequestHeader{}

	//UnmaeshalAll(r1, nil)


}
