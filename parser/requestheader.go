package parser

import (
	"fmt"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"reflect"
)

type RequestHeader struct {
	ProtocolVersion              *ProtocolVersion     `kmip:"420069"`
	MaximumResponseSize          *kmipbin.KmipInt     `kmip:"420050"`
	AsynchronousIndicator        *kmipbin.KmipBoolean `kmip:"420007"`
	AttestationCapableIndicator  *kmipbin.KmipBoolean `kmip:"4200D3"`
	AttestationType              *kmipbin.KmipEnum    `kmip:"4200C7"`
	Authentication               *Authentication      `kmip:"42000C"`
	BatchErrorContinuationOption *kmipbin.KmipEnum    `kmip:"42000E"`
	BatchOrderOption             *kmipbin.KmipBoolean `kmip:"420010"`
	TimeStamp                    *kmipbin.KmipDate    `kmip:"420092"`
	BatchCount                   *kmipbin.KmipInt     `kmip:"42000D"`
}

func (r *RequestHeader) Unmarshalrt(b []byte) ([]byte, error) {

	//var tagmap map[string]reflect.Value

	t := reflect.ValueOf(*r)

	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		fmt.Println(field.Type().String())
		//if field.Type().String() == "*parser.ProtocolVersion" {

		c := reflect.New(field.Type().Elem())
		//c := field.Elem().Interface().(*ProtocolVersion)
		fmt.Println(c.Elem().Interface())

	}
	//		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.String(), tag)
	return nil, nil

	/*for {
	str := hex.EncodeToString(b[:4])
	t,ok := tagmap[str]
	if !ok {
		return b, nil
	}
	*/
	//	}

	//for {
	//	if len(b) <= 0 {
	//		return b, nil
	//	}
	//	tag := kmipbin.KmipTagType{}
	//	tag.UnMarshalBin(b[:4])
	/*t := reflect.TypeOf(*r)

	// Get the type and kind of our user variable
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		// Get the field tag value
		//tag := field.Tag.Get("kmip")

		if field.Type.String() == "*kmipbin.KmipBoolean" {
			fmt.Println(field.Name)
		}

		//		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.String(), tag)
	}

	return nil, nil
	//
	//}*/
}

/*
func (r *RequestHeader) Unmarshal(b []byte) ([]byte, error) {

	for {
		if len(b) <= 0 {
			return b, nil
		}
		tag := kmipbin.KmipTagType{}
		tag.UnMarshalBin(b[:4])
		switch int(tag.Tag) {
		case 0x420069:
			b = b[8:]
			r.ProtocolVersion = &ProtocolVersion{}
			b, _ = r.ProtocolVersion.Unmarshal(b)
		case 0x42000D:
			b = b[8:]
			var bc kmipbin.KmipInt
			r.BatchCount = &bc
			r.BatchCount.UnMarshalBin(b[:8])
			b = b[8:]
		default:
			return b, nil

		}
	}
}*/
