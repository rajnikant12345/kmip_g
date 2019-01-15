package callbacks

import (
	"context"
	"github.com/rajnikant12345/kmip_g/objects"
	"github.com/rajnikant12345/kmip_g/kmiperror"
)


type CreateCallBackType func(context.Context  ,map[string]interface{} ,  []*objects.Name) ( string, kmiperror.KmipError)

type DestroyCallBackType func(context.Context , string ) ( string, kmiperror.KmipError)