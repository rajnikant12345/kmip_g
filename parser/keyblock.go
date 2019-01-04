package parser

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

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

type KeyWrappingData struct {
	WrappingMethod             *kmipbin.KmipEnum
	EncryptionKeyInformation   *EncryptionKeyInformation
	MACSignatureKeyInformation *MACSignatureKeyInformation
	MACSignature               *kmipbin.KmipByteString
	IVCounterNonce             *kmipbin.KmipByteString
	EncodingOption             *kmipbin.KmipEnum
}

type KeyValue struct {
	// KeyMaterial should be []byte, one of the Transparent*Key structs, or a custom struct if KeyFormatType is
	// an extension.
	KeyMaterial interface{}
	Attribute   []*Attribute
}

type KeyBlock struct {
	KeyFormatType          *kmipbin.KmipEnum
	KeyCompressionType     *kmipbin.KmipEnum
	KeyValue               interface{}
	CryptographicAlgorithm *kmipbin.KmipEnum
	CryptographicLength    *kmipbin.KmipInt
	KeywrappingData        *KeyWrappingData
	// it will be filled and then we will figureout what to do with it
	RawData				   []byte
}

func (k *KeyBlock) Unmarshal(bet *[]byte) {
	var l kmipbin.KmipLength
	l.UnMarshalBin((*bet)[4:8])
	le := kmipbin.PadLength(int(l))
	////////////////////////////////////////////

	k.RawData = make([]byte, 8+le)
	copy(k.RawData, (*bet)[:8+le])

	/////////////////////////////////////////////
	*bet = (*bet)[8+le:]
}