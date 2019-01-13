package operations

import "github.com/rajnikant12345/kmip_g/kmipbin"

func GetAttribute(name kmipbin.KmipTextString, m map[kmipbin.KmipTextString]interface{}) interface{} {
	if val,ok :=  m[name];ok {
		return val
	}
	return nil
}
