package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type Nonce struct {
	NonceID    kmipbin.KmipByteString
	NonceValue kmipbin.KmipByteString
}

type CredentialValue struct {
	Username               kmipbin.KmipTextString
	Password               kmipbin.KmipTextString
	DeviceSerialNumber     kmipbin.KmipTextString
	DeviceIdentifier       kmipbin.KmipTextString
	NetworkIdentifier      kmipbin.KmipTextString
	MachineIdentifier      kmipbin.KmipTextString
	MediaIdentifier        kmipbin.KmipTextString
	Nonce                  Nonce
	AttestationType        kmipbin.KmipEnum
	AttestationMeasurement kmipbin.KmipByteString
	AttestationAssertion   kmipbin.KmipByteString
}

type Credential struct {
	CredentialType  kmipbin.KmipEnum
	CredentialValue CredentialValue
}

type Authentication struct {
	Credential Credential
}
