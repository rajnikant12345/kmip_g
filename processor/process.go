package processor

import (
	"github.com/pkg/errors"
	"github.com/rajnikant12345/kmip_g/parser"
)

func ProcessPacket( req *parser.KmipStruct , res *parser.KmipStructResponse ) error {
	p := req.GetProtocol()
	if p == nil {
		return errors.New("No Protocol Packet Present")
	}
	batchCounter := req.GetBatchCount()

	if batchCounter == nil {
		return errors.New("Batch Item Not present")
	}

	for i:=0 ;i <  int(*batchCounter) ;i++ {

	}

	return nil
}
