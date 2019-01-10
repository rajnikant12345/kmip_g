package model

//resultstatus
const (
	Success          = 0
	OperationFailed  = 1
	OperationPending = 2
	OperationUndone  = 3
)

//ResultReasonEnumeration
const (
	ItemNotFound                     = 1
	ResponseTooLarge                 = 2
	AuthenticationNotSuccessful      = 3
	InvalidMessage                   = 4
	OperationNotSupported            = 5
	MissingData                      = 6
	InvalidField                     = 7
	FeatureNotSupported              = 8
	OperationCanceledByRequester     = 9
	CryptographicFailure             = 10
	IllegalOperation                 = 11
	PermissionDenied                 = 12
	ObjectArchived                   = 13
	IndexOutofBounds                 = 14
	ApplicationNamespaceNotSupported = 15
	KeyFormatTypeNotSupported        = 16
	KeyCompressionTypeNotSupported   = 17
	EncodingOptionError              = 18
	KeyValueNotPresent               = 19
	AttestationRequired              = 20
	AttestationFailed                = 21
	GeneralFailure                   = 0x100
)

//CredentialType
const (
	UsernameAndPassword = 0x00000001
	Device              = 0x00000002
	Attestation         = 0x00000003
)

//Compressiontype

const (
	ECPublicKeyTypeUncompressed = 0x00000001

	ECPublicKeyTypeX962CompressedPrime = 0x00000002

	ECPublicKeyTypeX962CompressedChar2 = 0x00000003

	ECPublicKeyTypeX962Hybrid = 0x00000004
)

//Key Format Type
const (
	Raw = 0x00000001

	Opaque = 0x00000002

	PKCS1 = 0x00000003

	PKCS8 = 0x00000004

	X509KeyFormatType = 0x00000005

	ECPrivateKey = 0x00000006

	TransparentSymmetricKey = 0x00000007

	TransparentDSAPrivateKey = 0x00000008

	TransparentDSAPublicKey = 0x00000009

	TransparentRSAPrivateKey = 0x0000000A

	TransparentRSAPublicKey = 0x0000000B

	TransparentDHPrivateKey = 0x0000000C

	TransparentDHPublicKey = 0x0000000D

	TransparentECDSAPrivateKey = 0x0000000E

	TransparentECDSAPublicKey = 0x0000000F

	TransparentECDHPrivateKey = 0x00000010

	TransparentECDHPublicKey = 0x00000011

	TransparentECMQVPrivateKey = 0x00000012

	TransparentECMQVPublicKey = 0x00000013
)

//Wrapping Method
const (
	Encrypt = 0x00000001

	MACsign = 0x00000002

	EncryptthenMACsign = 0x00000003

	MACsignThenencrypt = 0x00000004

	TR31 = 0x00000005
)

// Recommended Curve Enumeration
const (
	P192 = 0x00000001

	K163 = 0x00000002

	B163 = 0x00000003

	P224 = 0x00000004

	K233 = 0x00000005

	B233 = 0x00000006

	P256 = 0x00000007

	K283 = 0x00000008

	B283 = 0x00000009

	P384 = 0x0000000A

	K409 = 0x0000000B

	B409 = 0x0000000C

	P521 = 0x0000000D

	K571 = 0x0000000E

	B571 = 0x0000000F

	SECP112R1 = 0x00000010

	SECP112R2 = 0x00000011

	SECP128R1 = 0x00000012

	SECP128R2 = 0x00000013

	SECP160K1 = 0x00000014

	SECP160R1 = 0x00000015

	SECP160R2 = 0x00000016

	SECP192K1 = 0x00000017

	SECP224K1 = 0x00000018

	SECP256K1 = 0x00000019

	SECT113R1 = 0x0000001A

	SECT113R2 = 0x0000001B

	SECT131R1 = 0x0000001C

	SECT131R2 = 0x0000001D

	SECT163R1 = 0x0000001E

	SECT193R1 = 0x0000001F

	SECT193R2 = 0x00000020

	SECT239K1 = 0x00000021

	ANSIX9P192V2 = 0x00000022

	ANSIX9P192V3 = 0x00000023

	ANSIX9P239V1 = 0x00000024

	ANSIX9P239V2 = 0x00000025

	ANSIX9P239V3 = 0x00000026

	ANSIX9C2PNB163V1 = 0x00000027

	ANSIX9C2PNB163V2 = 0x00000028

	ANSIX9C2PNB163V3 = 0x00000029

	ANSIX9C2PNB176V1 = 0x0000002A

	ANSIX9C2TNB191V1 = 0x0000002B

	ANSIX9C2TNB191V2 = 0x0000002C

	ANSIX9C2TNB191V3 = 0x0000002D

	ANSIX9C2PNB208W1 = 0x0000002E

	ANSIX9C2TNB239V1 = 0x0000002F

	ANSIX9C2TNB239V2 = 0x00000030

	ANSIX9C2TNB239V3 = 0x00000031

	ANSIX9C2PNB272W1 = 0x00000032

	ANSIX9C2PNB304W1 = 0x00000033

	ANSIX9C2TNB359V1 = 0x00000034

	ANSIX9C2PNB368W1 = 0x00000035

	ANSIX9C2TNB431R1 = 0x00000036

	BRAINPOOLP160R1 = 0x00000037

	BRAINPOOLP160T1 = 0x00000038

	BRAINPOOLP192R1 = 0x00000039

	BRAINPOOLP192T1 = 0x0000003A

	BRAINPOOLP224R1 = 0x0000003B

	BRAINPOOLP224T1 = 0x0000003C

	BRAINPOOLP256R1 = 0x0000003D

	BRAINPOOLP256T1 = 0x0000003E

	BRAINPOOLP320R1 = 0x0000003F

	BRAINPOOLP320T1 = 0x00000040

	BRAINPOOLP384R1 = 0x00000041

	BRAINPOOLP384T1 = 0x00000042

	BRAINPOOLP512R1 = 0x00000043

	BRAINPOOLP512T1 = 0x00000044
)

