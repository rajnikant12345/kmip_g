package callbacks

import (
	"context"
	"github.com/rajnikant12345/kmip_g/objects"
	"github.com/rajnikant12345/kmip_g/kmiperror"
)


type CreateCallBackType func(context.Context  , []*objects.Attribute ) ( string, kmiperror.KmipError)

type DestroyCallBackType func(context.Context , string ) ( string, kmiperror.KmipError)

type CheckCallBackType func(context.Context , string ) ( string, map[string]interface{} , kmiperror.KmipError)

/////////////////////////////////////////////

type DeleteAttributeCallBack func(context.Context , string , []string) ( string , []*objects.Attribute ,kmiperror.KmipError)

type GetAttributeListCallBack func(context.Context , string ) ( string , []string, kmiperror.KmipError)

type GetAttributesCallBack func(context.Context , string , []string ) ( string , []*objects.Attribute , kmiperror.KmipError)

type ModifyAttributesCallBack func(context.Context , string , []*objects.Attribute  ) ( string , []*objects.Attribute , kmiperror.KmipError)

type LocateCallBack func(context.Context  ,[]*objects.Attribute ) ( []string , kmiperror.KmipError )

type GetKeyCallBack func(context.Context , string , *objects.KeyWrappingSpecification ) ( string , *objects.KeyBlock , kmiperror.KmipError)

type CreateKeyPairCallBack func(context.Context  ,[]*objects.Attribute , []*objects.Attribute )  ( string , string , kmiperror.KmipError)

type ReKeyCallBack func( context.Context , string   ) ( string, kmiperror.KmipError)

type ReKeyKeyPairCallBack func( context.Context , string   ) ( string, string, kmiperror.KmipError)