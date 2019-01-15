package operations

import (
	"github.com/rajnikant12345/kmip_g/objects"
	"github.com/rajnikant12345/kmip_g/kmipservice"
)

type Operation interface {
	DoOp( r *objects.KmipStruct , batchNum int, ks *kmipservice.KmipService) *objects.BatchItem
}
