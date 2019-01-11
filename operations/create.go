package operations

import (
	"github.com/rajnikant12345/kmip_g/objects"
)

type OpCreate struct {

}

func (op * OpCreate) DoOp( r *objects.KmipStruct , batchNum int) *objects.BatchItem {
	return nil
}
