package operations

import (
	"errors"
	"github.com/rajnikant12345/kmip_g/parser"
)

func Create(req *parser.KmipStruct , res *parser.KmipStructResponse, batchNum int) error {
	bi := req.GetBatcItem(batchNum)
	if bi == nil {
		errors.New("")
	}
}
