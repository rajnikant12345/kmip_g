package processor

import (
	"errors"
	"github.com/rajnikant12345/kmip_g/parser"
	"github.com/rajnikant12345/kmip_g/operations"
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

	for i:=0 ;i < int(*batchCounter) ;i++ {
		Process(req , res , i )
	}

	return nil
}

func Process( req *parser.KmipStruct , res *parser.KmipStructResponse, batchNum int ) error {
	batch :=  req.GetBatcItem(batchNum)
	if batch ==  nil {
		return errors.New("Batch Number Not Present")
	}
	operation := batch.GetOperation()
	if operation == nil {
		return errors.New("Operation Not Present IN KMIP")
	}
	switch *operation {
	case parser.OperationCreate:
		operations.Create(req , res, batchNum )
	default:
		return errors.New("Operation Not Supported")
	}
	return nil
}