// Certificate Type
const (
	X509CetificateType = 0x00000001
	CertificatePGP                = 0x00000002
)

// Digital Signature Algorithm
const (
	MD2withRSAEncryption = 0x00000001

	MD5withRSAEncryption = 0x00000002

	SHA1withRSAEncryption = 0x00000003

	SHA224withRSAEncryption = 0x00000004

	SHA256withRSAEncryption = 0x00000005

	SHA384withRSAEncryption = 0x00000006

	SHA512withRSAEncryption = 0x00000007

	RSASSAPSS = 0x00000008

	DSAwithSHA1 = 0x00000009

	DSAwithSHA224 = 0x0000000A

	DSAwithSHA256 = 0x0000000B

	ECDSAwithSHA1 = 0x0000000C

	ECDSAwithSHA224 = 0x0000000D

	ECDSAwithSHA256 = 0x0000000E

	ECDSAwithSHA384 = 0x0000000F

	ECDSAwithSHA512 = 0x00000010
)

// Split Key Method
const (
	XOR = 0x00000001

	PolynomialSharingGF216 = 0x00000002

	PolynomialSharingPrimeField = 0x00000003

	PolynomialSharingGF28 = 0x00000004
)

// Secret Data Type
const (
	Password = 0x00000001

	Seed = 0x00000002
)

// Name Type
const (
	NameUninterpretedTextString = 0x00000001

	NameURI = 00000002
)

// object type
const (
	Certificate = 0x00000001

	SymmetricKey = 0x00000002

	PublicKey = 0x00000003

	PrivateKey = 0x00000004

	SplitKey = 0x00000005

	Template = 0x00000006

	SecretData = 0x00000007

	OpaqueObject = 0x00000008

	PGPKey = 0x00000009
)

// crypto algorithms
const (
	DES = 0x00000001

	TDES = 0x00000002

	AES = 0x00000003

	RSA = 0x00000004

	DSA = 0x00000005

	ECDSA = 0x00000006

	HMACSHA1 = 0x00000007

	HMACSHA224 = 0x00000008

	HMACSHA256 = 0x00000009

	HMACSHA384 = 0x0000000A

	HMACSHA512 = 0x0000000B

	HMACMD5 = 0x0000000C

	DH = 0x0000000D

	ECDH = 0x0000000E

	ECMQV = 0x0000000F

	Blowfish = 0x00000010

	Camellia = 0x00000011

	CAST5 = 0x00000012

	IDEA = 0x00000013

	MARS = 0x00000014

	RC2 = 0x00000015

	RC4 = 0x00000016

	RC5 = 0x00000017

	SKIPJACK = 0x00000018

	Twofish = 0x00000019

	EC = 0x0000001A
)

// Block Cipher Mode Enumeration
const (
	CBC = 0x00000001

	ECB = 0x00000002

	PCBC = 0x00000003

	CFB = 0x00000004

	OFB = 0x00000005

	CTR = 0x00000006

	CMAC = 0x00000007

	CCM = 0x00000008

	GCM = 0x00000009

	CBCMAC = 0x0000000A

	XTS = 0x0000000B

	AESKeyWrapPadding = 0x0000000C

	NISTKeyWrap = 0x0000000D

	X9102AESKW = 0x0000000E

	X9102TDKW = 0x0000000F

	X9102AKW1 = 0x00000010

	X9102AKW2 = 0x00000011
)

//////Padding Method/////////
const (
	None = 0x00000001

	OAEP = 0x00000002

	PKCS5 = 0x00000003

	SSL3 = 0x00000004

	Zeros = 0x00000005

	ANSIX923 = 0x00000006

	ISO10126 = 0x00000007

	PKCS1v15 = 0x00000008

	X931 = 0x00000009

	PSS = 0x0000000A
)

//////Hashing Algorithm//////

const (
	MD2 = 0x00000001

	MD4 = 0x00000002

	MD5 = 0x00000003

	SHA1 = 0x00000004

	SHA224 = 0x00000005

	SHA256 = 0x00000006

	SHA384 = 0x00000007

	SHA512 = 0x00000008

	RIPEMD160 = 0x00000009

	Tiger = 0x0000000A

	Whirlpool = 0x0000000B

	SHA512224 = 0x0000000C

	SHA512256 = 0x0000000D
)

