package base_objects

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

type Attribute struct {
	AttributeName  *kmipbin.KmipTextString
	AttributeIndex *kmipbin.KmipInt
	AttributeValue interface{}
}

type Name struct {
	NameValue *kmipbin.KmipTextString
	NameType  *kmipbin.KmipEnum
}

type Nonce struct {
	NonceID    *kmipbin.KmipByteString
	NonceValue *kmipbin.KmipByteString
}

type CredentialValue struct {
	Username               *kmipbin.KmipTextString
	Password               *kmipbin.KmipTextString
	DeviceSerialNumber     *kmipbin.KmipTextString
	DeviceIdentifier       *kmipbin.KmipTextString
	NetworkIdentifier      *kmipbin.KmipTextString
	MachineIdentifier      *kmipbin.KmipTextString
	MediaIdentifier        *kmipbin.KmipTextString
	Nonce                  *Nonce
	AttestationType        *kmipbin.KmipEnum
	AttestationMeasurement *kmipbin.KmipByteString
	AttestationAssertion   *kmipbin.KmipByteString
}

type Credential struct {
	CredentialType  *kmipbin.KmipEnum
	CredentialValue *CredentialValue
}

type KeyBlock struct {
	KeyFormatType          *kmipbin.KmipEnum
	KeyCompressionTypem    *kmipbin.KmipEnum
	KeyValue               interface{}
	CryptographicAlgorithm *kmipbin.KmipEnum
	CryptographicLength    *kmipbin.KmipInt
	KeywrappingData        *KeyWrappingData
}

type KeyValue struct {
	KeyMaterial interface{}
	Attribute   *[]Attribute
}

type KeyWrappingData struct {
	WrappingMethod             *kmipbin.KmipEnum
	EncryptionKeyInformation   *EncryptionKeyInformation
	MACSignatureKeyInformation *MACSignatureKeyInformation
	MACSignature               *kmipbin.KmipByteString
	IVCounterNonce             *kmipbin.KmipByteString
	EncodingOption             *kmipbin.KmipEnum
}

type CryptographicParameters struct {
	BlockCipherMode           *kmipbin.KmipEnum
	PaddingMethod             *kmipbin.KmipEnum
	HashingAlgorithm          *kmipbin.KmipEnum
	KeyRoleType               *kmipbin.KmipEnum
	DigitalSignatureAlgorithm *kmipbin.KmipEnum
	CryptographicAlgorithm    *kmipbin.KmipEnum
	RandomIV                  *kmipbin.KmipBoolean
	IVLength                  *kmipbin.KmipInt
	TagLength                 *kmipbin.KmipInt
	FixedFieldLength          *kmipbin.KmipInt
	InvocationFieldLength     *kmipbin.KmipInt
	CounterLength             *kmipbin.KmipInt
	InitialCounterValue       *kmipbin.KmipInt
}

// MACSignatureKeyInformation 2.1.5 Table 11
type MACSignatureKeyInformation struct {
	UniqueIdentifier        *kmipbin.KmipTextString
	CryptographicParameters *CryptographicParameters
}

// EncryptionKeyInformation 2.1.5 Table 10
type EncryptionKeyInformation struct {
	UniqueIdentifier        *kmipbin.KmipTextString
	CryptographicParameters *CryptographicParameters
}

type KeyWrappingSpecification struct {
	WrappingMethod *kmipbin.KmipEnum
	EncryptionKeyInformation *EncryptionKeyInformation
	MACSignatureKeyInformation *MACSignatureKeyInformation
	AttributeName	*kmipbin.KmipTextString
	EncodingOption	*kmipbin.KmipEnum
}

