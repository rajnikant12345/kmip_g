package parser

import (
	"encoding/hex"
	"fmt"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"reflect"
	"strings"
)

func IsCompound(value reflect.Value) bool {
	switch value.Type().String() {
	case "kmipbin.KmipInt":
		return false
	case "kmipbin.KmipBoolean":
		return false
	case "kmipbin.KmipEnum":
		return false
	case "kmipbin.KmipDate":
		return false
	case "kmipbin.KmipLongInt":
		return false
	case "kmipbin.KmipInterval":
		return false
	case "kmipbin.KmipTextString":
		return false
	case "kmipbin.KmipByteString":
		return false
	}

	return true
}

func forwadBytes(b []byte) []byte {
	return nil
}


/*
t := reflect.ValueOf(v)

		for i := 0; i < t.NumField(); i++ {
			// Get the field, returns https://golang.org/pkg/reflect/#StructField
			field := t.Field(i)

			c := reflect.New(field.Type().Elem())
			if IsCompound(field) {
				fmt.Println("{{{{{")
				fmt.Println(field.Type().String())
				fmt.Println("struct",c.Elem().Interface())
				UnmaeshalAll(c.Elem().Interface(), b)
				fmt.Println("}}}}}")

			}else {
				fmt.Println(field.Type().String())
				fmt.Println(c.Elem().Interface())
			}

			//c := reflect.New(field.Type().Elem())
			//c := field.Elem().Interface().(*ProtocolVersion)


		}

/*	case "*kmipbin.KmipBoolean":
				b = b[8:]
				var k  kmipbin.KmipBoolean
				k.UnMarshalBin(b)
				b = b[8:]
				v := reflect.New(v.Type().Elem())
				v.Set(reflect.ValueOf(&k))
			case "*kmipbin.KmipEnum":
				b = b[8:]
				var k  kmipbin.KmipEnum
				k.UnMarshalBin(b)
				b = b[8:]
				v := reflect.New(v.Type().Elem())
				v.Set(reflect.ValueOf(&k))
			case "*kmipbin.KmipDate":
				b = b[8:]
				var k  kmipbin.KmipDate
				k.UnMarshalBin(b)
				b = b[8:]
				v := reflect.New(v.Type().Elem())
				v.Set(reflect.ValueOf(&k))
			case "*kmipbin.KmipLongInt":
				b = b[8:]
				var k  kmipbin.KmipLongInt
				k.UnMarshalBin(b)
				b = b[8:]
				v := reflect.New(v.Type().Elem())
				v.Set(reflect.ValueOf(&k))
			case "*kmipbin.KmipInterval":
				b = b[8:]
				var k  kmipbin.KmipInterval
				k.UnMarshalBin(b)
				b = b[8:]
				v := reflect.New(v.Type().Elem())
				v.Set(reflect.ValueOf(&k))*/


func UnmaeshalAllT(a *KmipStruct, b []byte) {
	//ty := reflect.TypeOf(*a)
	v := reflect.ValueOf(a)

	h := v.Elem().Field(0)
	b = b[8:]
	Dummy(&h , b)
}

func Dummy( v *reflect.Value, b []byte ) {

	tagmap := make(map[string]reflect.Value)

	ty := reflect.TypeOf(v.Interface())

	fmt.Println(ty.String())

	for i:=0;i< v.NumField();i++ {
		field := v.Field(i)
		tag := ty.Field(i).Tag.Get("kmip")

		if tag != "" {
			tagmap[tag] = field
		}
	}

	for {
		if len(b) < 8 {
			break
		}
		tag := strings.ToUpper(hex.EncodeToString(b[:3]))
		fmt.Println(tag)
		f, ok := tagmap[tag]
		if !ok {
			return
			//b = forwadBytes(b)
		}else {
			fmt.Println("rajni",f.Type().String())
			switch f.Type().String() {
			case "kmipbin.KmipInt":
				b = b[8:]
				var k  kmipbin.KmipInt
				k.UnMarshalBin(b[:8])
				f.Set(reflect.ValueOf(k))
				b = b[8:]
			default:
				b = b[8:]
				//fmt.Println(reflect.TypeOf(v).String())
				Dummy(&f, b )
			}
		}

	}
}
