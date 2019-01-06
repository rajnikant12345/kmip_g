package parser

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
)


type RequestPayload struct {
	ObjectType                   *kmipbin.KmipEnum
	UniqueIdentifier             *kmipbin.KmipTextString
	TemplateAttribute            *TemplateAttribute
	Attribute                    []*Attribute
	AsynchronousCorrelationValue *kmipbin.KmipByteString
	PublicKeyTemplateAttribute   *PublicKeyTemplateAttribute
	PrivateKeyTemplateAttribute  *PrivateKeyTemplateAttribute
	Template                     *Template
	SecretData                   *SecretData
	PrivateKey					 *PrivateKey
	PublicKey					 *PublicKey
	QueryFunction				[]*kmipbin.KmipEnum
}

func (p *RequestPayload) GetObjectType() *kmipbin.KmipEnum {
	return p.ObjectType
}

func (p *RequestPayload) GetUniqueIdentifier() *kmipbin.KmipTextString {
	return p.UniqueIdentifier
}

func (p *RequestPayload) GetTempleteAttribute() *TemplateAttribute {
	return p.TemplateAttribute
}

func (p *RequestPayload) GetTemplete() *Template {
	return p.Template
}

func (p *RequestPayload) GetAttributeCount() int {
	return len(p.Attribute)
}

func (p *RequestPayload) GetTempleteAttributeCount() int {

	return len(p.Attribute)
}

func (p *RequestPayload) GetAttributeAt(index int) *Attribute {
	if p.GetAttributeCount()-1 > index {
		return nil
	}
	return p.Attribute[index]
}
