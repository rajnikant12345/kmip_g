package objects

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
)


type KmipStruct struct {
	RequestMessage *RequestMessage
}

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
	ProtocolVersion				[]*ProtocolVersion
}

type ServerInformation struct {

}

type ResponsePayload struct {
	Operation					[]*kmipbin.KmipEnum
	ObjectType                   []*kmipbin.KmipEnum
	UniqueIdentifier             []*kmipbin.KmipTextString
	PrivateKeyUniqueIdentifier	 []*kmipbin.KmipTextString
	PublicKeyUniqueIdentifier	[]*kmipbin.KmipTextString
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
	ProtocolVersion				[]*ProtocolVersion
	VendorIdentification		*kmipbin.KmipTextString
	ServerInformation			*ServerInformation
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


func (r *KmipStruct) GetRequestMessage() *RequestMessage {
	if r.RequestMessage != nil {
		return r.RequestMessage
	}
	return nil
}

func (r *RequestMessage) GetRequestHeader() *RequestHeader {
	if r.RequestHeader != nil {
		return r.RequestHeader
	}
	return nil
}

func (r *ResponseHeader) SetBatchCount(val kmipbin.KmipInt) {
	r.BatchCount = &val
}

func (r *RequestMessage) GetBatchList( ) []*BatchItem {
	if len(r.BatchItem) != 0 {
		return r.BatchItem
	}
	return nil
}

func ( r *BatchItem) GetRequestPayLoad() *RequestPayload {
	if r.ResponsePayload != nil {
		return r.RequestPayload
	}
	return nil
}
