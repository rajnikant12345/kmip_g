package operations

import "github.com/rajnikant12345/kmip_g/objects"

type Operation interface {
	DoOp( r *objects.KmipStruct , batchNum int) *objects.BatchItem
}
