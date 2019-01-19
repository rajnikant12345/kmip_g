package operations

import (
	"github.com/rajnikant12345/kmip_g/enums/operation"
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

var OpMap map[kmipbin.KmipEnum]Operation

func init() {
	OpMap = make(map[kmipbin.KmipEnum]Operation)
	OpMap[operation.Create] = &OpCreate{}
	OpMap[operation.Destroy] = &OpDestroy{}
	OpMap[operation.Check] = &OpCheck{}
	OpMap[operation.DeleteAttribute] = &OpDelAttr{}
	OpMap[operation.GetAttributeList] = &OpGetAttrList{}
	OpMap[operation.ModifyAttribute] = &OpModAttr{}
	OpMap[operation.GetAttributes] = &OpGetAttr{}
	OpMap[operation.Locate] = &OpLocate{}


}
