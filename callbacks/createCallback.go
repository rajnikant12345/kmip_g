package callbacks

import (
	"context"
	"github.com/rajnikant12345/kmip_g/objects"
	"github.com/rajnikant12345/kmip_g/kmiperror"
)


type CreateCallBackType func(context.Context  ,map[string]interface{} ,  []*objects.Name) ( string, kmiperror.KmipError)

type DestroyCallBackType func(context.Context , string ) ( string, kmiperror.KmipError)

type CheckCallBackType func(context.Context , string ) ( string, map[string]interface{} , kmiperror.KmipError)

type DeleteAttributeCallBack func(context.Context , string , []string) ( string , []*objects.Attribute , kmiperror.KmipError)

type GetAttributeListCallBack func(context.Context , string ) ( string , []*objects.Attribute , kmiperror.KmipError)

type GetAttributesCallBack func(context.Context , string , []string ) ( string , []*objects.Attribute , kmiperror.KmipError)

type GetKeyCallBack func(context.Context , string , *objects.KeyWrappingSpecification ) ( string , *objects.KeyBlock , kmiperror.KmipError)

// common attribute, public key attribute, provate key atribute
type CreateKeyPairCallBack func(context.Context  ,map[string]interface{} , map[string]interface{}, map[string]interface{})  ( string , string , kmiperror.KmipError)

type LocateCallBack func(context.Context  ,map[string]interface{} , []*objects.Name ) ( []string , kmiperror.KmipError )

type ModifyAttributesCallBack func(context.Context , string , map[string]interface{}  ) ( string , []*objects.Attribute , kmiperror.KmipError)

type ReKeyCallBack func( context.Context , string   ) ( string, kmiperror.KmipError)

type ReKeyKeyPairCallBack func( context.Context , string   ) ( string, string, kmiperror.KmipError)