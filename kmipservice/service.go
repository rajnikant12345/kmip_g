package kmipservice

import "github.com/rajnikant12345/kmip_g/callbacks"

type KmipService struct {
	CreateCallBack  callbacks.CreateCallBackType
	DestroyCallBack callbacks.DestroyCallBackType
	CheckCallBack   callbacks.CheckCallBackType
	DeleteAttributeCallBack	callbacks.DeleteAttributeCallBack
}
