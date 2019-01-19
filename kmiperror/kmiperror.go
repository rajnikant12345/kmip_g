package kmiperror

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"github.com/rajnikant12345/kmip_g/enums/resultreason"
	"github.com/rajnikant12345/kmip_g/enums/operation"
)

type KmipError struct {
	ResultReason  kmipbin.KmipEnum
	ResultStatus  kmipbin.KmipEnum
	ResultMessage kmipbin.KmipTextString
	Operation     kmipbin.KmipEnum
}


// General error codes
var MessageCannotBeParsed = KmipError{resultreason.InvalidMessage, resultreason.OperationFailed, "Message cannot be parsed", 0}
var InvalidMessageStructure = KmipError{resultreason.InvalidMessage, resultreason.OperationFailed, "Invalid Message structure" , 0}
var IllegalOperationStructure = KmipError{resultreason.IllegalOperation, resultreason.OperationFailed, "Illegal operation" , 0}
var PermissionDenied = KmipError{resultreason.PermissionDenied, resultreason.OperationFailed, "Permission denied" , 0}
var ItemNorFound = KmipError{resultreason.ItemNotFound, resultreason.OperationFailed, "Item not found" , 0}

var GeneralFailure = KmipError{resultreason.GeneralFailure, resultreason.OperationFailed, "GeneralFailure" , 0}
var InvalidField = KmipError{resultreason.InvalidField, resultreason.OperationFailed, "Type is not recognized" , 0}
var FeatureNotSupported = KmipError{resultreason.FeatureNotSupported, resultreason.OperationFailed, "Not Supported" , 0}

//create eror codes

var CreteObjectTypeNotFound = KmipError{resultreason.InvalidField, resultreason.OperationFailed, "Object Type is not recognized" , operation.Create}
var CreteObjectTemplateNotFound = KmipError{resultreason.ItemNotFound, resultreason.OperationFailed, "Templates that do not exist are given in request" , operation.Create}
var CreteObjectIncorrectAttributeValue = KmipError{resultreason.InvalidField, resultreason.OperationFailed, "Incorrect attribute value(s) specified" , operation.Create}
var CreteObjectError = KmipError{resultreason.CryptographicFailure, resultreason.OperationFailed, "Error creating cryptographic object" , operation.Create}
var CreteObjectErrorMultipleInstance = KmipError{resultreason.IndexOutofBounds, resultreason.OperationFailed, "Trying to set more instances than the server supports of an attribute that MAY have multiple instances" , operation.Create}
var CreateObjectDuplicateName = KmipError{resultreason.InvalidField, resultreason.OperationFailed, "Trying to create a new object with the same Name attribute value as an existing object" , operation.Create}
var CreateObjectApplicatioNamespace = KmipError{resultreason.ApplicationNamespaceNotSupported, resultreason.OperationFailed, "The particular Application Namespace is not supported, and Application Data cannot be generated if it was omitted from the client request" , operation.Create}
var CreateObjectTemplateArchived = KmipError{ resultreason.ObjectArchived, resultreason.OperationFailed, "Template object is archived" , operation.Create}




//Destroy error codes
var DestroyPermissionDenied = KmipError{resultreason.PermissionDenied, resultreason.OperationFailed, "Permission denied" , operation.Destroy}
var DestroyItemNotFound = KmipError{resultreason.ItemNotFound, resultreason.OperationFailed, "Item not found" , operation.Destroy}
var DestroyOblectArchived = KmipError{resultreason.ObjectArchived, resultreason.OperationFailed, "Invalid Message structure" , operation.Destroy}



