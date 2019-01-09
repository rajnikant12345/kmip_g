package objects

import "github.com/rajnikant12345/kmip_g/kmipbin"

type RequestPayload struct {
	ObjectType                   *kmipbin.KmipEnum
	UniqueIdentifier             []*kmipbin.KmipTextString
	CertificateRequestType		*kmipbin.KmipEnum
	CertificateRequest			*kmipbin.KmipByteString
	KeyWrappingSpecification	*KeyWrappingSpecification
	SplitKeyParts				*kmipbin.KmipInt
	SplitKeyThreshold			*kmipbin.KmipInt
	SplitKeyMethod				*kmipbin.KmipEnum
	TemplateAttribute            *TemplateAttribute
	SplitKey					*SplitKey
	PutFunction					*kmipbin.KmipEnum
	SymmetricKey				*SymmetricKey
	Attribute                    []*Attribute
	AsynchronousCorrelationValue *kmipbin.KmipByteString
	CommonTemplateAttribute		 *CommonTemplateAttribute
	PrivateKeyTemplateAttribute  *PrivateKeyTemplateAttribute
	PublicKeyTemplateAttribute   *PublicKeyTemplateAttribute
	Template                     *Template
	SecretData                   *SecretData
	PGPKey						*PGPKey
	PrivateKey                   *PrivateKey
	PublicKey                    *PublicKey
	QueryFunction                []*kmipbin.KmipEnum
	AttributeName				 []*kmipbin.KmipTextString
	RevocationReason			 *RevocationReason
	CompromiseOccurrenceDate	*kmipbin.KmipDate

	Certificate					*Certificate
	KeyFormatType				*kmipbin.KmipEnum
	ProtocolVersion				*ProtocolVersion
}

type ResponsePayload struct {
	ObjectType                   *kmipbin.KmipEnum
	UniqueIdentifier             []*kmipbin.KmipTextString
	CertificateRequestType		*kmipbin.KmipEnum
	CertificateRequest			*kmipbin.KmipByteString
	KeyWrappingSpecification	*KeyWrappingSpecification
	SplitKeyParts				*kmipbin.KmipInt
	SplitKeyThreshold			*kmipbin.KmipInt
	SplitKeyMethod				*kmipbin.KmipEnum
	TemplateAttribute            *TemplateAttribute
	SplitKey					*SplitKey
	PutFunction					*kmipbin.KmipEnum
	SymmetricKey				*SymmetricKey
	Attribute                    []*Attribute
	AsynchronousCorrelationValue *kmipbin.KmipByteString
	CommonTemplateAttribute		 *CommonTemplateAttribute
	PrivateKeyTemplateAttribute  *PrivateKeyTemplateAttribute
	PublicKeyTemplateAttribute   *PublicKeyTemplateAttribute
	Template                     *Template
	SecretData                   *SecretData
	PGPKey						*PGPKey
	PrivateKey                   *PrivateKey
	PublicKey                    *PublicKey
	QueryFunction                []*kmipbin.KmipEnum
	AttributeName				 []*kmipbin.KmipTextString
	RevocationReason			 *RevocationReason
	CompromiseOccurrenceDate	*kmipbin.KmipDate
	Certificate					*Certificate
	KeyFormatType				*kmipbin.KmipEnum
	ProtocolVersion				*ProtocolVersion
}

type MessageExtension struct {
	VendorIdentification *kmipbin.KmipTextString
	CriticalityIndicator *kmipbin.KmipBoolean
	VendorExtension      *VendorExtension
}

type VendorExtension struct {
}

type BatchItem struct {
	Operation                    *kmipbin.KmipEnum
	UniqueBatchItemID            *kmipbin.KmipByteString
	ResultStatus                 *kmipbin.KmipEnum
	ResultReason                 *kmipbin.KmipEnum
	ResultMessage                *kmipbin.KmipTextString
	AsynchronousCorrelationValue *kmipbin.KmipByteString
	ResponsePayload              *ResponsePayload
	RequestPayload               *RequestPayload
	MessageExtension             *MessageExtension
}

type ProtocolVersion struct {
	ProtocolVersionMajor *kmipbin.KmipInt
	ProtocolVersionMinor *kmipbin.KmipInt
}

type RequestHeader struct {
	ProtocolVersion              *ProtocolVersion
	MaximumResponseSize          *kmipbin.KmipInt
	AsynchronousIndicator        *kmipbin.KmipBoolean
	AttestationCapableIndicator  *kmipbin.KmipBoolean
	AttestationType              *kmipbin.KmipEnum
	Authentication               *Authentication
	BatchErrorContinuationOption *kmipbin.KmipEnum
	BatchOrderOption             *kmipbin.KmipBoolean
	TimeStamp                    *kmipbin.KmipDate
	BatchCount                   *kmipbin.KmipInt
}

type Authentication struct {
	Credential *Credential
}

type RequestMessage struct {
	RequestHeader *RequestHeader
	BatchItem     []*BatchItem
}

type KmipStruct struct {
	RequestMessage  *RequestMessage
	ResponseMessage *ResponseMessage
}

type KmipStructResponse struct {
	ResponseMessage *ResponseMessage
}

type ResponseMessage struct {
	ResponseHeader *ResponseHeader
	BatchItem      []*BatchItem
}

type ResponseHeader struct {
	ProtocolVersion *ProtocolVersion
	TimeStamp       *kmipbin.KmipDate
	Nonce           *Nonce
	AttestationType *kmipbin.KmipEnum
	BatchCount      *kmipbin.KmipInt
}
