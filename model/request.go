package model

import "github.com/rajnikant12345/kmip_g/objects"

func ValidateRequestMessage(r *objects.KmipStruct) error {

}

func GetProtocolVersion(r *objects.KmipStruct) *objects.ProtocolVersion {
	if r ==  nil {
		return nil
	}
	rm := r.RequestMessage
	if rm == nil {

	}
}
