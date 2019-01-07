package objects

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"strings"
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
	if *a.AttributeName == "X.509 Certificate Subject" {
		a.AttributeValue = new(X509CertificateSubject)
	}
	if *a.AttributeName == "X.509 Certificate Issuer" {
		a.AttributeValue = new(X509CertificateIssuer)
	}
	if *a.AttributeName == "Certificate Identifier" {
		a.AttributeValue = new(CertificateIdentifier)
	}
	if *a.AttributeName == "Certificate Subject" {
		a.AttributeValue = new(CertificateSubject)
	}
	if *a.AttributeName == "Certificate Issuer" {
		a.AttributeValue = new(CertificateIssuer)
	}
	if *a.AttributeName == "Digital Signature Algorithm" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Digest" {
		a.AttributeValue = new(Digest)
	}
	if *a.AttributeName == "Operation Policy Name" {
		a.AttributeValue = new(kmipbin.KmipTextString)
	}
	if *a.AttributeName == "Cryptographic Usage Mask" {
		a.AttributeValue = new(kmipbin.KmipInt)
	}
	if *a.AttributeName == "Lease Time" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Usage Limits" {
		a.AttributeValue = new(UsageLimits)
	}
	if *a.AttributeName == "State" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Initial Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Activation Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Process Start Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Protect Stop Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Deactivation Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Destroy Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Compromise Occurrence Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Compromise Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Revocation Reason" {
		a.AttributeValue = new(RevocationReason)
	}
	if *a.AttributeName == "Archive Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Object Group" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Fresh" {
		a.AttributeValue = new(kmipbin.KmipBoolean)
	}
	if *a.AttributeName == "Link" {
		a.AttributeValue = new(Link)
	}
	if *a.AttributeName == "Application Specific Information" {
		a.AttributeValue = new(ApplicationSpecificInformation)
	}
	if *a.AttributeName == "Contact Information" {
		a.AttributeValue = new(kmipbin.KmipTextString)
	}
	if *a.AttributeName == "Last Change Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if strings.HasPrefix(string(*a.AttributeName),"x-") {
		a.AttributeValue = make([]byte,0)
	}
	if *a.AttributeName == "Alternative Name" {
		a.AttributeValue = new(AlternativeName)
	}
	if *a.AttributeName == "Key Value Present" {
		a.AttributeValue = new(kmipbin.KmipBoolean)
	}
	if *a.AttributeName == "Key Value Location" {
		a.AttributeValue = new(KeyValueLocation)
	}
	if *a.AttributeName == "Original Creation Date" {
		a.AttributeValue = new(KeyValueLocation)
	}
}
