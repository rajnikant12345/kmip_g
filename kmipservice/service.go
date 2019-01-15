package kmipservice

import "github.com/rajnikant12345/kmip_g/callbacks"

type KmipService struct {
	CreateCallBack           callbacks.CreateCallBackType
	DestroyCallBack          callbacks.DestroyCallBackType
	CheckCallBack            callbacks.CheckCallBackType
	DeleteAttributeCallBack  callbacks.DeleteAttributeCallBack
	GetAttributeListCallBack callbacks.GetAttributeListCallBack
	GetAttributesCallBack    callbacks.GetAttributesCallBack
	GetKeyCallBack           callbacks.GetKeyCallBack
	CreateKeyPairCallBack    callbacks.CreateKeyPairCallBack
	LocateCallBack           callbacks.LocateCallBack
	ModifyAttributesCallBack callbacks.ModifyAttributesCallBack
	ReKeyCallBack            callbacks.ReKeyCallBack
	ReKeyKeyPairCallBack     callbacks.ReKeyKeyPairCallBack
}