/////////Key Role Type///////////

const (
	BDK = 0x00000001

	CVK = 0x00000002

	DEK = 0x00000003

	MKAC = 0x00000004

	MKSMC = 0x00000005

	MKSMI = 0x00000006

	MKDAC = 0x00000007

	MKDN = 0x00000008

	MKCP = 0x00000009

	MKOTH = 0x0000000A

	KEK = 0x0000000B

	MAC16609 = 0x0000000C

	MAC97971 = 0x0000000D

	MAC97972 = 0x0000000E

	MAC97973 = 0x0000000F

	MAC97974 = 0x00000010

	MAC97975 = 0x00000011

	ZPK = 0x00000012

	PVKIBM = 0x00000013

	PVKPVV = 0x00000014

	PVKOTH = 0x00000015
)

////////////////////////////////////////

const (
	/////////////

	PreActive = 0x00000001

	Active = 0x00000002

	Deactivated = 0x00000003

	Compromised = 0x00000004

	Destroyed = 0x00000005

	DestroyedCompromised = 0x00000006

	////////////////////

	Unspecified = 0x00000001

	KeyCompromise = 0x00000002

	CACompromise = 0x00000003

	AffiliationChanged = 0x00000004

	Superseded = 0x00000005

	CessationofOperation = 0x00000006

	PrivilegeWithdrawn = 0x00000007

	////////////////

	CertificateLink = 0x00000101

	PublicKeyLink = 0x00000102

	PrivateKeyLink = 0x00000103

	DerivationBaseObjectLink = 0x00000104

	DerivedKeyLink = 0x00000105

	ReplacementObjectLink = 0x00000106

	ReplacedObjectLink = 0x00000107

	ParentLink = 0x00000108

	ChildLink = 0x00000109

	PreviousLink = 0x0000010A

	NextLink = 0x0000010B

	//////////////////

	PBKDF2 = 0x00000001

	HASH = 0x00000002

	HMAC = 0x00000003

	ENCRYPT = 0x00000004

	NIST800108C = 0x00000005

	NIST800108F = 0x00000006

	NIST800108DPI = 0x00000007

	/////////////////////

	CRMF = 0x00000001

	PKCS10 = 0x00000002

	PEM = 0x00000003

	PGP = 0x00000004

	/////////////////////

	Valid = 0x00000001

	Invalid = 0x00000002

	Unknown = 0x00000003

	/////////////////

	QueryOperations = 0x00000001

	QueryObjects = 0x00000002

	QueryServerInformation = 0x00000003

	QueryApplicationNamespaces = 0x00000004

	QueryExtensionList = 0x00000005

	QueryExtensionMap = 0x00000006

	QueryAttestationTypes = 0x00000007

	/////////////////////

	Canceled = 0x00000001

	UnabletoCancel = 0x00000002

	Completed = 0x00000003

	Failed = 0x00000004

	Unavailable = 0x00000005

	//////////////////

	New = 0x00000001

	Replace = 0x00000002

	///////////////////////

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

	///////////////////

	Continue = 0x00000001

	Stop = 0x00000002

	Undo = 0x00000003

	/////////////////

	Byte = 0x00000001

	Object = 0x00000002

	////////////////////

	NoEncoding = 0x00000001

	TTLVEncoding = 0x00000002

	///////////////

	GroupMemberFresh = 0x00000001

	GroupMemberDefault = 0x00000002

	////////Alternative Name Type//////////

	UninterpretedTextString = 0x00000001

	URI = 0x00000002

	ObjectSerialNumber = 0x00000003

	EmailAddress = 0x00000004

	DNSName = 0x00000005

	X500DistinguishedName = 0x00000006

	IPAddress = 0x00000007

	/////////Key Value Location Type////////////////

	LocationUninterpretedTextString = 0x00000001

	LocationURI = 0x00000002

	///////////////////////////

	TPMQuote = 0x00000001

	TCGIntegrityReport = 0x00000002

	SAMLAssertion = 0x00000003

	//////////Usage Mask/////////////////

	Sign = 0x00000001

	Verify = 0x00000002

	UsageEncrypt = 0x00000004

	Decrypt = 0x00000008

	WrapKey = 0x00000010

	UnwrapKey = 0x00000020

	Export = 0x00000040

	MACGenerate = 0x00000080

	MACVerify = 0x00000100

	DeriveKey = 0x00000200

	ContentCommitment = 0x00000400

	KeyAgreement = 0x00000800

	CertificateSign = 0x00001000

	CRLSign = 0x00002000

	GenerateCryptogram = 0x00004000

	ValidateCryptogram = 0x00008000

	TranslateEncrypt = 0x00010000

	TranslateDecrypt = 0x00020000

	TranslateWrap = 0x00040000

	TranslateUnwrap = 0x00080000

	//////////////////////////////////

	Onlinestorage = 0x00000001

	Archivalstorage = 0x00000002
)
