package server

import (
	"github.com/rajnikant12345/kmip_g/objects"
	"github.com/rajnikant12345/kmip_g/operations"
	"github.com/rajnikant12345/kmip_g/parser"
	"io"
)

func KmipLoop(conn io.ReadWriter) {
	var res *objects.KmipStructResponse
	dec,_ := parser.NewDecoder(conn)
	req := objects.KmipStruct{}
	err := dec.Decode(&req)

	if err != nil {
		res = operations.MakeKmipResponse(err)
	}else {
		res = operations.DoKmip(&req)
	}

	encd,_ := parser.NewEncoder(conn)

	encd.Encode(res)

}