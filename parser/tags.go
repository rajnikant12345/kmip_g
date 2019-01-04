package parser

var Tags = map[string]string{
	"None":                                  "000000",
	"ActivationDate":                        "420001",
	"ApplicationData":                       "420002",
	"ApplicationNamespace":                  "420003",
	"ApplicationSpecificInformation":        "420004",
	"ArchiveDate":                           "420005",
	"AsynchronousCorrelationValue":          "420006",
	"AsynchronousIndicator":                 "420007",
	"Attribute":                             "420008",
	"AttributeIndex":                        "420009",
	"AttributeName":                         "42000A",
	"AttributeValue":                        "42000B",
	"Authentication":                        "42000C",
	"BatchCount":                            "42000D",
	"BatchErrorContinuationOption":          "42000E",
	"BatchItem":                             "42000F",
	"BatchOrderOption":                      "420010",
	"BlockCipherMode":                       "420011",
	"CancellationResult":                    "420012",
	"Certificate":                           "420013",
	"CertificateIdentifier":                 "420014",
	"CertificateIssuer":                     "420015",
	"CertificateIssuerAlternativeName":      "420016",
	"CertificateIssuerDistinguishedName":    "420017",
	"CertificateRequest":                    "420018",
	"CertificateRequestType":                "420019",
	"CertificateSubject":                    "42001A",
	"CertificateSubjectAlternativeName":     "42001B",
	"CertificateSubjectDistinguishedName":   "42001C",
	"CertificateType":                       "42001D",
	"CertificateValue":                      "42001E",
	"CommonTemplateAttribute":               "42001F",
	"CompromiseDate":                        "420020",
	"CompromiseOccurrenceDate":              "420021",
	"ContactInformation":                    "420022",
	"Credential":                            "420023",
	"CredentialType":                        "420024",
	"CredentialValue":                       "420025",
	"CriticalityIndicator":                  "420026",
	"CRTCoefficient":                        "420027",
	"CryptographicAlgorithm":                "420028",
	"CryptographicDomainParameters":         "420029",
	"CryptographicLength":                   "42002A",
	"CryptographicParameters":               "42002B",
	"CryptographicUsageMask":                "42002C",
	"CustomAttribute":                       "42002D",
	"D":                                     "42002E",
	"DeactivationDate":                      "42002F",
	"DerivationData":                        "420030",
	"DerivationMethod":                      "420031",
	"DerivationParameters":                  "420032",
	"DestroyDate":                           "420033",
	"Digest":                                "420034",
	"DigestValue":                           "420035",
	"EncryptionKeyInformation":              "420036",
	"G":                                     "420037",
	"HashingAlgorithm":                      "420038",
	"InitialDate":                           "420039",
	"InitializationVector":                  "42003A",
	"Issuer":                                "42003B",
	"IterationCount":                        "42003C",
	"IV/Counter/Nonce":                      "42003D",
	"J":                                     "42003E",
	"Key":                                   "42003F",
	"KeyBlock":                              "420040",
	"KeyCompressionType":                    "420041",
	"KeyFormatType":                         "420042",
	"KeyMaterial":                           "420043",
	"KeyPartIdentifier":                     "420044",
	"KeyValue":                              "420045",
	"KeyWrappingData":                       "420046",
	"KeyWrappingSpecification":              "420047",
	"LastChangeDate":                        "420048",
	"LeaseTime":                             "420049",
	"Link":                                  "42004A",
	"LinkType":                              "42004B",
	"LinkedObjectIdentifier":                "42004C",
	"MAC/Signature":                         "42004D",
	"MAC/SignatureKeyInformation":           "42004E",
	"MaximumItems":                          "42004F",
	"MaximumResponseSize":                   "420050",
	"MessageExtension":                      "420051",
	"Modulus":                               "420052",
	"Name":                                  "420053",
	"NameType":                              "420054",
	"NameValue":                             "420055",
	"ObjectGroup":                           "420056",
	"ObjectType":                            "420057",
	"Offset":                                "420058",
	"OpaqueDataType":                        "420059",
	"OpaqueDataValue":                       "42005A",
	"OpaqueObject":                          "42005B",
	"Operation":                             "42005C",
	"OperationPolicyName":                   "42005D",
	"P":                                     "42005E",
	"PaddingMethod":                         "42005F",
	"PrimeExponentP":                        "420060",
	"PrimeExponentQ":                        "420061",
	"PrimeFieldSize":                        "420062",
	"PrivateExponent":                       "420063",
	"PrivateKey":                            "420064",
	"PrivateKeyTemplateAttribute":           "420065",
	"PrivateKeyUniqueIdentifier":            "420066",
	"ProcessStartDate":                      "420067",
	"ProtectStopDate":                       "420068",
	"ProtocolVersion":                       "420069",
	"ProtocolVersionMajor":                  "42006A",
	"ProtocolVersionMinor":                  "42006B",
	"PublicExponent":                        "42006C",
	"PublicKey":                             "42006D",
	"PublicKeyTemplateAttribute":            "42006E",
	"PublicKeyUniqueIdentifier":             "42006F",
	"PutFunction":                           "420070",
	"Q":                                     "420071",
	"QString":                               "420072",
	"Qlength":                               "420073",
	"QueryFunction":                         "420074",
	"RecommendedCurve":                      "420075",
	"ReplacedUniqueIdentifier":              "420076",
	"RequestHeader":                         "420077",
	"RequestMessage":                        "420078",
	"RequestPayload":                        "420079",
	"ResponseHeader":                        "42007A",
	"ResponseMessage":                       "42007B",
	"ResponsePayload":                       "42007C",
	"ResultMessage":                         "42007D",
	"ResultReason":                          "42007E",
	"ResultStatus":                          "42007F",
	"RevocationMessage":                     "420080",
	"RevocationReason":                      "420081",
	"RevocationReasonCode":                  "420082",
	"KeyRoleType":                           "420083",
	"Salt":                                  "420084",
	"SecretData":                            "420085",
	"SecretDataType":                        "420086",
	"SerialNumber":                          "420087",
	"ServerInformation":                     "420088",
	"SplitKey":                              "420089",
	"SplitKeyMethod":                        "42008A",
	"SplitKeyParts":                         "42008B",
	"SplitKeyThreshold":                     "42008C",
	"State":                                 "42008D",
	"StorageStatusMask":                     "42008E",
	"SymmetricKey":                          "42008F",
	"Template":                              "420090",
	"TemplateAttribute":                     "420091",
	"TimeStamp":                             "420092",
	"UniqueBatchItemID":                     "420093",
	"UniqueIdentifier":                      "420094",
	"UsageLimits":                           "420095",
	"UsageLimitsCount":                      "420096",
	"UsageLimitsTotal":                      "420097",
	"UsageLimitsUnit":                       "420098",
	"Username":                              "420099",
	"ValidityDate":                          "42009A",
	"ValidityIndicator":                     "42009B",
	"VendorExtension":                       "42009C",
	"VendorIdentification":                  "42009D",
	"WrappingMethod":                        "42009E",
	"X":                                     "42009F",
	"Y":                                     "4200A0",
	"Password":                              "4200A1",
	"DeviceIdentifier":                      "4200A2",
	"EncodingOption":                        "4200A3",
	"ExtensionInformation":                  "4200A4",
	"ExtensionName":                         "4200A5",
	"ExtensionTag":                          "4200A6",
	"ExtensionType":                         "4200A7",
	"Fresh":                                 "4200A8",
	"MachineIdentifier":                     "4200A9",
	"MediaIdentifier":                       "4200AA",
	"NetworkIdentifier":                     "4200AB",
	"ObjectGroupMember":                     "4200AC",
	"CertificateLength":                     "4200AD",
	"DigitalSignatureAlgorithm":             "4200AE",
	"CertificateSerialNumber":               "4200AF",
	"DeviceSerialNumber":                    "4200B0",
	"IssuerAlternativeName":                 "4200B1",
	"IssuerDistinguishedName":               "4200B2",
	"SubjectAlternativeName":                "4200B3",
	"SubjectDistinguishedName":              "4200B4",
	"X.509CertificateIdentifier":            "4200B5",
	"X.509CertificateIssuer":                "4200B6",
	"X.509CertificateSubject":               "4200B7",
	"KeyValueLocation":                      "4200B8",
	"KeyValueLocationValue":                 "4200B9",
	"KeyValueLocationType":                  "4200BA",
	"KeyValuePresent":                       "4200BB",
	"OriginalCreationDate":                  "4200BC",
	"PGPKey":                                "4200BD",
	"PGPKeyVersion":                         "4200BE",
	"AlternativeName":                       "4200BF",
	"AlternativeNameValue":                  "4200C0",
	"AlternativeNameType":                   "4200C1",
	"Data":                                  "4200C2",
	"SignatureData":                         "4200C3",
	"DataLength":                            "4200C4",
	"RandomIV":                              "4200C5",
	"MACData":                               "4200C6",
	"AttestationType":                       "4200C7",
	"Nonce":                                 "4200C8",
	"NonceID":                               "4200C9",
	"NonceValue":                            "4200CA",
	"AttestationMeasurement":                "4200CB",
	"AttestationAssertion":                  "4200CC",
	"IVLength":                              "4200CD",
	"TagLength":                             "4200CE",
	"FixedFieldLength":                      "4200CF",
	"CounterLength":                         "4200D0",
	"InitialCounterValue":                   "4200D1",
	"InvocationFieldLength":                 "4200D2",
	"AttestationCapableIndicator":           "4200D3",
	"OffsetItems":                           "4200D4",
	"LocatedItems":                          "4200D5",
	"CorrelationValue":                      "4200D6",
	"InitIndicator":                         "4200D7",
	"FinalIndicator":                        "4200D8",
	"RNGParameters":                         "4200D9",
	"RNGAlgorithm":                          "4200DA",
	"DRBGAlgorithm":                         "4200DB",
	"FIPS186Variation":                      "4200DC",
	"PredictionResistance":                  "4200DD",
	"RandomNumberGenerator":                 "4200DE",
	"ValidationInformation":                 "4200DF",
	"ValidationAuthorityType":               "4200E0",
	"ValidationAuthorityCountry":            "4200E1",
	"ValidationAuthorityURI":                "4200E2",
	"ValidationVersionMajor":                "4200E3",
	"ValidationVersionMinor":                "4200E4",
	"ValidationType":                        "4200E5",
	"ValidationLevel":                       "4200E6",
	"ValidationCertificateIdentifier":       "4200E7",
	"ValidationCertificateURI":              "4200E8",
	"ValidationVendorURI":                   "4200E9",
	"ValidationProfile":                     "4200EA",
	"ProfileInformation":                    "4200EB",
	"ProfileName":                           "4200EC",
	"ServerURI":                             "4200ED",
	"ServerPort":                            "4200EE",
	"StreamingCapability":                   "4200EF",
	"AsynchronousCapability":                "4200F0",
	"AttestationCapability":                 "4200F1",
	"UnwrapMode":                            "4200F2",
	"DestroyAction":                         "4200F3",
	"ShreddingAlgorithm":                    "4200F4",
	"RNGMode":                               "4200F5",
	"ClientRegistrationMethod":              "4200F6",
	"CapabilityInformation":                 "4200F7",
	"KeyWrapType":                           "4200F8",
	"BatchUndoCapability":                   "4200F9",
	"BatchContinueCapability":               "4200FA",
	"PKCS#12FriendlyName":                   "4200FB",
	"Description":                           "4200FC",
	"Comment":                               "4200FD",
	"AuthenticatedEncryptionAdditionalData": "4200FE",
	"AuthenticatedEncryptionTag":            "4200FF",
	"SaltLength":                            "420100",
	"MaskGenerator":                         "420101",
	"MaskGeneratorHashingAlgorithm":         "420102",
	"PSource":                               "420103",
	"TrailerField":                          "420104",
	"ClientCorrelationValue":                "420105",
	"ServerCorrelationValue":                "420106",
	"DigestedData":                          "420107",
	"CertificateSubjectCN":                  "420108",
	"CertificateSubjectO":                   "420109",
	"CertificateSubjectOU":                  "42010A",
	"CertificateSubjectEmail":               "42010B",
	"CertificateSubjectC":                   "42010C",
	"CertificateSubjectST":                  "42010D",
	"CertificateSubjectL":                   "42010E",
	"CertificateSubjectUID":                 "42010F",
	"CertificateSubjectSerialNumber":        "420110",
	"CertificateSubjectTitle":               "420111",
	"CertificateSubjectDC":                  "420112",
	"CertificateSubjectDNQualifier":         "420113",
	"CertificateIssuerCN":                   "420114",
	"CertificateIssuerO":                    "420115",
	"CertificateIssuerOU":                   "420116",
	"CertificateIssuerEmail":                "420117",
	"CertificateIssuerC":                    "420118",
	"CertificateIssuerST":                   "420119",
	"CertificateIssuerL":                    "42011A",
	"CertificateIssuerUID":                  "42011B",
	"CertificateIssuerSerialNumber":         "42011C",
	"CertificateIssuerTitle":                "42011D",
	"CertificateIssuerDC":                   "42011E",
	"CertificateIssuerDNQualifier":          "42011F",
	"Sensitive":                             "420120",
	"AlwaysSensitive":                       "420121",
	"Extractable":                           "420122",
	"NeverExtractable":                      "420123",
	"ReplaceExisting":                       "420124",
}
