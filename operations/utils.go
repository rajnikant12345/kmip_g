package operations

import (
	"github.com/rajnikant12345/kmip_g/objects"
	"github.com/rajnikant12345/kmip_g/kmiperror"
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

func ReadUniqueIdOfPayLoad(payLoad *objects.RequestPayload , idPlaceHolder *kmipbin.KmipTextString  ) *kmipbin.KmipTextString {
	if len(payLoad.UniqueIdentifier) == 0 && payLoad.UniqueIdentifier[0] == nil && *idPlaceHolder == "" {
		return nil
	}
	if payLoad.UniqueIdentifier[0] != nil {
		return payLoad.UniqueIdentifier[0]
	}

	if *idPlaceHolder != "" {
		return idPlaceHolder
	}
	return nil

}

func GetAttribute(name string, m map[string]interface{}) interface{} {
	if val,ok :=  m[name];ok {
		return val
	}
	return nil
}

func ReadCommonTemplateAttributes( attribute *objects.TemplateAttribute) (map[string]interface{} , []*objects.Name , *kmiperror.KmipError) {
	var AttributeMap = make(map[string]interface{})

	var NameList []*objects.Name

	if attribute != nil {
		if len(attribute.Attribute) != 0 {
			for _, v := range attribute.Attribute {
				name := string(*v.AttributeName)
				_, ok := AttributeMap[name]
				if ok && *v.AttributeName != "Name" {
					return nil, nil, &kmiperror.CreteObjectErrorMultipleInstance
				}

				if *v.AttributeName == "Name" {
					NameList = append(NameList , v.AttributeValue.(*objects.Name))
				}else {
					AttributeMap[name] = v.AttributeValue
				}
			}
		}
	}
	return AttributeMap , NameList , nil
}

func ReadPublicKeyTemplateAttributes( attribute *objects.TemplateAttribute) (map[string]interface{} , []*objects.Name , *kmiperror.KmipError) {
	var AttributeMap = make(map[string]interface{})

	var NameList []*objects.Name

	if attribute != nil {
		if len(attribute.Attribute) != 0 {
			for _, v := range attribute.Attribute {
				name := string(*v.AttributeName)
				_, ok := AttributeMap[name]
				if ok && *v.AttributeName != "Name" {
					return nil, nil, &kmiperror.CreteObjectErrorMultipleInstance
				}

				if *v.AttributeName == "Name" {
					NameList = append(NameList , v.AttributeValue.(*objects.Name))
				}else {
					AttributeMap[name] = v.AttributeValue
				}
			}
		}
	}
	return AttributeMap , NameList , nil
}


func ReadPrivateKeyTemplateAttributes( attribute *objects.TemplateAttribute) (map[string]interface{} , []*objects.Name , *kmiperror.KmipError) {
	var AttributeMap = make(map[string]interface{})

	var NameList []*objects.Name

	if attribute != nil {
		if len(attribute.Attribute) != 0 {
			for _, v := range attribute.Attribute {
				name := string(*v.AttributeName)
				_, ok := AttributeMap[name]
				if ok && *v.AttributeName != "Name" {
					return nil, nil, &kmiperror.CreteObjectErrorMultipleInstance
				}

				if *v.AttributeName == "Name" {
					NameList = append(NameList , v.AttributeValue.(*objects.Name))
				}else {
					AttributeMap[name] = v.AttributeValue
				}
			}
		}
	}
	return AttributeMap , NameList , nil
}



func ReadTemplateAttributes( attribute *objects.TemplateAttribute) (map[string]interface{} , []*objects.Name , *kmiperror.KmipError) {

	var AttributeMap = make(map[string]interface{})

	var NameList []*objects.Name

	if attribute != nil {
		if len(attribute.Attribute) != 0 {
			for _, v := range attribute.Attribute {
				name := string(*v.AttributeName)
				_, ok := AttributeMap[name]
				if ok && *v.AttributeName != "Name" {
					return nil, nil, &kmiperror.CreteObjectErrorMultipleInstance
				}

				if *v.AttributeName == "Name" {
					NameList = append(NameList , v.AttributeValue.(*objects.Name))
				}else {
					AttributeMap[name] = v.AttributeValue
				}
			}
		}
	}
	return AttributeMap , NameList , nil
}

func ReadAttributes(attribute []*objects.Attribute) (map[string]interface{} , []*objects.Name , *kmiperror.KmipError) {
	var AttributeMap = make(map[string]interface{})

	var NameList []*objects.Name

	if attribute != nil {
		if len(attribute) != 0 {
			for _, v := range attribute {
				name := string(*v.AttributeName)
				_, ok := AttributeMap[name]
				if ok && *v.AttributeName != "Name" {
					return nil, nil, &kmiperror.CreteObjectErrorMultipleInstance
				}

				if *v.AttributeName == "Name" {
					NameList = append(NameList , v.AttributeValue.(*objects.Name))
				}else {
					AttributeMap[name] = v.AttributeValue
				}
			}
		}
	}
	return AttributeMap , NameList , nil

}
