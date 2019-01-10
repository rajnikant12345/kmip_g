package operation

const (
	QueryOperations = 0x00000001

	QueryObjects = 0x00000002

	QueryServerInformation = 0x00000003

	QueryApplicationNamespaces = 0x00000004

	QueryExtensionList = 0x00000005

	QueryExtensionMap = 0x00000006

	QueryAttestationTypes = 0x00000007
)

//Put Function

const (
	New = 0x00000001

	Replace = 0x00000002
)

//operations

const (
	Create = 0x00000001

	CreateKeyPair = 0x00000002

	Register = 0x00000003

	Rekey = 0x00000004

	OperationDeriveKey = 0x00000005

	Certify = 0x00000006

	Recertify = 0x00000007

	Locate = 0x00000008

	Check = 0x00000009

	Get = 0x0000000A

	GetAttributes = 0x0000000B

	GetAttributeList = 0x0000000C

	AddAttribute = 0x0000000D

	ModifyAttribute = 0x0000000E

	DeleteAttribute = 0x0000000F

	ObtainLease = 0x00000010

	GetUsageAllocation = 0x00000011

	Activate = 0x00000012

	Revoke = 0x00000013

	Destroy = 0x00000014

	Archive = 0x00000015

	Recover = 0x00000016

	Validate = 0x00000017

	Query = 0x00000018

	Cancel = 0x00000019

	Poll = 0x0000001A

	Notify = 0x0000001B

	Put = 0x0000001C

	RekeyKeyPair = 0x0000001D

	DiscoverVersions = 0x0000001E

	OperationEncrypt = 0x0000001F

	OperationDecrypt = 0x00000020

	OperatonSign = 0x00000021

	SignatureVerify = 0x00000022

	MAC = 0x00000023

	OperationMACVerify = 0x00000024

	RNGRetrieve = 0x00000025

	RNGSeed = 0x00000026

	Hash = 0x00000027

	CreateSplitKey = 0x00000028

	JoinSplitKey = 0x00000029
)

//Usage Limits Unit Enumeration

const (
	Byte = 0x00000001

	Object = 0x00000002
)

//Object Group Member Enumeration

const (
	GroupMemberFresh = 0x00000001

	GroupMemberDefault = 0x00000002
)


//Alternative Name Type


const (
	UninterpretedTextString = 0x00000001

	URI = 0x00000002

	ObjectSerialNumber = 0x00000003

	EmailAddress = 0x00000004

	DNSName = 0x00000005

	X500DistinguishedName = 0x00000006

	IPAddress = 0x00000007
)

// attestation

const (
	TPMQuote = 0x00000001

	TCGIntegrityReport = 0x00000002

	SAMLAssertion = 0x00000003
)

// storage

const (
	Onlinestorage = 0x00000001

	Archivalstorage = 0x00000002
)