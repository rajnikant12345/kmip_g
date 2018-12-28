package parser

import (
	"encoding/xml"
	"fmt"
	"io"
)

func Parser(rw io.ReadWriter) {
	kmipBin,err := TTLVReader(rw)
	if err != nil {
		return
	}
	r := RequestMessage{}
	r.Unmarshal(kmipBin[8:])

	x,_ := xml.MarshalIndent(&r,"","  ")

	fmt.Println(string(x))

}