package parser

import (
	"io"
	"github.com/rajnikant12345/kmip_g/objects"
	"reflect"
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/enums/resultreason"
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
		return &kmiperror.KmipError{resultreason.OperationFailed, resultreason.InvalidMessage, "Invalid Message structure"}
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
		return &kmiperror.KmipError{resultreason.OperationFailed, resultreason.InvalidMessage, "Message cannot be parsed"}
	}
	b = b[8:]
	err = Dummy(&k, &b)
	if err != nil {
		return &kmiperror.KmipError{resultreason.OperationFailed, resultreason.InvalidMessage, "Invalid Message structure"}
	}
	v.Elem().Field(0).Set(k)
	return nil
}


func Parser(rw io.ReadWriter) *objects.KmipStruct {
	kmipBin, err := TTLVReader(rw)
	if err != nil {
		return nil
	}
	r := objects.KmipStruct{}
	UnmaeshalAllRequest(&r, kmipBin)
	return &r
}

func ResPonseParser(rw io.ReadWriter) *objects.KmipStructResponse {
	kmipBin, err := TTLVReader(rw)
	if err != nil {
		return nil
	}
	r := objects.KmipStructResponse{}
	UnmaeshalAllResponse(&r, kmipBin)
	return &r
}

func MarshalAllRequest(a *objects.KmipStruct) []byte {
	v := reflect.ValueOf(a)
	return DummyMarshal(&v, "")

}

func MarshalAllResponse(a *objects.KmipStructResponse) []byte {
	v := reflect.ValueOf(a)
	return DummyMarshal(&v, "")

}
