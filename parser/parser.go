package parser

import (
	"github.com/rajnikant12345/kmip_g/objects"
	"reflect"
	"github.com/rajnikant12345/kmip_g/kmiperror"
)


func UnmaeshalAllRequest(a *objects.KmipStruct, b []byte) *kmiperror.KmipError  {
	v := reflect.ValueOf(a)
	typ := reflect.TypeOf(v.Elem().Field(0).Interface()).Elem()
	k := reflect.New(typ)
	err := validateLength(b)
	if err != nil {
		return err
	}
	b = b[8:]
	err = Dummy(&k, &b)
	if err != nil {
		return &kmiperror.InvalidMessageStructure
	}
	v.Elem().Field(0).Set(k)
	return nil
}

func UnmaeshalAllResponse(a *objects.KmipStructResponse, b []byte)  *kmiperror.KmipError  {
	v := reflect.ValueOf(a)
	typ := reflect.TypeOf(v.Elem().Field(0).Interface()).Elem()
	k := reflect.New(typ)
	err := validateLength(b)
	if err != nil {
		return &kmiperror.MessageCannotBeParsed
	}
	b = b[8:]
	err = Dummy(&k, &b)
	if err != nil {
		return &kmiperror.InvalidMessageStructure
	}
	v.Elem().Field(0).Set(k)
	return nil
}


func MarshalAllRequest(a *objects.KmipStruct) []byte {
	v := reflect.ValueOf(a)
	return DummyMarshal(&v, "")

}

func MarshalAllResponse(a *objects.KmipStructResponse) []byte {
	v := reflect.ValueOf(a)
	return DummyMarshal(&v, "")

}
