package objects

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

type Attribute struct {
	AttributeName  *kmipbin.KmipTextString
	AttributeIndex *kmipbin.KmipInt
	AttributeValue interface{}
}



func (a *Attribute) CreateAttribute() {

	if a.AttributeName == nil {
		return
	}
	if *a.AttributeName == "Unique Identifier" {
		a.AttributeValue = new(kmipbin.KmipTextString)
	}
	if *a.AttributeName == "Name" {
		a.AttributeValue = new(Name)
	}
	if *a.AttributeName == "Object Type" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Cryptographic Algorithm" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Cryptographic Length" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Cryptographic Parameters" {
		a.AttributeValue = new(CryptographicParameters)
	}
	if *a.AttributeName == "Cryptographic Domain Parameters" {
		a.AttributeValue = new(CryptographicDomainParameters)
	}
	if *a.AttributeName == "Certificate Type" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Certificate Length" {
		a.AttributeValue = new(kmipbin.KmipInt)
	}
	if *a.AttributeName == "X.509 Certificate Identifier" {
		a.AttributeValue = new(X509CertificateIdentifier)
	}

}
